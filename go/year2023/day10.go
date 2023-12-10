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
	return "", fmt.Errorf("Day 10 part 2 not implemented")
}
