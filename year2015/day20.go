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
	start := 770000
	return lowestHouseNumberToGetMoreThanNumberOfPresents(numberOfPresents, start), nil
}

func lowestHouseNumberToGetMoreThanNumberOfPresents(n, start int) int {
	houseNumber := start
	maxHouseNumber := 10_000_000
	for {
		presents := 0
		for elf := houseNumber; elf >= 1; elf-- {
			if houseNumber%elf == 0 {
				presents += 10 * elf
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

func Day20Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
