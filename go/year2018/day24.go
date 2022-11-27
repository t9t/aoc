package year2018

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
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
	parseTypes := func(s string) []string {
		if s == "" {
			return []string{}
		}
		// "[immune|weak] to radiation, bludgeoning"
		return strings.Split(strings.Split(s, " to ")[1], ", ")
	}

	parseWeaknessesAndImmunities := func(s string) ([]string, []string) {
		if s == "" {
			return []string{}, []string{}
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

	contains := func(ss []string, s string) bool {
		for _, c := range ss {
			if c == s {
				return true
			}
		}
		return false
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
			if !g.isInfection {
				//g.attackDamage += boost
			}
			inputGroups = append(inputGroups, &g)
		}
	}

	for boost := 0; ; boost++ {
		solutionFound, unitCount := func() (bool, int) {
			groups := make([]*day24group, len(inputGroups))
			for i, g := range inputGroups {
				groupCopy := *g
				if !g.isInfection {
					groupCopy.attackDamage += boost
				}
				groups[i] = &groupCopy
			}

			for {
				// Reset
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
				sort.Slice(groups, func(i, j int) bool {
					l, r := groups[i], groups[j]
					lep, rep := l.unitCount*l.attackDamage, r.unitCount*r.attackDamage
					if lep == rep {
						return l.initiative > r.initiative
					}
					return lep > rep
				})
				totalTargets := 0
				for _, g := range groups {
					candidates := make([]*day24group, 0)
					for _, candidate := range groups {
						if g.isInfection == candidate.isInfection {
							continue
						} else if candidate.isTargeted {
							continue
						} else if contains(candidate.immunities, g.damageType) {
							continue
						}
						candidates = append(candidates, candidate)
					}
					if len(candidates) == 0 {
						continue
					}
					sort.Slice(candidates, func(i, j int) bool {
						/* If an attacking group is considering two defending groups to which it would deal equal damage, it chooses
						to target the defending group with the largest effective power; if there is still a tie, it chooses the
						defending group with the highest initiative. If it cannot deal any defending groups damage, it does not
						choose a target. Defending groups can only be chosen as a target by one attacking group.*/
						l, r := candidates[i], candidates[j]
						ldmg, rdmg := g.attackDamage, g.attackDamage
						if contains(l.weaknesses, g.damageType) {
							ldmg *= 2
						}
						if contains(r.weaknesses, g.damageType) {
							rdmg *= 2
						}
						if ldmg == rdmg {
							lep, rep := l.unitCount*l.attackDamage, r.unitCount*r.attackDamage
							if lep == rep {
								return l.initiative > r.initiative
							}
							return lep > rep
						}
						return ldmg > rdmg
					})
					g.target = candidates[0]
					candidates[0].isTargeted = true
					totalTargets++
				}
				if totalTargets == 0 {
					// Couldn't assign any targets, deadlock, ignore this round
					return false, 0
				}

				// Attacking phase
				sort.Slice(groups, func(i, j int) bool {
					return groups[i].initiative > groups[j].initiative
				})

				for _, g := range groups {
					if g.unitCount <= 0 || g.target == nil {
						continue
					}
					dmg := g.unitCount * g.attackDamage
					if contains(g.target.weaknesses, g.damageType) {
						dmg *= 2
					}
					g.target.unitCount -= dmg / g.target.unitHitPoints
				}
			}
		}()
		if solutionFound {
			return strconv.Itoa(unitCount), nil
		}
	}
}

type day24group struct {
	unitCount     int
	unitHitPoints int
	weaknesses    []string
	immunities    []string
	attackDamage  int
	damageType    string
	initiative    int

	isInfection bool // false = immune system

	target     *day24group
	isTargeted bool
}
