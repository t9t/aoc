package year2015

import "strings"

func Day5Part1(input string) (int, error) {
	nice := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if isNice(line) {
			nice++
		}
	}
	return nice, nil
}

func isNice(s string) bool {
	return vowelCount(s) >= 3 && containsALetterTwiceInARow(s) && !containsUnNiceStrings(s)
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
	// It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
	return strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy")
}
