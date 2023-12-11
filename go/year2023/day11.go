package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(11, Day11Part1, Day11Part2)
}

func Day11Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	outLines := make([]string, 0)
	for _, line := range lines {
		outLines = append(outLines, line)
		if strings.Count(line, ".") == len(line) {
			outLines = append(outLines, line)
		}
	}
	lines = outLines

	outLines = make([]string, len(lines))
	for x := 0; x < len(lines[0]); x += 1 {
		dots := 0
		for y := 0; y < len(lines); y++ {
			c := lines[y][x]
			if c == '.' {
				dots += 1
			}
			outLines[y] = outLines[y] + string(c)
		}
		if dots == len(lines) {
			for y := 0; y < len(lines); y++ {
				outLines[y] = outLines[y] + "."
			}
		}
	}
	lines = outLines

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
			dx := pos2.x - pos1.x
			if dx < 0 {
				dx = -dx
			}
			dy := pos2.y - pos1.y
			if dy < 0 {
				dy = -dy
			}
			totalDist += dx + dy
		}
	}

	return strconv.Itoa(totalDist / 2), nil
}

func Day11Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 11 part 2 not implemented")
}
