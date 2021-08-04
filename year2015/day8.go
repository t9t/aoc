package year2015

import (
	"fmt"
	"strings"
)

func Day8Part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	n := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		n += len(line)
		n -= characterLength(line)
	}
	return n, nil
}

func characterLength(inputString string) int {
	v := inputString[1 : len(inputString)-1]
	length := 0
	for i := 0; i < len(v); i++ {
		c := v[i]
		length++
		if c != '\\' {
			// regular character, just continue to the next
			continue
		}

		i++
		next := v[i]
		if next == '\\' || next == '"' {
			// escaping a \ or " which we counted, so continue on
			continue
		} else if next != 'x' {
			// not escaping anything, also count this character and move on
			length++
			continue
		}

		// hex char escaping; we counted the character already, so just increment i by 2 more
		i += 2
	}
	return length
}

func Day8Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
