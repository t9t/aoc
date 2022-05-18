//go:build itest_2015

package itest

import (
	"aoc/year2015"
	"testing"
)

func Test_2015(t *testing.T) {
	year2015.RegisterAll()

	runTest(t, 2015)
}
