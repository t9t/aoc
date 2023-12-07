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
	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
	cardScores := make(map[byte]int)
	// J now in front
	for i, card := range []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'} {
		cardScores[card] = i
	}

	const (
		fiveOfAKind  = 0
		fourOfAKind  = 1
		fullHouse    = 2
		threeOfAKind = 3
		twoPair      = 4
		onePair      = 5
		highCard     = 6
	)

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
			return fiveOfAKind // Five of a kind
		} else if maxCount == 4 {
			return fourOfAKind
		} else if len(counts) == 2 {
			return fullHouse
		} else if maxCount == 3 {
			return threeOfAKind
		} else if len(counts) == 3 {
			return twoPair
		} else if maxCount == 2 {
			return onePair
		} else if len(counts) == 5 {
			return highCard
		}
		panic(fmt.Sprintf("invalid hand? %s -> %+v", s, counts))
	}

	scoreHandJokerized := func(s string) int {
		jokerCount := strings.Count(s, "J")
		if jokerCount == 0 {
			return scoreHand(s)
		}
		maxScore := scoreHand(s)
		for card := range cardScores {
			if card == 'J' {
				continue
			}
			rep := strings.ReplaceAll(s, "J", string([]byte{card}))
			score := scoreHand(rep)
			if score < maxScore {
				maxScore = score
			}
		}
		return maxScore
	}

	jokerize := func(s string, score int) int {
		jokerCount := strings.Count(s, "J")
		if jokerCount == 0 {
			return score
		}

		if score == fiveOfAKind {
			return fiveOfAKind
		}
		if score == fourOfAKind {
			// XJJJJ or XXXXJ
			return fiveOfAKind
		}
		if score == fullHouse {
			// XXJJJ or JJXXX
			return fiveOfAKind
		}
		if score == threeOfAKind {
			// XYJJJ or XYYYJ or XXXYJ
			return fourOfAKind
		}
		if score == twoPair {
			if jokerCount == 2 {
				// 2 others + 2 jokers = 4 of a kind
				return fourOfAKind
			} else {
				if jokerCount != 1 {
					panic(fmt.Sprintf("%s - %d - %d", s, score, jokerCount))
				}
				// 2 others + 1 joker = 3 of a kind
				return threeOfAKind
			}
		}
		if score == onePair {
			if jokerCount == 2 {
				// 1 of 3 different ones + 2 jokers = 3 of a kind
				return threeOfAKind
			} else {
				if jokerCount != 1 {
					panic(fmt.Sprintf("%s - %d - %d", s, score, jokerCount))
				}
				// 2 of a pair + 1 joker = 3 of a kind
				return threeOfAKind
			}
		}
		if score == highCard {
			// ABCDJ -> ABCJJ
			return twoPair
		}
		panic(fmt.Sprintf("cannot jokerize %s", s))
	}

	func(any) {}(jokerize)

	lines := strings.Split(input, "\n")
	sort.Slice(lines, func(i, j int) bool {
		l, r := lines[i], lines[j]
		ls, rs := scoreHandJokerized(l), scoreHandJokerized(r)
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
