package year2018

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(12, Day12Part1, Day12Part2)
}

func Day12Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	initial := strings.Split(lines[0], ": ")[1]

	rules := make(map[string]byte)
	for i := 2; i < len(lines); i++ {
		parts := strings.Split(lines[i], " => ")
		rules[parts[0]] = parts[1][0]
	}

	row := "...." + initial + "...."
	leftMost := -4
	for generation := 0; generation < 20; generation++ {
		var sb strings.Builder
		sb.WriteString("....")
		for i := 2; i < len(row)-2; i++ {
			surrounding := row[i-2 : i+3]
			c, f := rules[surrounding]
			if !f {
				c = '.'
			}
			sb.WriteByte(c)
		}
		sb.WriteString("....")
		row = sb.String()
		leftMost -= 2
	}

	sum := 0
	for i, c := range row {
		if c == '#' {
			sum += i + leftMost
		}
	}

	return strconv.Itoa(sum), nil
}

func Day12Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 12 part 2 not implemented")
}
