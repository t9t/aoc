package year2018

import (
	"testing"
)

func Test_Day9Part1(t *testing.T) {
	basicMultiTest(t, Day9Part1, []testInput{
		{"9 players; last marble is worth 25 points", "32"},
		{"10 players; last marble is worth 1618 points", "8317"},
		{"13 players; last marble is worth 7999 points", "146373"},
		{"17 players; last marble is worth 1104 points", "2764"},
		{"21 players; last marble is worth 6111 points", "54718"},
		{"30 players; last marble is worth 5807 points", "37305"},
	})
}

func Test_Day9Part2(t *testing.T) {
	basicMultiTest(t, Day9Part2, []testInput{
		{"", ""},
	})
}
