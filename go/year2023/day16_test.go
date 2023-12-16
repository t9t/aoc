package year2023

import (
	"testing"
)

func Test_Day16Part1(t *testing.T) {
	basicMultiTest(t, Day16Part1, []testInput{
		{`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`, "46"},
	})
}

func Test_Day16Part2(t *testing.T) {
	basicMultiTest(t, Day16Part2, []testInput{
		{"", ""},
	})
}
