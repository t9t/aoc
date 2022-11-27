package year2018

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func init() {
	mustRegisterPair(24, Day24Part1, Day24Part2)
}

func Day24Part1(input string) (string, error) {
	return day24(input, false)
}

func Day24Part2(input string) (string, error) {
	return day24(input, true)
}

func day24(input string, part2 bool) (string, error) {
	parseTypes := func(s string) map[string]struct{} {
		if s == "" {
			return map[string]struct{}{}
		}
		// "[immune|weak] to radiation, bludgeoning"
		m := make(map[string]struct{})
		for _, item := range strings.Split(strings.Split(s, " to ")[1], ", ") {
			m[item] = struct{}{}
		}
		return m
	}

	parseWeaknessesAndImmunities := func(s string) (map[string]struct{}, map[string]struct{}) {
		if s == "" {
			return map[string]struct{}{}, map[string]struct{}{}
		}
		// " (weak to cold; immune to fire, bludgeoning)"
		parts := strings.Split(s[2:len(s)-1], "; ")
		if len(parts) == 1 {
			parts = append(parts, "")
		}
		left, right := parseTypes(parts[0]), parseTypes(parts[1])
		if strings.HasPrefix(parts[0], "immune") {
			return right, left
		}
		return left, right
	}

	// 989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3
	re := regexp.MustCompile(`(\d+) units each with (\d+) hit points( \(.+\))? with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	parseGroup := func(line string) (day24group, error) {
		matches := re.FindStringSubmatch(line)
		nums := make([]int, 4)
		g := day24group{}
		for i, s := range []string{matches[1], matches[2], matches[4], matches[6]} {
			if n, err := strconv.Atoi(s); err != nil {
				return g, err
			} else {
				nums[i] = n
			}
		}
		g.unitCount, g.unitHitPoints, g.attackDamage, g.initiative = nums[0], nums[1], nums[2], nums[3]
		g.damageType = matches[5]
		g.weaknesses, g.immunities = parseWeaknessesAndImmunities(matches[3])

		return g, nil
	}

	lines := strings.Split(input, "\n")
	inputGroups := make([]*day24group, 0)
	isInfection := false
	for _, line := range lines {
		if line == "Immune System:" {
			isInfection = false
			continue
		} else if line == "Infection:" {
			isInfection = true
			continue
		} else if line == "" {
			continue
		}
		if g, err := parseGroup(line); err != nil {
			return "", err
		} else {
			g.isInfection = isInfection
			inputGroups = append(inputGroups, &g)
		}
	}

	rec := &recorder{}
	rec.measurements = make(map[string]time.Duration)
	rec.startTimes = make(map[string]time.Time)
	rec.counters = make(map[string]int)

	for boost := 0; ; boost++ {
		groups := make([]*day24group, len(inputGroups))
		for i, g := range inputGroups {
			groupCopy := *g
			if !g.isInfection {
				groupCopy.attackDamage += boost
			}
			groupCopy.rec = rec
			groups[i] = &groupCopy
		}

		solutionFound, unitCount := func() (bool, int) {
			rec.start("cycle")
			defer rec.stop("cycle")

			for {
				// Reset
				rec.start("reset")
				newGroups := make([]*day24group, 0)
				infectionFound, immuneFound := false, false
				for _, g := range groups {
					if g.unitCount > 0 {
						g.target, g.isTargeted = nil, false
						newGroups = append(newGroups, g)
						if g.isInfection {
							infectionFound = true
						} else {
							immuneFound = true
						}
					}
				}
				rec.stop("reset")
				if !infectionFound || !immuneFound {
					if part2 && !immuneFound {
						return false, 0
					}

					totalUnitCount := 0
					for _, g := range newGroups {
						totalUnitCount += g.unitCount
					}
					return true, totalUnitCount
				}
				groups = newGroups

				// Target selection phase
				rec.start("targetsort")
				sort.Slice(groups, func(i, j int) bool { return groups[i].greaterEpOrInitiativeThan(groups[j]) })
				rec.stop("targetsort")
				totalTargets := 0
				rec.start("targetselect")
				for _, g := range groups {
					rec.start("candiselect")
					candidates := make([]*day24group, 0)
					for _, candidate := range groups {
						if g.isInfection == candidate.isInfection {
							continue
						} else if candidate.isTargeted {
							continue
						} else if _, f := candidate.immunities[g.damageType]; f {
							continue
						}
						candidates = append(candidates, candidate)
					}
					rec.stop("candiselect")
					if len(candidates) == 0 {
						rec.start("nocandi")
						rec.stop("nocandi")
						continue
					}
					rec.start("candisort")
					sort.Slice(candidates, func(i, j int) bool {
						/* If an attacking group is considering two defending groups to which it would deal equal damage, it chooses
						to target the defending group with the largest effective power; if there is still a tie, it chooses the
						defending group with the highest initiative. If it cannot deal any defending groups damage, it does not
						choose a target. Defending groups can only be chosen as a target by one attacking group.*/
						l, r := candidates[i], candidates[j]
						ldmg, rdmg := g.effectivePowerAgainst(l.weaknesses), g.effectivePowerAgainst(r.weaknesses)
						if ldmg == rdmg {
							return l.greaterEpOrInitiativeThan(r)
						}
						return ldmg > rdmg
					})
					rec.stop("candisort")
					g.target = candidates[0]
					candidates[0].isTargeted = true
					totalTargets++
				}
				rec.stop("targetselect")
				if totalTargets == 0 {
					// Couldn't assign any targets, deadlock, ignore this round
					rec.start("notarget")
					rec.stop("notarget")
					return false, 0
				}

				// Attacking phase
				rec.start("attacksort")
				sort.Slice(groups, func(i, j int) bool {
					return groups[i].initiative > groups[j].initiative
				})
				rec.stop("attacksort")

				rec.start("attacking")
				for _, g := range groups {
					if g.unitCount <= 0 || g.target == nil {
						continue
					}
					dmg := g.effectivePowerAgainst(g.target.weaknesses)
					g.target.unitCount -= dmg / g.target.unitHitPoints
				}
				rec.stop("attacking")
			}
		}()
		if solutionFound {
			keys := make([]string, 0)
			for k := range rec.measurements {
				keys = append(keys, k)
			}
			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
			for _, k := range keys {
				fmt.Printf("%s: %d @ %+v / %+v\n", k, rec.counters[k], rec.measurements[k], rec.measurements[k]/time.Duration(rec.counters[k]))
			}
			return strconv.Itoa(unitCount), nil
		}
	}
}

type recorder struct {
	measurements map[string]time.Duration
	startTimes   map[string]time.Time
	counters     map[string]int
}

func (r *recorder) start(s string) {
	_, f := r.startTimes[s]
	if f {
		panic(s + " already started")
	}
	r.startTimes[s] = time.Now()
}

func (r *recorder) stop(s string) {
	startTime, f := r.startTimes[s]
	if !f {
		panic(s + " not started")
	}
	r.record(s, startTime)
	delete(r.startTimes, s)
}

func (r *recorder) record(s string, startTime time.Time) {
	t, f := r.measurements[s]
	if !f {
		t = time.Duration(0)
	}
	r.measurements[s] = t + time.Since(startTime)
	n, f := r.counters[s]
	if !f {
		n = 0
	}
	r.counters[s] = n + 1
}

type day24group struct {
	unitCount     int
	unitHitPoints int
	weaknesses    map[string]struct{}
	immunities    map[string]struct{}
	attackDamage  int
	damageType    string
	initiative    int

	isInfection bool // false = immune system

	target     *day24group
	isTargeted bool
	rec        *recorder
}

func (d *day24group) effectivePower() int {
	return d.unitCount * d.attackDamage
}

func (d *day24group) effectivePowerAgainst(weaknesses map[string]struct{}) int {
	ep := d.effectivePower()
	if _, f := weaknesses[d.damageType]; f {
		return ep * 2
	}
	return ep
}

func (d *day24group) greaterEpOrInitiativeThan(other *day24group) bool {
	l, r := d, other
	lep, rep := l.effectivePower(), r.effectivePower()
	if lep == rep {
		return l.initiative > r.initiative
	}
	return lep > rep
}
