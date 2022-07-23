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

		for _, u := range units {
			//fmt.Printf("Round %d unit %d: %+v\n", rounds, i, u)
			if u.hp <= 0 {
				continue
			}

			c, x, y := u.c, u.x, u.y
			target := byte('E')
			if c == 'E' {
				target = 'G'
			}

			shouldMove := true
			for _, a := range around(x, y) {
				if _, t := get(a.x, a.y); t != nil {
					if t.c == target && t.hp > 0 {
						shouldMove = false
					}
				}
			}

			if shouldMove {

				targets := make([]xAndY, 0)
				func() {
					for ty, tline := range grid {
						for tx := range tline {
							tc, tu := get(tx, ty)
							if tc == target && tu.hp > 0 {
								d := abs(ty-y) + abs(tx-x)
								if d == 1 {
									//panic("nope")
									shouldMove = false
									return
								} else {
									targets = append(targets, xAndY{x: tx, y: ty})
								}
							}
						}
					}
				}()

				sort.Slice(targets, func(i, j int) bool {
					l, r := targets[i], targets[j]
					ldx, ldy := abs(x-l.x), abs(y-l.y)
					rdx, rdy := abs(x-r.x), abs(y-r.y)
					if ldy == rdy {
						return ldx < rdx
					}
					return ldy < rdy
				})

				//fmt.Printf("  %dx%d %c; Targets: %d; Should move: %t\n", x, y, c, len(targets), shouldMove)

				if len(targets) == 0 {
					fmt.Printf("No targets exist for %dx%d %c %d\n", u.x, u.y, u.c, u.hp)
					// No targets exist
					score := 0
					for sy, line := range grid {
						for sx := range line {
							if su := findUnit(sx, sy); su != nil && su.hp > 0 {
								score += su.hp
							}
						}
					}
					drawGrid2(fmt.Sprintf("Final (round %d)", round), true)
					//fmt.Printf("bla bla %d\n", (rounds-2)*score)
					fmt.Printf("Rounds: %d; score: %d\n", round, score)
					return true, strconv.Itoa((round - 1) * score)
				}

				//shortest := math.MaxInt32
				allPathsToTargets := make([][]xAndY, 0)
				for _, t := range targets {
					//fmt.Printf("  Target %d: %+v (%d x %d)\n", ti, t, abs(x-t.x), abs(y-t.y))
					empty := emptyAround(t.x, t.y)
					for _, kek := range empty {
						tx, ty := kek.x, kek.y
						//fmt.Printf("  - Examining %d x %d\n", tx, ty)
						type xyl struct{ x, y, l int }

						lengths := map[xAndY]int{{x: tx, y: ty}: 0}
						examineNext := []xyl{{x: tx, y: ty, l: 0}}
						for len(examineNext) > 0 {
							current := examineNext[0]
							examineNext = examineNext[1:]
							cx, cy := current.x, current.y
							for _, next := range emptyAround(cx, cy) {
								nx, ny := next.x, next.y
								nxy := xAndY{x: nx, y: ny}
								if _, found := lengths[nxy]; !found {
									lengths[nxy] = current.l + 1
									examineNext = append(examineNext, xyl{x: nx, y: ny, l: current.l + 1})
								}
							}
						}
						//fmt.Printf("    %d lengths\n", len(lengths))
						//for k, v := range lengths {
						//fmt.Printf("      %dx%d = %d\n", k.x, k.y, v)
						//}

						var start xAndY
						startLen := math.MaxInt32
						for _, a := range emptyAround(x, y) {
							if length, found := lengths[a]; found {
								if length < startLen {
									startLen = length
									start = a
								}
							}
						}
						if startLen == math.MaxInt32 {
							//panic("fuck my life")
							continue
						}
						//fmt.Printf("    Start: %+v; len: %d\n", start, startLen)

						path := make([]xAndY, 0)
						path = append(path, xAndY{x: x, y: y})
						path = append(path, start)
						next := start
						if startLen != 0 {
							nextLen := startLen - 1
							for {
								done, boop := false, false
								for _, a := range emptyAround(next.x, next.y) {
									if length, found := lengths[a]; found {
										if length == nextLen {
											path = append(path, a)
											next = a
											boop = true
											nextLen = nextLen - 1
											if length == 0 {
												done = true
											}
											break
										}
									}
								}
								if done {
									break
								}
								if !boop {
									fmt.Printf("next: %v %d\n", next, nextLen)
									panic("die")
								}
							}
						}

						//fmt.Printf("    Path: %+v\n", path)
						allPathsToTargets = append(allPathsToTargets, path)
					}
				}
				if len(allPathsToTargets) == 0 {
					if debug {
						fmt.Printf("  %dx%d %c no allPathsToTargets\n", x, y, c)
					}
					continue
				}
				minPathLength := math.MaxInt32
				shortests := make([][]xAndY, 0)
				for _, path := range allPathsToTargets {
					if len(path) < minPathLength {
						minPathLength = len(path)
						shortests = [][]xAndY{path}
					} else if len(path) == minPathLength {
						shortests = append(shortests, path)
					}
				}
				step := xAndY{x: math.MaxInt32, y: math.MaxInt32}
				for _, path := range shortests {
					second := path[1]
					if second.y < step.y || (second.y == step.y && second.x < step.x) {
						step = second
					}
				}
				toc, _ := get(step.x, step.y)
				if toc != '.' {
					panic(fmt.Sprintf("no dot %dx%d %c", step.x, step.y, toc))
				}
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

		asdf := round > 46
		drawGrid2(fmt.Sprintf("After %d rounds:\n", round), asdf)
	}
}
