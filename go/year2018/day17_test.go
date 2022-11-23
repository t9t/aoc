package year2018

import (
	"testing"
)

const day17input = `x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504`

func Test_Day17Part1(t *testing.T) {
	basicTest(t, Day17Part1, day17input, "57")
}

func Test_Day17Part2(t *testing.T) {
	basicTest(t, Day17Part2, day17input, "29")
}
