package year2023

import (
	"testing"
)

const day16input = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func Test_Day16Part1(t *testing.T) {
	basicMultiTest(t, Day16Part1, []testInput{
		{day16input, "46"},
	})
}

func Test_Day16Part2(t *testing.T) {
	basicMultiTest(t, Day16Part2, []testInput{
		{day16input, "51"},
	})
}
