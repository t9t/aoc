package year2023

import (
	"testing"
)

const day9input = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func Test_Day9Part1(t *testing.T) {
	basicMultiTest(t, Day9Part1, []testInput{
		{day9input, "114"},
	})
}

func Test_Day9Part2(t *testing.T) {
	basicMultiTest(t, Day9Part2, []testInput{
		{day9input, "2"},
	})
}
