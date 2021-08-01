package year2015

import (
	"reflect"
	"testing"
)

func Test_traverseSingleSanta(t *testing.T) {
	tests := []struct {
		instructions string
		want         map[coords]int
		wantErr      bool
	}{
		{">", map[coords]int{
			{0, 0}: 1,
			{1, 0}: 1,
		}, false},
		{"^>v<", map[coords]int{
			{0, 0}:  2,
			{0, -1}: 1,
			{1, -1}: 1,
			{1, 0}:  1,
		}, false},
		{"^v^v^v^v^v", map[coords]int{
			{0, 0}:  6,
			{0, -1}: 5,
		}, false},
		{"invalid input", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.instructions, func(t *testing.T) {
			got, err := traverseSingleSanta(tt.instructions)
			if (err != nil) != tt.wantErr {
				t.Errorf("traverse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_traverseDoubleSanta(t *testing.T) {
	tests := []struct {
		instructions string
		want         map[coords]int
		wantErr      bool
	}{
		{"^v", map[coords]int{
			{0, 0}:  2,
			{0, -1}: 1,
			{0, 1}:  1,
		}, false},
		{"^>v<", map[coords]int{
			{0, 0}:  4,
			{0, -1}: 1,
			{1, 0}:  1,
		}, false},
		{"^v^v^v^v^v", map[coords]int{
			{0, 0}:  2,
			{0, -1}: 1,
			{0, -2}: 1,
			{0, -3}: 1,
			{0, -4}: 1,
			{0, -5}: 1,
			{0, 1}:  1,
			{0, 2}:  1,
			{0, 3}:  1,
			{0, 4}:  1,
			{0, 5}:  1,
		}, false},
		{"invalid input", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.instructions, func(t *testing.T) {
			got, err := traverseDoubleSanta(tt.instructions)
			if (err != nil) != tt.wantErr {
				t.Errorf("traverse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
