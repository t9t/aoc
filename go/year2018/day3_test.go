package year2018

import (
	"testing"
)

func Test_Day3Part1(t *testing.T) {
	basicTest(t, Day3Part1, "#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2", "4")
}

func Test_Day3Part2(t *testing.T) {
	basicTest(t, Day3Part2, "#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2", "3")
}
