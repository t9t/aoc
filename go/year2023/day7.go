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
	return day7(input, false)
}

func Day7Part2(input string) (string, error) {
	return day7(input, true)
}

func day7(input string, joker bool) (string, error) {
	cardScores := make(map[byte]int)
	cards := []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	if joker {
		cards = []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
	}
	for i, card := range cards {
		cardScores[card] = i
	}

	type entry struct {
		hand       string
		score, bid int
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
			return fiveOfAKind
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

	jokerize := func(s string, score int) int {
		jokerCount := strings.Count(s, "J")
		if !joker || jokerCount == 0 {
			return score
		}

		if score == fiveOfAKind {
			return fiveOfAKind
		} else if score == fourOfAKind {
			return fiveOfAKind
		} else if score == fullHouse {
			return fiveOfAKind
		} else if score == threeOfAKind {
			return fourOfAKind
		} else if score == twoPair {
			if jokerCount == 2 {
				return fourOfAKind
			} else if jokerCount == 1 {
				return fullHouse
			}
		} else if score == onePair {
			if jokerCount == 2 {
				return threeOfAKind
			} else if jokerCount == 1 {
				return threeOfAKind
			}
		} else if score == highCard {
			return onePair
		}
		panic(fmt.Sprintf("cannot jokerize %s", s))
	}

	lines := strings.Split(input, "\n")
	entries := make([]entry, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid line (expected 2 parts but got %d): %s", len(parts), line)
		}
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("invalid line %s: %w", line, err)
		}
		hand := parts[0]
		entries[i] = entry{hand: hand, score: jokerize(hand, scoreHand(hand)), bid: bid}
	}

	sort.Slice(entries, func(i, j int) bool {
		l, r := entries[i], entries[j]
		if l.score == r.score {
			for k := 0; k < len(l.hand); k++ {
				lc, rc := l.hand[k], r.hand[k]
				if lc == rc {
					continue
				}
				return cardScores[lc] < cardScores[rc]
			}
			panic("woops")
		}
		return l.score > r.score
	})

	sum := 0
	for i, e := range entries {
		sum += (i + 1) * e.bid
	}
	return strconv.Itoa(sum), nil
}
