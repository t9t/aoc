package year2023

import (
	"testing"
)

func Test_Day3Part1(t *testing.T) {
	basicMultiTest(t, Day3Part1, []testInput{
		{`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, "4361"},
	})
}

func Test_Day3Part2(t *testing.T) {
	basicMultiTest(t, Day3Part2, []testInput{
		{"", ""},
	})
}
