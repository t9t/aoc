package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(6, Day6Part1, Day6Part2)
}

func Day6Part1(input string) (string, error) {
	re := regexp.MustCompile(`(?m)(\d+)`)
	toNumbers := func(s string) ([]int, error) {
		var ret []int
		for _, m := range re.FindAllString(s, -1) {
			n, err := strconv.Atoi(m)
			if err != nil {
				return nil, err
			}
			ret = append(ret, n)
		}
		return ret, nil
	}

	return day6(input, toNumbers)
}

func Day6Part2(input string) (string, error) {
	parseLine := func(s string) ([]int, error) {
		n, err := strconv.Atoi(strings.ReplaceAll(strings.Split(s, ":")[1], " ", ""))
		if err != nil {
			return nil, err
		}
		return []int{n}, nil
	}

	return day6(input, parseLine)
}

func day6(input string, parseFunc func(string) ([]int, error)) (string, error) {
	lines := strings.Split(input, "\n")
	if len(lines) != 2 {
		return "", fmt.Errorf("invalid input (expected 2 lines but got %d): %s", len(lines), input)
	}
	times, err := parseFunc(lines[0])
	if err != nil {
		return "", fmt.Errorf("invalid Times line %s: %w", lines[0], err)
	}
	distances, err := parseFunc(lines[1])
	if err != nil {
		return "", fmt.Errorf("invalid Distances line %s: %w", lines[1], err)
	}

	mul := 1
	for i, time := range times {
		dist := distances[i]

		beats := 0 // sponsored by Dwight Shrute
		for hold := 1; hold < time; hold++ {
			moveTime := time - hold
			moveDist := moveTime * hold
			if moveDist > dist {
				beats++
			}
		}
		mul *= beats
	}
	return strconv.Itoa(mul), nil
}
