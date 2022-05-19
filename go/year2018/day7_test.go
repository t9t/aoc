package year2018

import (
	"testing"
)

var day7input = `
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`

func Test_Day7Part1(t *testing.T) {
	basicTest(t, Day7Part1, day7input, "CABDFE")
}

func Test_Day7Part2(t *testing.T) {
	basicTest(t, func(input string) (string, error) {
		return day7Part2(input, 2, 0)
	}, day7input, "15")
}
