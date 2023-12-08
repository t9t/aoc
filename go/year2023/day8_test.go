package year2023

import (
	"testing"
)

func Test_Day8Part1(t *testing.T) {
	basicMultiTest(t, Day8Part1, []testInput{
		{`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`, "2"},
		{`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`, "6"},
	})
}

func Test_Day8Part2(t *testing.T) {
	basicMultiTest(t, Day8Part2, []testInput{
		{"", ""},
	})
}
