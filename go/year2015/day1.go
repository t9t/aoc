package year2015

import (
	"fmt"
)

func Day1Part1(input string) (int, error) {
	floor := 0
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		} else {
			return 0, fmt.Errorf("invalid input, unexpected character %s", string(c))
		}
	}
	return floor, nil
}

func Day1Part2(input string) (int, error) {
	floor := 0
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		} else {
			return 0, fmt.Errorf("invalid input, unexpected character %s", string(c))
		}
		if floor == -1 {
			return i + 1, nil
		}
	}
	return 0, fmt.Errorf("invalid input, never entered the basement")
}
