package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(10, Day10Part1, Day10Part2)
}

func Day10Part1(input string) (string, error) {
	type xy struct{ x, y int }

	lines := strings.Split(input, "\n")
	var sPos xy

findSPos:
	for y, line := range lines {
		for x, r := range line {
			if r == 'S' {
				sPos = xy{x: x, y: y}
				break findSPos
			}
		}
	}

	findConnections := func(x, y int) []xy {
		r := lines[y][x]

		var ret []xy
		if y > 0 {
			above := lines[y-1][x]
			if r == 'S' || r == '|' || r == 'L' || r == 'J' {
				if above == 'S' || above == '|' || above == '7' || above == 'F' {
					ret = append(ret, xy{x: x, y: y - 1})
				}
			}
		}
		if x < len(lines[y])-1 {
			right := lines[y][x+1]
			if r == 'S' || r == '-' || r == 'L' || r == 'F' {
				if right == 'S' || right == '-' || right == 'J' || right == '7' {
					ret = append(ret, xy{x: x + 1, y: y})
				}
			}
		}
		if y < len(lines)-1 {
			below := lines[y+1][x]
			if r == 'S' || r == '|' || r == '7' || r == 'F' {
				if below == 'S' || below == '|' || below == 'L' || below == 'J' {
					ret = append(ret, xy{x: x, y: y + 1})
				}
			}
		}
		if x > 0 {
			if r == 'S' || r == '-' || r == 'J' || r == '7' {
				if left := lines[y][x-1]; left == 'S' || left == '-' || left == 'L' || left == 'F' {
					ret = append(ret, xy{x: x - 1, y: y})
				}
			}
		}

		if len(ret) != 2 {
			panic(fmt.Sprintf("expected len 2 but is %d: %+v (input: %dx%d)", len(ret), ret, x, y))
		}
		return ret
	}

	distances := make(map[xy]int)
	for _, p := range findConnections(sPos.x, sPos.y) {
		prev := sPos
		dist := 1
		distances[p] = dist
		for {
			dist += 1
			connections := findConnections(p.x, p.y)
			var next xy
			if connections[0] == prev {
				next = connections[1]
			} else if connections[1] == prev {
				next = connections[0]
			} else {
				panic("💥")
			}

			if next == sPos {
				break
			}

			curDist, found := distances[next]
			if found {
				distances[next] = min(dist, curDist)
			} else {
				distances[next] = dist
			}

			prev = p
			p = next
		}
	}

	farthest := 0
	for _, dist := range distances {
		if dist > farthest {
			farthest = dist
		}
	}

	return strconv.Itoa(farthest), nil
}

func Day10Part2(input string) (string, error) {
	type xy struct{ x, y int }

	lines := strings.Split(input, "\n")
	var sPos xy

findSPos:
	for y, line := range lines {
		for x, r := range line {
			if r == 'S' {
				sPos = xy{x: x, y: y}
				break findSPos
			}
		}
	}

	findConnections := func(x, y int) []xy {
		r := lines[y][x]

		var ret []xy
		if y > 0 {
			above := lines[y-1][x]
			if r == 'S' || r == '|' || r == 'L' || r == 'J' {
				if above == 'S' || above == '|' || above == '7' || above == 'F' {
					ret = append(ret, xy{x: x, y: y - 1})
				}
			}
		}
		if x < len(lines[y])-1 {
			right := lines[y][x+1]
			if r == 'S' || r == '-' || r == 'L' || r == 'F' {
				if right == 'S' || right == '-' || right == 'J' || right == '7' {
					ret = append(ret, xy{x: x + 1, y: y})
				}
			}
		}
		if y < len(lines)-1 {
			below := lines[y+1][x]
			if r == 'S' || r == '|' || r == '7' || r == 'F' {
				if below == 'S' || below == '|' || below == 'L' || below == 'J' {
					ret = append(ret, xy{x: x, y: y + 1})
				}
			}
		}
		if x > 0 {
			if r == 'S' || r == '-' || r == 'J' || r == '7' {
				if left := lines[y][x-1]; left == 'S' || left == '-' || left == 'L' || left == 'F' {
					ret = append(ret, xy{x: x - 1, y: y})
				}
			}
		}

		if len(ret) != 2 {
			panic(fmt.Sprintf("expected len 2 but is %d: %+v (input: %dx%d)", len(ret), ret, x, y))
		}
		return ret
	}

	distances := make(map[xy]int)
	distances[sPos] = 0
	sConnections := findConnections(sPos.x, sPos.y)
	for _, p := range sConnections {
		prev := sPos
		dist := 1
		distances[p] = dist
		for {
			dist += 1
			connections := findConnections(p.x, p.y)
			var next xy
			if connections[0] == prev {
				next = connections[1]
			} else if connections[1] == prev {
				next = connections[0]
			} else {
				panic("💥")
			}

			if next == sPos {
				break
			}

			curDist, found := distances[next]
			if found {
				distances[next] = min(dist, curDist)
			} else {
				distances[next] = dist
			}

			prev = p
			p = next
		}
	}

	sConn1, sConn2 := sConnections[0], sConnections[1]
	s1dx, s1dy := sConn1.x-sPos.x, sConn1.y-sPos.y
	s2dx, s2dy := sConn2.x-sPos.x, sConn2.y-sPos.y
	var s rune
	if s1dx == 0 && s1dy == -1 {
		// above
		if s2dx == 1 && s2dy == 0 {
			// right
			s = 'L'
		} else if s2dx == 0 && s2dy == 1 {
			// below
			s = '|'
		} else if s2dx == -1 && s2dy == 0 {
			// left
			s = 'J'
		} else {
			panic("above not connecting")
		}
	} else if s1dx == 1 && s1dy == 0 {
		// right
		if s2dx == 0 && s2dy == 1 {
			// below
			s = 'F'
		} else if s2dx == -1 && s2dy == 0 {
			// left
			s = '-'
		} else {
			panic("right not connecting")
		}
	} else if s1dx == 0 && s1dy == 1 {
		// below
		if s2dx == -1 && s2dy == 0 {
			// left
			s = '7'
		} else {
			panic("below not connecting")
		}
	} else {
		panic("unable to determine S")
	}

	count := 0
	for y, line := range lines {
		var prevCorner rune = 0
		inside := false
		for x, r := range line {
			_, pipe := distances[xy{x: x, y: y}]
			if pipe {
				if r == 'S' {
					r = s
				}
				if r == '|' {
					inside = !inside
				} else if r == '-' {
					// nothing to do
				} else {
					// corner
					if r != 'L' && r != 'J' && r != '7' && r != 'F' {
						panic(fmt.Sprintf("expected corner but got: %c", r))
					} else if prevCorner != 0 {
						if prevCorner == 'L' {
							if r == '7' {
								inside = !inside
							} else if r == 'J' {
								// nothing to do
							} else {
								panic("unexpected corner situation 1")
							}
						} else if prevCorner == 'F' {
							if r == 'J' {
								inside = !inside
							} else if r == '7' {
								// nothing to do
							} else {
								panic("unexpected corner situation 2")
							}
						} else {
							panic("unexpected corner situation 3")
						}

						prevCorner = 0
					} else {
						prevCorner = r
					}
				}
			} else if inside {
				count += 1
			} else {
				// Outside
			}
		}
	}

	return strconv.Itoa(count), nil
}
