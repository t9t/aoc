package year2015

import (
	"testing"
)

func TestDay8Part1(t *testing.T) {
	input := `""
"abc"
"aaa\"aaa"
"\x27"`
	want := 12 // (2+5+10+6) - (0+3+7+1) = 23-11 = 12
	got, err := Day8Part1(input)
	if err != nil {
		t.Errorf("Day8Part1() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("Day8Part1() = %v, want %v", got, want)
	}
}

func TestDay8Part2(t *testing.T) {
	input := `""
"abc"
"aaa\"aaa"
"\x27"`
	want := 19 // (6+9+16+11) - (2+5+10+6) = 42-23 = 19
	got, err := Day8Part2(input)
	if err != nil {
		t.Errorf("Day8Part2() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("Day8Part2() = %v, want %v", got, want)
	}
}

func Test_characterLength(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{`""`, 0},
		{`"abc"`, 3},
		{`"aaa\"aaa"`, 7},
		{`"\x27"`, 1},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := characterLength(tt.input); got != tt.want {
				t.Errorf("characterLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countEncodedStringLength(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{`""`, 6},
		{`"abc"`, 9},
		{`"aaa\"aaa"`, 16},
		{`"\x27"`, 11},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := countEncodedStringLength(tt.input); got != tt.want {
				t.Errorf("countEncodedStringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
