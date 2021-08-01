package year2015

import (
	"fmt"
	"strings"
)

func Day1Part1(input string) (int, error) {
	input = strings.TrimSpace(input)

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
