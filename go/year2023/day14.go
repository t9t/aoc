package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(14, Day14Part1, Day14Part2)
}

func Day14Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	for x := 0; x < len(lines[0]); x += 1 {
		for y := 0; y < len(lines); y += 1 {
			line := lines[y]
			c := line[x]
			if c != 'O' || y == 0 || lines[y-1][x] != '.' {
				continue
			}

			targety := y - 1
			for oy := y - 1; oy >= 0; oy -= 1 {
				if lines[oy][x] != '.' {
					targety = oy
				} else {
					break // O or #
				}
			}
			lineBytes := []byte(line)
			lineBytes[x] = '.'
			lines[y] = string(lineBytes)

			lineBytes = []byte(lines[targety])
			lineBytes[x] = 'O'
			lines[targety] = string(lineBytes)
		}
	}

	sum := 0
	for y, line := range lines {
		load := len(lines) - y
		for _, r := range line {
			if r == 'O' {
				sum += load
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func Day14Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 14 part 2 not implemented")
}
