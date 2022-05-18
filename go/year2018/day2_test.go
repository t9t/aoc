package year2018

import (
	"testing"
)

func Test_Day2Part1(t *testing.T) {
	basicTest(t, Day2Part1, "abcdef\nbababc\nabbcde\nabcccd\naabcdd\nabcdee\nababab", "12")
}

func Test_Day2Part2(t *testing.T) {
	basicTest(t, Day2Part2, "abcde\nfghij\nklmno\npqrst\nfguij\naxcye\nwvxyz", "fgij")
}
