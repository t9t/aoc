package year2015

import (
	"strconv"
	"testing"
)

func Test_lowestHouseNumberToGetMoreThanNumberOfPresents(t *testing.T) {
	tests := []struct {
		n            int
		max          int
		presentCount int
		want         int
	}{
		{60, 0, 10, 4},
		{130, 0, 10, 8},
		{5521, 0, 10, 210},
		{1_000_000, 0, 10, 27720},

		{5521, 2, 11, 336},
		{1_000_000, 2, 11, 60608},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			houseNumber, err := lowestHouseNumberToGetMoreThanNumberOfPresents(tt.n, tt.max, tt.presentCount)
			if err != nil {
				t.Errorf("lowestHouseNumberToGetMoreThanNumberOfPresents() error = %v", err)
				return
			}
			if houseNumber != tt.want {
				t.Errorf("lowestHouseNumberToGetMoreThanNumberOfPresents() = %v, want %v", houseNumber, tt.want)
			}
		})
	}
}
