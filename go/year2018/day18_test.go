package year2018

import (
	"testing"
)

func Test_Day18Part1(t *testing.T) {
	input := `.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.`
	basicTest(t, Day18Part1, input, "1147")
}
