package year2015

import (
	"strings"
)

func Day5Part1(input string) (int, error) {
	return countNiceWords(input, isNice)
}

func Day5Part2(input string) (int, error) {
	return countNiceWords(input, isActuallyNice)
}

func countNiceWords(input string, testFunc func(input string) bool) (int, error) {
	nice := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if testFunc(line) {
			nice++
		}
	}
	return nice, nil
}

func isNice(s string) bool {
	return vowelCount(s) >= 3 && containsALetterTwiceInARow(s) && !containsUnNiceStrings(s)
}

func isActuallyNice(s string) bool {
	return hasRepeatingNonOverlappingPair(s) && hasLetterSandwichedInPair(s)
}

func vowelCount(s string) int {
	vowels := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels++
		}
	}
	return vowels
}

func containsALetterTwiceInARow(s string) bool {
	if len(s) < 1 {
		return false
	}
	prev := s[0]
	for i := 1; i < len(s); i++ {
		c := s[i]
		if c == prev {
			return true
		}
		prev = c
	}
	return false
}

func containsUnNiceStrings(s string) bool {
	return strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy")
}

func hasRepeatingNonOverlappingPair(s string) bool {
	if len(s) < 3 {
		return false
	}
	li := len(s) - 1
	for i := 0; i < li; i++ {
		test := s[i : i+2]
		for j := i + 2; j < li; j++ {
			if test == s[j:j+2] {
				return true
			}
		}
	}
	return false
}

func hasLetterSandwichedInPair(s string) bool {
	if len(s) < 3 {
		return false
	}
	li := len(s) - 1
	for i := 1; i < li; i++ {
		if s[i-1] == s[i+1] {
			return true
		}
	}
	return false
}
