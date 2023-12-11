package year2023

import (
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(11, Day11Part1, Day11Part2)
}

func Day11Part1(input string) (string, error) {
	return day11(input, 2)
}

func Day11Part2(input string) (string, error) {
	return day11(input, 1_000_000)
}

func day11(input string, expansion int) (string, error) {
	lines := strings.Split(input, "\n")
	emptyRows := make(map[int]struct{})
	for y, line := range lines {
		if strings.Count(line, ".") == len(line) {
			emptyRows[y] = struct{}{}
		}
	}

	emptyCols := make(map[int]struct{})
	for x := 0; x < len(lines[0]); x += 1 {
		dots := 0
		for y := 0; y < len(lines); y++ {
			c := lines[y][x]
			if c == '.' {
				dots += 1
			} else {
				break
			}
		}
		if dots == len(lines) {
			emptyCols[x] = struct{}{}
		}
	}

	type xy struct{ x, y int }
	galaxies := make(map[xy]struct{})
	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				galaxies[xy{x: x, y: y}] = struct{}{}
			}
		}
	}

	totalDist := 0
	for pos1 := range galaxies {
		for pos2 := range galaxies {
			if pos1 == pos2 {
				continue
			}
			dx := 0
			for x := min(pos1.x, pos2.x); x < max(pos1.x, pos2.x); x += 1 {
				if _, f := emptyCols[x]; f {
					dx += expansion
				} else {
					dx += 1
				}
			}

			dy := 0
			for y := min(pos1.y, pos2.y); y < max(pos1.y, pos2.y); y += 1 {
				if _, f := emptyRows[y]; f {
					dy += expansion
				} else {
					dy += 1
				}
			}
			totalDist += dx + dy
		}
	}

	return strconv.Itoa(totalDist / 2), nil
}
