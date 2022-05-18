package year2018

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(2, Day2Part1, Day2Part2)
}

func Day2Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	twos, threes := 0, 0
	for _, line := range lines {
		counts := make(map[rune]int)
		for _, c := range line {
			if n, found := counts[c]; found {
				counts[c] = n + 1
			} else {
				counts[c] = 1
			}
		}
		incrementedTwos, incrementedThrees := false, false
		for _, n := range counts {
			if n == 2 && !incrementedTwos {
				twos++
				incrementedTwos = true
			}
			if n == 3 && !incrementedThrees {
				threes++
				incrementedThrees = true
			}
		}
	}

	return strconv.Itoa(twos * threes), nil
}

func Day2Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	for _, left := range lines {
		for _, right := range lines {
			if left == right {
				continue
			}

			differences, common := 0, strings.Builder{}
			for i := range left {
				c1, c2 := left[i], right[i]
				if c1 == c2 {
					common.WriteByte(c1)
				} else {
					differences++
					if differences > 1 {
						break
					}
				}
			}
			if differences == 1 {
				return common.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no answer found")
}
