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
	discardLower := discardUpper + diff
	kept := make([]byte, 0, len(input))
	for _, b := range []byte(input) {
		if discardUpper != 0 && (b == discardUpper || b == discardLower) {
			continue
		}

		if len(kept) == 0 {
			kept = append(kept, b)
			continue
		}

		lastPos := len(kept) - 1
		last := kept[lastPos]

		if last+diff == b || last-diff == b {
			kept = kept[:lastPos]
		} else {
			kept = append(kept, b)
		}
	}
	return len(kept)
}
