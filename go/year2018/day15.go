package year2018

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(15, Day15Part1, Day15Part2)
}

func Day15Part1(input string) (string, error) {
	_, outcome := day15Part2(input, 3, false)
	return outcome, nil
}

func Day15Part2(input string) (string, error) {
	for i := 4; i < math.MaxInt32; i++ {
		noElfDied, outcome := day15Part2(input, i, true)
		if noElfDied {
			return outcome, nil
		}
	}
	return "", fmt.Errorf("no answer found")
}

func day15Part2(input string, elfAttackDamge int, elfDeathIsLoss bool) (bool, string) {
	type unit struct {
		c        byte
		hp, x, y int
	}
	type xAndY struct{ x, y int }

	walls := make([][]bool, 0)
	units := make([]*unit, 0)

	// Parse input into walls and units
	for y, line := range strings.Split(input, "\n") {
		row := make([]bool, 0)
		for x, c := range []byte(line) {
			if c == 'G' || c == 'E' {
				units = append(units, &unit{c: c, hp: 200, x: x, y: y})
			}
			row = append(row, c == '#')
		}
		walls = append(walls, row)
	}

	maxX, maxY := len(walls[0])-1, len(walls)-1
	deltas := []xAndY{{x: 0, y: -1}, {x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: 1}}

	// Helper functions
	findUnit := func(x, y int) *unit {
		for _, unit := range units {
			if unit.x == x && unit.y == y && unit.hp > 0 {
				return unit
			}
		}
		return nil
	}

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	emptyAround := func(x, y int) []xAndY {
		arounds := make([]xAndY, 0)
		for _, d := range deltas {
			cx, cy := x+d.x, y+d.y
			if cx >= 0 && cx <= maxX && cy >= 0 && cy <= maxY {
				if !walls[cy][cx] {
					if unit := findUnit(cx, cy); unit == nil || unit.hp <= 0 {
						arounds = append(arounds, xAndY{x: cx, y: cy})
					}
				}
			}
		}
		return arounds
	}

	buildDistanceMap := func(startX, startY int) map[xAndY]int {
		type xyd struct{ x, y, d int }
		distances := make(map[xAndY]int)
		queue := []xyd{{x: startX, y: startY, d: 0}}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			for _, a := range emptyAround(current.x, current.y) {
				if _, found := distances[a]; !found {
					distances[a] = current.d + 1
					queue = append(queue, xyd{x: a.x, y: a.y, d: current.d + 1})
				}
			}
		}
		return distances
	}

	// Battle
	round := 0
	for {
		round++

		// Sort units top-to-bottom then left-to-right
		sort.Slice(units, func(i, j int) bool {
			l, r := units[i], units[j]
			if l.y == r.y {
				return l.x < r.x
			}
			return l.y < r.y
		})

		for _, current := range units {
			if current.hp <= 0 {
				continue
			}

			// Find any targetable units
			targetUnits := make([]*unit, 0)
			shouldMove := true
			for _, other := range units {
				if other.c != current.c && other.hp > 0 {
					if abs(other.x-current.x)+abs(other.y-current.y) == 1 {
						// Directly adjacent, we should attack, not move
						shouldMove = false
						break
					}
					targetUnits = append(targetUnits, other)
				}
			}

			if shouldMove {
				if len(targetUnits) == 0 {
					// There are no units to target, the game is over, calcuate & return the score
					score := 0
					for _, unit := range units {
						if unit.hp > 0 {
							score += unit.hp
						}
					}
					return true, strconv.Itoa((round - 1) * score)
				}

				distances := buildDistanceMap(current.x, current.y)

				// Find the position next to a target that is the shortest distance away
				shortestDistance := math.MaxInt32
				var targetPos xAndY
				for _, targetUnit := range targetUnits {
					for _, pos := range emptyAround(targetUnit.x, targetUnit.y) {
						if dist, found := distances[pos]; found {
							if dist < shortestDistance {
								shortestDistance, targetPos = dist, pos
							}
						}
					}
				}

				if shortestDistance == math.MaxInt32 {
					// No target position is reachable, so we don't move
					continue
				}

				targetDistances := buildDistanceMap(targetPos.x, targetPos.y)
				var step xAndY
				if abs(targetPos.x-current.x)+abs(targetPos.y-current.y) == 1 {
					// We are directly adjacent to the target position, simply move there
					step = targetPos
				} else {
					// Find the empty spot next to the current unit that is the shortest distance away from the target position
					lowestDist := math.MaxInt32
					for _, pos := range emptyAround(current.x, current.y) {
						if dist, found := targetDistances[pos]; found {
							if dist < lowestDist {
								lowestDist, step = dist, pos
							}
						}
					}
				}

				// Move to the chosen spot
				current.x, current.y = step.x, step.y
			} // if shouldMove

			// Find the enemy unit next to the current unit with the lowest HP, if any
			var toAttack *unit
			for _, delta := range deltas {
				otherUnit := findUnit(current.x+delta.x, current.y+delta.y)
				if otherUnit == nil || otherUnit.c == current.c || otherUnit.hp <= 0 {
					// No unit there, or not an enemy, or already dead
					continue
				}
				if toAttack == nil {
					toAttack = otherUnit
				} else if otherUnit.hp < toAttack.hp {
					toAttack = otherUnit
				}
			}

			// If an adjacent enemy is found, attack it
			if toAttack != nil {
				dmg := 3
				if current.c == 'E' {
					dmg = elfAttackDamge
				}
				toAttack.hp -= dmg
				if elfDeathIsLoss && toAttack.c == 'E' && toAttack.hp <= 0 {
					// RIP elf
					return false, ""
				}
			}
		}
	}
}
