package year2023

import (
	"testing"
)

func Test_Day15Part1(t *testing.T) {
	basicMultiTest(t, Day15Part1, []testInput{
		{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", "1320"},
	})
}

func Test_Day15Part2(t *testing.T) {
	basicMultiTest(t, Day15Part2, []testInput{
		{"", ""},
	})
}
