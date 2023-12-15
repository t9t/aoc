package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(15, Day15Part1, Day15Part2)
}

func Day15Part1(input string) (string, error) {
	parts := strings.Split(strings.ReplaceAll(input, "\n", ""), ",")
	sum := 0
	for _, part := range parts {
		hash := 0
		for _, r := range part {
			hash += int(r)
			hash *= 17
			hash %= 256
		}
		sum += hash
	}

	return strconv.Itoa(sum), nil
}

func Day15Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 15 part 2 not implemented")
}
