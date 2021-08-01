package year2015

import (
	"reflect"
	"testing"
)

func Test_traverse(t *testing.T) {
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
			got, err := traverse(tt.instructions)
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
