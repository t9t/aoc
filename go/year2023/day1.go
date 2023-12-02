package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(1, Day1Part1, Day1Part2)
}

func Day1Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		first, last := 0, 0
		for _, c := range line {
			if c >= 49 && c <= 57 {
				last = int(c - 48)
				if first == 0 {
					first = last
				}
			}
		}
		sum += first*10 + last
	}

	return strconv.Itoa(sum), nil
}

func Day1Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 1 part 2 not implemented")
}
