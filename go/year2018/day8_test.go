package year2018

import (
	"testing"
)

func Test_Day8Part1(t *testing.T) {
	basicTest(t, Day8Part1, "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", "138")
}

func Test_Day8Part2(t *testing.T) {
	basicTest(t, Day8Part2, "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", "66")
}
