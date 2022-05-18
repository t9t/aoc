package year2018

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(1, Day1Part1, Day1Part2)
}

func Day1Part1(input string) (string, error) {
	return day1(input, true)
}

func Day1Part2(input string) (string, error) {
	return day1(input, false)
}

func day1(input string, part1 bool) (string, error) {
	lines := strings.Split(input, "\n")
	out := 0
	seen := make(map[int]struct{})
	seen[out] = struct{}{}

	for {
		for _, line := range lines {
			n, err := strconv.Atoi(line)
			if err != nil {
				return "", fmt.Errorf("invalid line %s (%w)", line, err)
			}

			out += n
			if _, found := seen[out]; found {
				return strconv.Itoa(out), nil
			}
			seen[out] = struct{}{}
		}
		if part1 {
			return strconv.Itoa(out), nil
		}
	}
}
