package year2015

import (
	"fmt"
	"strconv"
	"strings"
)

func Day20Part1(input string) (int, error) {
	numberOfPresents, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}

	// TODO: remove start param and improve solution a lot
	start := 775000
	return lowestHouseNumberToGetMoreThanNumberOfPresents(numberOfPresents, start, 0, 10), nil
}

func Day20Part2(input string) (int, error) {
	numberOfPresents, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}

	// TODO: remove start param and improve solution a lot
	start := 785000
	return lowestHouseNumberToGetMoreThanNumberOfPresents(numberOfPresents, start, 50, 11), nil
}

func lowestHouseNumberToGetMoreThanNumberOfPresents(n, start, maxElfVisits, elfPresents int) int {
	houseNumber := start
	maxHouseNumber := 10_000_000
	for {
		presents := 0
		for elf := houseNumber; elf >= 1; elf-- {
			elfVisit := houseNumber / elf
			if maxElfVisits > 0 && elfVisit > maxElfVisits {
				break
			}

			if houseNumber%elf == 0 {
				presents += elfPresents * elf
				if presents >= n {
					return houseNumber
				}
			}
		}

		if houseNumber >= maxHouseNumber {
			panic(fmt.Sprintf("houseNumber %d >= maxHouseNumber %d", houseNumber, maxHouseNumber))
		}

		houseNumber++
	}
}
