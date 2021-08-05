package year2015

import "testing"

func Test_lookAndSay(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := lookAndSay(tt.input); got != tt.want {
				t.Errorf("lookAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
