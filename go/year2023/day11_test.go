package year2023

import (
	"testing"
)

func Test_Day11Part1(t *testing.T) {
	basicMultiTest(t, Day11Part1, []testInput{
		{`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`, "374"},
	})
}

func Test_Day11Part2(t *testing.T) {
	basicMultiTest(t, Day11Part2, []testInput{
		{"", ""},
	})
}
