package year2023

import (
	"testing"
)

func Test_Day6Part1(t *testing.T) {
	basicMultiTest(t, Day6Part1, []testInput{
		{`Time:      7  15   30
Distance:  9  40  200`, "288"},
	})
}

func Test_Day6Part2(t *testing.T) {
	basicMultiTest(t, Day6Part2, []testInput{
		{"", ""},
	})
}
