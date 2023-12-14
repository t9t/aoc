package year2023

import (
	"testing"
)

func Test_Day14Part1(t *testing.T) {
	basicMultiTest(t, Day14Part1, []testInput{
		{`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`, "136"},
	})
}

func Test_Day14Part2(t *testing.T) {
	basicMultiTest(t, Day14Part2, []testInput{
		{"", ""},
	})
}
