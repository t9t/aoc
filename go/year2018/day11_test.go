package year2018

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day11Part1(t *testing.T) {
	basicMultiTest(t, Day11Part1, []testInput{
		{"18", "33,45"},
		{"42", "21,61"},
	})
}

func Test_Day11Part2(t *testing.T) {
	basicMultiTest(t, Day11Part2, []testInput{
		{"18", "90,269,16"},
		{"42", "232,251,12"},
	})
}

func Test_Day11CalculatePowerLevel(t *testing.T) {
	tests := []struct{ x, y, sn, want int }{
		{3, 5, 8, 4},
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("x: %d; y: %d; sn: %d", tt.x, tt.y, tt.sn), func(t *testing.T) {
			got := day11{}.calculatePowerLevel(tt.x, tt.y, tt.sn)
			assert.Equal(t, tt.want, got)
		})
	}
}
