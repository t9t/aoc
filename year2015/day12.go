package year2015

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day12Part1(input string) (int, error) {
	numbers, err := listNumbersIn(input)
	if err != nil {
		return 0, err
	}
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total, nil
}

var numberMatchingRegexp = regexp.MustCompile(`(-?\d+)`)

func listNumbersIn(input string) ([]int, error) {
	matches := numberMatchingRegexp.FindAllString(input, -1)
	out := make([]int, len(matches))
	for i, match := range matches {
		n, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		out[i] = n
	}
	return out, nil
}

func Day12Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
