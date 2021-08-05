package year2015

import (
	"testing"
)

func Test_includesAnIncreasingStraightOfAtLeastThreeLetters(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"a", false},
		{"ab", false},
		{"abd", false},
		{"abdc", false},
		{"abc", true},
		{"bcd", true},
		{"cde", true},
		{"xyz", true},
		{"blablablabc", true},
		{"hijklmmn", true},
		{"abbceffg", false},
		{"abbcegjk", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := includesAnIncreasingStraightOfAtLeastThreeLetters(tt.input); got != tt.want {
				t.Errorf("includesAnIncreasingStraightOfAtLeastThreeLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsConfusingLetters(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abcixyz", true},
		{"abcoxyz", true},
		{"abclxyz", true},
		{"abcxyz", false},
		{"hijklmmn", true},
		{"abbceffg", false},
		{"abbcegjk", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := containsConfusingLetters(tt.input); got != tt.want {
				t.Errorf("containsConfusingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsAtLeastTwoDifferentNonOverlappingPairsOfLetters(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"paaxyzbbq", true},
		{"aaaa", false},
		{"aabb", true},
		{"aaaaa", false},
		{"aaa", false},
		{"aaba", false},
		{"abab", false},
		{"aabbaa", true},
		{"hijklmmn", false},
		{"abbceffg", true},
		{"abbcegjk", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := containsAtLeastTwoDifferentNonOverlappingPairsOfLetters(tt.input); got != tt.want {
				t.Errorf("containsAtLeastTwoDifferentNonOverlappingPairsOfLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAValidPassword(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"aa", false},
		{"aabb", false},
		{"hijklmmn", false},
		{"abbceffg", false},
		{"abbcegjk", false},
		{"abcdffaa", true},
		{"ghjaabcc", true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := isAValidPassword(tt.input); got != tt.want {
				t.Errorf("isAValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextPassword(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"a", "b"},
		{"xx", "xy"},
		{"xy", "xz"},
		{"xz", "ya"},
		{"ya", "yb"},
		{"az", "ba"},
		{"azzzz", "baaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := nextPassword(tt.input); got != tt.want {
				t.Errorf("nextPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextValidPassword(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := nextValidPassword(tt.input); got != tt.want {
				t.Errorf("nextValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
