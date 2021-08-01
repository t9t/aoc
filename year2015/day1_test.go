package year2015

import (
	"testing"
)

func Test_day1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"floor 0, 1", "(())", 0, false},
		{"floor 0, 2", "()()", 0, false},
		{"floor 3, 1", "(((", 3, false},
		{"floor 3, 2", "(()(()(", 3, false},
		{"floor 3, 3", "))(((((", 3, false},
		{"floor -1, 1", "())", -1, false},
		{"floor -1, 2", "))(", -1, false},
		{"floor -3, 1", ")))", -3, false},
		{"floor -3, 2", ")())())", -3, false},
		{"invalid floor", "-", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Day1Part1(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("day1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("day1() = %v, want %v", got, tt.want)
			}
		})
	}
}
