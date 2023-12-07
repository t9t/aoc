package year2023

import (
	"testing"
)

func Test_Day7Part1(t *testing.T) {
	basicMultiTest(t, Day7Part1, []testInput{
		{`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`, "6440"},
	})
}

func Test_Day7Part2(t *testing.T) {
	basicMultiTest(t, Day7Part2, []testInput{
		{"", ""},
	})
}
