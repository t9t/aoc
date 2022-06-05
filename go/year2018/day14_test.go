package year2018

import (
	"testing"
)

func Test_Day14Part1(t *testing.T) {
	basicMultiTest(t, Day14Part1, []testInput{
		{"9", "5158916779"},
		{"5", "0124515891"},
		{"18", "9251071085"},
		{"2018", "5941429882"},
	})
}

func Test_Day14Part2(t *testing.T) {
	basicMultiTest(t, Day14Part2, []testInput{
		{"", ""},
	})
}
