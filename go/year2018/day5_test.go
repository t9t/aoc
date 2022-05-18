package year2018

import (
	"testing"
)

func Test_Day5Part1(t *testing.T) {
	basicMultiTest(t, Day5Part1, []testInput{
		{"aA", "0"},
		{"abBA", "0"},
		{"abAB", "4"},
		{"aabAAB", "6"},
		{"dabCBAcaDA", "10"},
		{"dabCBAcCcaDA", "10"},
		{"dabAaCBAcCcaDA", "10"},
		{"dabAcCaCBAcCcaDA", "10"},
	})
}

func Test_Day5Part2(t *testing.T) {
	basicMultiTest(t, Day5Part2, []testInput{
		{"", ""},
	})
}
