package year2015

import "testing"

func TestDay4Part1(t *testing.T) {
	tests := []struct {
		input   string
		want    int
		wantErr bool
	}{
		{"abcdef", 609043, false},
		{"pqrstuv", 1048970, false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Day4Part1(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Day4Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day4Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
