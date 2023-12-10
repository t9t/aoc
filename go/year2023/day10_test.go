package year2023

import (
	"testing"
)

func Test_Day10Part1(t *testing.T) {
	basicMultiTest(t, Day10Part1, []testInput{
		{`.....
.S-7.
.|.|.
.L-J.
.....`, "4"},
		{`-L|F7
7S-7|
L|7||
-L-J|
L|-JF`, "4"},
		{`..F7.
.FJ|.
SJ.L7
|F--J
LJ...`, "8"},
		{`7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`, "8"},
	})
}

func Test_Day10Part2(t *testing.T) {
	basicMultiTest(t, Day10Part2, []testInput{
		{"", ""},
	})
}
