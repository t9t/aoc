package year2015

import (
	"testing"
)

func Test_containsUnNiceStrings(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"ab", true},
		{"cd", true},
		{"pq", true},
		{"xy", true},
		{"abcdpqxy", true},
		{"abcdefghijklmnopqrstuvwxyz", true},
		{"efghijklmnorstuvwz", false},
		{"dafjghakjhsdhakhdajsduasid", false},
		{"a", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := containsUnNiceStrings(tt.input); got != tt.want {
				t.Errorf("containsUnNiceStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsALetterTwiceInARow(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"xx", true},
		{"abcdde", true},
		{"aabbccdd", true},
		{"abcdefghijklmnopqrstuvwxyz", false},
		{"dasfjbhaskjjsda", true},
		{"a", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := containsALetterTwiceInARow(tt.input); got != tt.want {
				t.Errorf("containsALetterTwiceInARow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vowelCount(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"a", 1},
		{"khgdfkshdksf", 0},
		{"aaa", 3},
		{"aeiou", 5},
		{"aeiouaeiou", 10},
		{"aei", 3},
		{"xazegov", 3},
		{"aeiouaeiouaeiou", 15},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := vowelCount(tt.input); got != tt.want {
				t.Errorf("vowelCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNice(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := isNice(tt.input); got != tt.want {
				t.Errorf("isNice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay5Part1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"ugknbfddgicrmopn\naaa\njchzalrnumimnmhp\nhaegwjzuvuyypxyu\ndvszwmarrgswjxmb", 2},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Day5Part1(tt.input)
			if err != nil {
				t.Errorf("Day5Part1() error = %v, wantErr %v", err, false)
				return
			}
			if got != tt.want {
				t.Errorf("Day5Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasRepeatingNonOverlappingPair(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"a", false},
		{"aa", false},
		{"aba", false},
		{"xyxy", true},
		{"axyxy", true},
		{"aabcdefgaa", true},
		{"aaa", false},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", true},
		{"ieodomkazucvgmuy", false},
		{"qiqqlmcgnhzparyg", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := hasRepeatingNonOverlappingPair(tt.input); got != tt.want {
				t.Errorf("hasRepeatingNonOverlappingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasLetterSandwichedInPair(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"a", false},
		{"ab", false},
		{"abb", false},
		{"xyx", true},
		{"abcdefeghi", true},
		{"aaa", true},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := hasLetterSandwichedInPair(tt.input); got != tt.want {
				t.Errorf("hasLetterSandwichedInPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isActuallyNice(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
		{"qiqqlmcgnhzparyg", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := isActuallyNice(tt.input); got != tt.want {
				t.Errorf("isActuallyNice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay5Part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"qjhvhtzxzqqjkmpb\nxxyxx\nuurcxstgmygtbstg\nieodomkazucvgmuy", 2},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Day5Part2(tt.input)
			if err != nil {
				t.Errorf("Day5Part2() error = %v, wantErr %v", err, false)
				return
			}
			if got != tt.want {
				t.Errorf("Day5Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
