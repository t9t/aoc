package year2015

import (
	"fmt"
	"strconv"
	"strings"
)

func Day10Part1(input string) (int, error) {
	current := strings.TrimSpace(input)
	for i := 0; i < 40; i++ {
		current = lookAndSay(current)
	}
	return len(current), nil
}

func lookAndSay(s string) string {
	if len(s) == 0 {
		return s
	}
	var buf strings.Builder
	var out strings.Builder
	prev := s[0]
	buf.WriteByte(prev)

	writeOut := func() {
		out.WriteString(strconv.Itoa(buf.Len()))
		out.WriteByte(prev)
	}

	for i := 1; i < len(s); i++ {
		c := s[i]
		if c != prev {
			writeOut()
			buf.Reset()
		}

		buf.WriteByte(c)
		prev = c
	}
	writeOut()
	return out.String()
}

func Day10Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
