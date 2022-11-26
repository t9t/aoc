package year2018

import (
	"testing"
)

func Test_Day19Part1(t *testing.T) {
	input := `#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`

	basicTest(t, Day19Part1, input, "6")
}

func Test_Day19Part2(t *testing.T) {
	basicMultiTest(t, Day19Part2, []testInput{
		{"", ""},
	})
}
