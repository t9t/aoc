package year2018

import (
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(5, Day5Part1, Day5Part2)
}

func Day5Part1(input string) (string, error) {
	return strconv.Itoa(reduce(input, 0)), nil
}

func Day5Part2(input string) (string, error) {
	diff := byte('a') - byte('A')
	_ = func(upper byte) string {
		sb := strings.Builder{}
		lower := upper + diff
		for _, c := range []byte(input) {
			if c != upper && c != lower {
				sb.WriteByte(c)
			}
		}
		return sb.String()
	}

	shortest := len(input)
	for b := byte('A'); b <= byte('Z'); b++ {
		l := reduce(input, b)
		if l < shortest {
			shortest = l
		}
	}
	return strconv.Itoa(shortest), nil
}

func reduce(input string, discardUpper byte) int {
	diff := byte('a') - byte('A')
	discardLower := byte(0)
	if discardUpper != 0 {
		discardLower = discardUpper + diff
	}
	kept := make([]byte, len(input))
	ptr := 0
	for _, b := range []byte(input) {
		if b == discardUpper || b == discardLower {
			continue
		}

		if ptr == 0 {
			kept[ptr] = b
			ptr++
			continue
		}

		lastPos := ptr - 1
		last := kept[lastPos]

		if last+diff == b || last-diff == b {
			ptr--
		} else {
			kept[ptr] = b
			ptr++
		}
	}
	return ptr
}
