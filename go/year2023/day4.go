package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(4, Day4Part1, Day4Part2)
}

func Day4Part1(input string) (string, error) {
	totalPoints := 0
	for _, line := range strings.Split(input, "\n") {
		_, lists, found := strings.Cut(line, ": ")
		if !found {
			return "", fmt.Errorf("invalid line (no ': '): %s", line)
		}

		winning, owned, found := strings.Cut(lists, " | ")
		if !found {
			return "", fmt.Errorf("invalid line (no ' | '): %s", line)
		}

		toNumbers := func(s string) (map[int]struct{}, error) {
			numbers := make(map[int]struct{})
			for _, item := range strings.Split(strings.TrimSpace(s), " ") {
				if item == "" {
					continue
				}
				n, err := strconv.Atoi(strings.TrimSpace(item))
				if err != nil {
					return nil, err
				}
				numbers[n] = struct{}{}
			}
			return numbers, nil
		}

		winningNumbers, err := toNumbers(winning)
		if err != nil {
			return "", err
		}

		ownedNumbers, err := toNumbers(owned)
		if err != nil {
			return "", err
		}

		score := 0
		for n := range ownedNumbers {
			if _, found := winningNumbers[n]; found {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		totalPoints += score
	}
	return strconv.Itoa(totalPoints), nil
}

func Day4Part2(input string) (string, error) {
	type card struct {
		matchingNumbers int
		count           int
	}
	cards := make([]*card, 0)
	cardIndex := 0
	for _, line := range strings.Split(input, "\n") {
		cardNum, lists, found := strings.Cut(line, ": ")
		if !found {
			return "", fmt.Errorf("invalid line (no ': '): %s", line)
		}

		cardNumber, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(cardNum, "Card ")))
		if err != nil {
			return "", err
		}

		if cardNumber != cardIndex+1 {
			return "", fmt.Errorf("assumption violated, cardNumber %d does not match cardIndex %d+1", cardNumber, cardIndex)
		}

		winning, owned, found := strings.Cut(lists, " | ")
		if !found {
			return "", fmt.Errorf("invalid line (no ' | '): %s", line)
		}

		toNumbers := func(s string) (map[int]struct{}, error) {
			numbers := make(map[int]struct{})
			for _, item := range strings.Split(strings.TrimSpace(s), " ") {
				if item == "" {
					continue
				}
				n, err := strconv.Atoi(strings.TrimSpace(item))
				if err != nil {
					return nil, err
				}
				numbers[n] = struct{}{}
			}
			return numbers, nil
		}

		winningNumbers, err := toNumbers(winning)
		if err != nil {
			return "", err
		}

		ownedNumbers, err := toNumbers(owned)
		if err != nil {
			return "", err
		}

		matching := 0
		for n := range ownedNumbers {
			if _, found := winningNumbers[n]; found {
				matching += 1
			}
		}
		cards = append(cards, &card{
			matchingNumbers: matching,
			count:           1,
		})
		cardIndex += 1
	}

	l := len(cards)
	for i, card := range cards {
		if card.matchingNumbers == 0 {
			continue
		}

		for d := i + 1; d <= i+card.matchingNumbers; d++ {
			if d == l {
				break
			}
			cards[d].count += card.count
		}
	}

	totalCardCount := 0
	for _, card := range cards {
		totalCardCount += card.count
	}
	return strconv.Itoa(totalCardCount), nil
}
