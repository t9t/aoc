package year2015

import (
	"fmt"
	"strings"
)

func Day11Part1(input string) (int, error) {
	fmt.Printf("next: %s\n", nextValidPassword(strings.TrimSpace(input)))
	return 0, fmt.Errorf("not implemented")
}

func Day11Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func nextValidPassword(s string) string {
	for {
		s = nextPassword(s)
		if isAValidPassword(s) {
			return s
		}
	}
}

func nextPassword(s string) string {
	out := make([]byte, len(s))
	increment := true
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if increment {
			c++
		}
		if c > 'z' {
			c = 'a'
		} else {
			increment = false
		}
		out[i] = c
	}
	return string(out)
}

func isAValidPassword(s string) bool {
	return includesAnIncreasingStraightOfAtLeastThreeLetters(s) &&
		!containsConfusingLetters(s) &&
		containsAtLeastTwoDifferentNonOverlappingPairsOfLetters(s)
}

// Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
func includesAnIncreasingStraightOfAtLeastThreeLetters(s string) bool {
	if len(s) < 3 {
		return false
	}
	for i := 0; i < len(s)-2; i++ {
		c := s[i]
		if c == (s[i+1]-1) && c == (s[i+2]-2) {
			return true
		}
	}
	return false
}

// Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
func containsConfusingLetters(s string) bool {
	return strings.Contains(s, "i") || strings.Contains(s, "o") || strings.Contains(s, "l")
}

// Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
func containsAtLeastTwoDifferentNonOverlappingPairsOfLetters(s string) bool {
	if len(s) < 4 {
		return false
	}
	found := make(map[string]bool)
	for i := 0; i < len(s)-1; i++ {
		c1 := s[i]
		c2 := s[i+1]
		if c1 != c2 {
			continue
		}
		found[fmt.Sprintf("%c%c", c1, c2)] = true
		if len(found) >= 2 {
			return true
		}
	}
	return false
}
