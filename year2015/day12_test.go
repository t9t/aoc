package year2015

import (
	"reflect"
	"testing"
)

func Test_listNumbersIn(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{"no numbers", []int{}},
		{`[1,2,3]`, []int{1, 2, 3}},
		{`{"a":2,"b":4}`, []int{2, 4}},
		{`[[[3]]]`, []int{3}},
		{`{"a":{"b":4},"c":-1}`, []int{4, -1}},
		{`{"a":[-1,1]}`, []int{-1, 1}},
		{`[-1,{"a":1}]`, []int{-1, 1}},
		{`[]`, []int{}},
		{`{}`, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := listNumbersIn(tt.input)
			if err != nil {
				t.Errorf("listNumbersIn() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listNumbersIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay12Part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"no numbers", 0},
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
