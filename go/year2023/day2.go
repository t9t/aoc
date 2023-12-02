package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(2, Day2Part1, Day2Part2)
}

func Day2Part1(input string) (string, error) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		left, right, found := strings.Cut(line, ": ")
		if !found {
			return "", fmt.Errorf("invalid line (could not find :) %s", line)
		}

		samples := strings.Split(right, "; ")
		allPossible := true
		for _, sample := range samples {
			cubes := strings.Split(sample, ", ")
			possible := true
			for _, cube := range cubes {
				countString, color, found := strings.Cut(cube, " ")
				if !found {
					return "", fmt.Errorf("invalid line (could not split count/color) %s", line)
				}
				count, err := strconv.Atoi(countString)
				if err != nil {
					return "", err
				}

				if color == "red" && count > 12 {
					possible = false
					break
				} else if color == "green" && count > 13 {
					possible = false
					break
				} else if color == "blue" && count > 14 {
					possible = false
					break
				}
			}

			if !possible {
				allPossible = false
				break
			}
		}

		if allPossible {
			_, gameNumString, found := strings.Cut(left, " ")
			if !found {
				return "", fmt.Errorf("invalid line (could not split game number) %s", line)
			}
			gameNum, err := strconv.Atoi(gameNumString)
			if err != nil {
				return "", err
			}
			sum += gameNum
		}
	}
	return strconv.Itoa(sum), nil
}

func Day2Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 2 part 2 not implemented")
}
