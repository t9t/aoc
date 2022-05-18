package year2018

import (
	"testing"
)

func Test_Day1Part1(t *testing.T) {
	basicMultiTest(t, Day1Part1, []testInput{
		{"+1\n-2\n+3\n+1", "3"},
		{"+1\n+1\n+1", "3"},
		{"+1\n+1\n-2", "0"},
		{"-1\n-2\n-3", "-6"},
	})
}

func Test_Day1Part2(t *testing.T) {
	basicMultiTest(t, Day1Part1, []testInput{
		{"+1\n-2\n+3\n+1", "3"},
		{"+1\n+1\n+1", "3"},
		{"+1\n+1\n-2", "0"},
		{"-1\n-2\n-3", "-6"},
	})
}
