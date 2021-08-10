package year2015

import (
	"strconv"
	"testing"
)

func Test_lowestHouseNumberToGetMoreThanNumberOfPresents(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{10, 1},
		{60, 4},
		{130, 8},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if houseNumber := lowestHouseNumberToGetMoreThanNumberOfPresents(tt.n, 1, 0, 10); houseNumber != tt.want {
				t.Errorf("lowestHouseNumberToGetMoreThanNumberOfPresents() = %v, want %v", houseNumber, tt.want)
			}
		})
	}
}
