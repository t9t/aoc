package year2018

import (
	"testing"
)

func Test_Day7Part1(t *testing.T) {
	input := `
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
	`
	basicTest(t, Day7Part1, input, "CABDFE")
}

func Test_Day7Part2(t *testing.T) {
	basicMultiTest(t, Day7Part2, []testInput{
		{"", ""},
	})
}
