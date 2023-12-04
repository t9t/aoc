package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(4, Day4Part1, Day4Part2)
}

func Day4Part1(input string) (string, error) {
	totalPoints := 0
	for _, line := range strings.Split(input, "\n") {
		_, lists, found := strings.Cut(line, ": ")
		if !found {
			return "", fmt.Errorf("invalid line (no ': '): %s", line)
		}

		winning, owned, found := strings.Cut(lists, " | ")
		if !found {
			return "", fmt.Errorf("invalid line (no ' | '): %s", line)
		}

		toNumbers := func(s string) (map[int]struct{}, error) {
			numbers := make(map[int]struct{})
			for _, item := range strings.Split(strings.TrimSpace(s), " ") {
				if item == "" {
					continue
				}
				n, err := strconv.Atoi(strings.TrimSpace(item))
				if err != nil {
					return nil, err
				}
				numbers[n] = struct{}{}
			}
			return numbers, nil
		}

		winningNumbers, err := toNumbers(winning)
		if err != nil {
			return "", err
		}

		ownedNumbers, err := toNumbers(owned)
		if err != nil {
			return "", err
		}

		score := 0
		for n := range ownedNumbers {
			if _, found := winningNumbers[n]; found {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		totalPoints += score
	}
	return strconv.Itoa(totalPoints), nil
}

func Day4Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 4 part 2 not implemented")
}
