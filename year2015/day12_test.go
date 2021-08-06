package year2015

import (
	"testing"
)

func TestDay12Part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{`{"no": "numbers"}`, 0},
		{`[1,2,3]`, 6},
		{`{"a":2,"b":4}`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`[]`, 0},
		{`{}`, 0},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Day12Part1(tt.input)
			if err != nil {
				t.Errorf("Day12Part1() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("Day12Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay12Part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{`[1,2,3]`, 6},
		{`[1,{"c":"red","b":2},3]`, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{`[1,"red",5]`, 6},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Day12Part2(tt.input)
			if err != nil {
				t.Errorf("Day12Part2() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("Day12Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
