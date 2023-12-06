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
	lines := strings.Split(input, "\n")
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

	times, err := toNumbers(lines[0])
	if err != nil {
		return "", fmt.Errorf("invalid line %s: %w\n", lines[0], err)
	}
	distances, err := toNumbers(lines[1])
	if err != nil {
		return "", fmt.Errorf("invalid line %s: %w\n", lines[1], err)
	}

	return day6(times, distances)
}

func Day6Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	parseLine := func(s string) (int, error) {
		return strconv.Atoi(strings.ReplaceAll(strings.Split(s, ":")[1], " ", ""))
	}

	times, err := parseLine(lines[0])
	if err != nil {
		return "", fmt.Errorf("invalid line %s: %w\n", lines[0], err)
	}
	distances, err := parseLine(lines[1])
	if err != nil {
		return "", fmt.Errorf("invalid line %s: %w\n", lines[1], err)
	}

	return day6([]int{times}, []int{distances})
}

func day6(times, distances []int) (string, error) {
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
