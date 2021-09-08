package year2015

import (
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

func Day8Part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	n := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		n += countEncodedStringLength(line)
		n -= len(line)
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

func countEncodedStringLength(s string) int {
	length := 2 // starting and ending quotes
	for i := 0; i < len(s); i++ {
		c := s[i]
		length++
		if c == '\\' || c == '"' {
			length++
		}
	}
	return length
}
