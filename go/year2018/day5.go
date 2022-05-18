package year2018

import (
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(5, Day5Part1, Day5Part2)
}

func Day5Part1(input string) (string, error) {
	return strconv.Itoa(reduce(input)), nil
}

func Day5Part2(input string) (string, error) {
	diff := byte('a') - byte('A')
	inputWithout := func(upper byte) string {
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
		l := reduce(inputWithout(b))
		if l < shortest {
			shortest = l
		}
	}
	return strconv.Itoa(shortest), nil
}

func reduce(input string) int {
	diff := byte('a') - byte('A')
	for {
		anyReduced := false
		for i := 0; i < len(input)-1; i++ {
			left := input[i]
			right := input[i+1]

			if left+diff == right || left-diff == right {
				input = input[:i] + input[i+2:]
				anyReduced = true
				i++
			}
		}
		if !anyReduced {
			break
		}
	}
	return len(input)
}
