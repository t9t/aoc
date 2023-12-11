package year2023

import (
	"testing"
)

const day11input = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func Test_Day11Part1(t *testing.T) {
	basicMultiTest(t, Day11Part1, []testInput{
		{day11input, "374"},
	})
}

func Test_Day11Part2_10(t *testing.T) {
	basicMultiTest(t, func(input string) (string, error) {
		return day11(input, 10)
	}, []testInput{
		{day11input, "1030"},
	})
}

func Test_Day11Part2_100(t *testing.T) {
	basicMultiTest(t, func(input string) (string, error) {
		return day11(input, 100)
	}, []testInput{
		{day11input, "8410"},
	})
}

func Test_Day11Part2_1_000_000(t *testing.T) {
	basicMultiTest(t, func(input string) (string, error) {
		return day11(input, 1_000_000)
	}, []testInput{
		{day11input, "82000210"},
	})
}

func Test_Day11Part2(t *testing.T) {
	basicMultiTest(t, Day11Part2, []testInput{
		{day11input, "82000210"},
	})
}
