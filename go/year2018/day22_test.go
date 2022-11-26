package year2018

import (
	"testing"
)

func Test_Day22Part1(t *testing.T) {
	input := "depth: 510\ntarget: 10,10"
	basicTest(t, Day22Part1, input, "114")
}
