package year2023

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(7, Day7Part1, Day7Part2)
}

func Day7Part1(input string) (string, error) {
	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
	cardScores := make(map[byte]int)
	for i, card := range []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'} {
		cardScores[card] = i
	}

	scoreHand := func(s string) int {
		s = strings.Split(s, " ")[0]
		counts := make(map[rune]int)
		maxCount := 0
		for _, c := range s {
			nn := counts[c] + 1
			counts[c] = nn
			if nn > maxCount {
				maxCount = nn
			}
		}
		if maxCount == 5 {
			return 0 // Five of a kind
		} else if maxCount == 4 {
			return 1 // Four of a kind
		} else if len(counts) == 2 {
			return 2 // Full house (must be XXXYY or XXYYY - cannot be XYYYY as that's four of a kind)
		} else if maxCount == 3 {
			return 3 // Three of a kind
		} else if len(counts) == 3 {
			return 4 // Two pair
		} else if maxCount == 2 {
			return 5 // One pair
		} else if len(counts) == 5 {
			return 6 // High card
		}
		panic(fmt.Sprintf("invalid hand? %s -> %+v", s, counts))
	}

	lines := strings.Split(input, "\n")
	sort.Slice(lines, func(i, j int) bool {
		l, r := lines[i], lines[j]
		ls, rs := scoreHand(l), scoreHand(r)
		if ls == rs {
			for k := 0; k < len(l); k++ {
				lc, rc := l[k], r[k]
				if lc == rc {
					continue
				}
				lcs, rcs := cardScores[lc], cardScores[rc]
				return lcs < rcs
			}
			panic("woops")
		}
		return ls > rs
	})

	sum := 0
	for i, line := range lines {
		bid, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			return "", fmt.Errorf("invalid line %s: %w", line, err)
		}
		sum += (i + 1) * bid
	}
	return strconv.Itoa(sum), nil
}

func Day7Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 7 part 2 not implemented")
}
