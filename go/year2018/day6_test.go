package year2018

import (
	"testing"
)

func Test_Day6Part1(t *testing.T) {
	basicTest(t, Day6Part1, "1, 1\n1, 6\n8, 3\n3, 4\n5, 5\n8, 9", "17")
}

func Test_Day6Part2(t *testing.T) {
	basicMultiTest(t, Day6Part2, []testInput{
		{"", ""},
	})
}
