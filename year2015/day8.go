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
		length++
		sub := v[i:]
		if strings.HasPrefix(sub, `\x`) {
			// \xNN escaping, increment i by 1+3 = 4 to move beyond escape sequence
			i += 3
		} else if strings.HasPrefix(sub, `\\`) || strings.HasPrefix(sub, `\"`) {
			// \\ or \", increment i by 1 to move to next character
			i += 1
		}
	}
	return length
}

func Day8Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
