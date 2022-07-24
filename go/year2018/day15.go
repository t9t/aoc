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
	_, r := day15Part2(input, 3, false)
	return r, nil
}

func Day15Part2(input string) (string, error) {
	i := 4
	for {
		//fmt.Printf("trying ap %d\n", i)
		ok, r := day15Part2(input, i, true)
		if ok {
			return r, nil
		}
		i++
		if i >= 50 {
			panic("kaput")
		}
	}
}

func day15Part2(input string, ap int, elfDeathIsLoss bool) (bool, string) {
	debug := false
	type unit struct {
		c    byte
		hp   int
		x, y int
	}
	type xAndY struct{ x, y int }

	grid := make([][]byte, 0)
	units := make([]*unit, 0)

	findUnit := func(x, y int) *unit {
		for _, u := range units {
			if u.x == x && u.y == y && u.hp > 0 {
				return u
			}
		}
		return nil
	}

	drawGrid2 := func(caption string, do bool) {
		if !debug || !do {
			return
		}
		fmt.Println(caption)
		for y, line := range grid {
			for x, c := range line {
				if u := findUnit(x, y); u != nil {
					c = u.c
				}
				fmt.Printf("%c", c)
			}

			unitsOnLine := make([]*unit, 0)
			for _, u := range units {
				if u.y == y {
					unitsOnLine = append(unitsOnLine, u)
				}
			}
			sort.Slice(unitsOnLine, func(i, j int) bool {
				l, r := unitsOnLine[i], unitsOnLine[j]
				return l.x < r.x
			})

			fmt.Printf("   ")

			for n, u := range unitsOnLine {
				if n > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%c(%d)", u.c, u.hp)
			}

			fmt.Println()
		}
	}

	drawGrid := func(caption string) {
		drawGrid2(caption, debug)
	}

	for y, line := range strings.Split(input, "\n") {
		row := make([]byte, 0)
		for x, c := range []byte(line) {
			if c == 'G' || c == 'E' {
				units = append(units, &unit{c: c, hp: 200, x: x, y: y})
				c = '.'
			}
			row = append(row, c)
		}
		grid = append(grid, row)
	}

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	maxX, maxY := len(grid[0])-1, len(grid)-1
	invalid := byte(0)

	get := func(x, y int) (byte, *unit) {
		if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
			if u := findUnit(x, y); u != nil && u.hp > 0 {
				return u.c, u
			} else {
				return grid[y][x], nil
			}
		}
		return invalid, nil
	}

	around := func(x, y int) []xAndY {
		return []xAndY{{x: x, y: y - 1}, {x: x - 1, y: y}, {x: x + 1, y: y}, {x: x, y: y + 1}}
	}

	emptyAround := func(x, y int) []xAndY {
		arounds := make([]xAndY, 0)
		for _, candidate := range around(x, y) {
			cx, cy := candidate.x, candidate.y
			if c, _ := get(cx, cy); c == '.' {
				arounds = append(arounds, candidate)
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
			cx, cy, cd := current.x, current.y, current.d

			for _, a := range emptyAround(cx, cy) {
				if _, ok := distances[a]; ok {
					continue
				}
				distances[a] = cd + 1
				queue = append(queue, xyd{x: a.x, y: a.y, d: cd + 1})
			}
		}
		return distances
	}

	drawGrid("Initially:")

	round := 0
	for {
		round++

		sort.Slice(units, func(i, j int) bool {
			l, r := units[i], units[j]
			if l.y == r.y {
				return l.x < r.x
			}
			return l.y < r.y
		})

		for i, u := range units {
			func(n int) {}(i)
			//fmt.Printf("Round %d unit %d: %c %+v\n", round, i, u.c, u)
			if u.hp <= 0 {
				continue
			}

			c, x, y := u.c, u.x, u.y
			target := byte('E')
			if c == 'E' {
				target = 'G'
			}

			//fmt.Printf("a\n")
			targetUnits := make([]*unit, 0)
			shouldMove := func() bool {
				for _, ou := range units {
					if ou.c == target && ou.hp > 0 {
						if abs(ou.x-x)+abs(ou.y-y) == 1 {
							// Directly adjacent, we should attack
							return false
						}
						targetUnits = append(targetUnits, ou)
					}
				}
				return true
			}()

			//fmt.Printf("b\n")
			if shouldMove {
				//fmt.Printf("c\n")
				if len(targetUnits) == 0 {
					//fmt.Printf("d\n")
					score := 0
					for _, su := range units {
						if su.hp > 0 {
							score += su.hp
						}
					}
					fmt.Printf("round: %d; score: %d\n", round, score)
					return true, strconv.Itoa((round - 1) * score)
				}

				targetPositions := make([]xAndY, 0)
				for _, tu := range targetUnits {
					for _, a := range emptyAround(tu.x, tu.y) {
						targetPositions = append(targetPositions, a)
					}
				}

				//fmt.Printf("e\n")

				distances := buildDistanceMap(x, y)

				//fmt.Printf("f targetPositions: %+v\n", targetPositions)
				//fmt.Printf("distances: %+v\n", distances)

				shortestDistance := math.MaxInt32
				var targetPos xAndY
				for _, pos := range targetPositions {
					if dist, found := distances[pos]; found {
						if dist < shortestDistance {
							shortestDistance = dist
							targetPos = pos
						}
					}
				}
				//fmt.Printf("g\n")
				//fmt.Printf("shortest: %d; targetPos: %+v; distances: %+v\n", shortestDistance, targetPos, distances)
				if shortestDistance == math.MaxInt32 {
					//drawGrid2("Oh no?", true)
					//panic("impossible game state")
					//fmt.Printf("Cannea move\n")
					continue
				}

				targetDistances := buildDistanceMap(targetPos.x, targetPos.y)
				//drawGrid("...")
				//fmt.Printf("targetDistances %+v\n", targetDistances)
				var step xAndY
				if abs(targetPos.x-u.x)+abs(targetPos.y-u.y) == 1 {
					step = targetPos
				} else {
					step = func() xAndY {
						lowestDist := math.MaxInt32
						var lowest xAndY
						for _, a := range emptyAround(u.x, u.y) {
							if dist, found := targetDistances[a]; found {
								if dist < lowestDist {
									lowestDist = dist
									lowest = a
								}
							}
						}
						if lowestDist == math.MaxInt32 {
							panic(":(")
						}
						return lowest
					}()
				}

				//fmt.Printf("i\n")

				u.x = step.x
				u.y = step.y
			}

			var toAttack *unit
			for _, other := range around(u.x, u.y) {
				ox, oy := other.x, other.y
				ou := findUnit(ox, oy)
				if ou == nil {
					continue
				}
				if ou.c != target {
					continue
				}
				if ou.hp <= 0 {
					continue
				}
				if toAttack == nil {
					toAttack = ou
					continue
				}
				if ou.hp < toAttack.hp {
					toAttack = ou
				}
			}
			if toAttack != nil {
				dmg := 3
				if toAttack.c == 'G' {
					dmg = ap
				}
				toAttack.hp -= dmg
				if elfDeathIsLoss {
					if toAttack.c == 'E' && toAttack.hp <= 0 {
						drawGrid2(fmt.Sprintf("Elf died (round %d)", round), true)
						// elf died
						return false, ""
					}
				}
				continue
			}
			///
		}

		//asdf := round > 46
		drawGrid2(fmt.Sprintf("After %d rounds:\n", round), false)
	}
}
