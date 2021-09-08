package year2015

import (
	"fmt"
	"strconv"
	"strings"
)

func Day20Part1(input string) (int, error) {
	return day20(input, 0, 10)
}

func Day20Part2(input string) (int, error) {
	return day20(input, 50, 11)
}

func day20(input string, maxElfVisits, elfPresentsPerHouse int) (int, error) {
	numberOfPresents, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}

	return lowestHouseNumberToGetMoreThanNumberOfPresents(numberOfPresents, maxElfVisits, elfPresentsPerHouse)
}

func lowestHouseNumberToGetMoreThanNumberOfPresents(n, maxElfVisits, elfPresentsPerHouse int) (int, error) {
	size := n / 10
	presentsPerHouseNumber := make([]int, size)
	for i := 1; i < size; i++ {
		elfVisits := 0
		for j := i; j < size; j += i {
			presentsPerHouseNumber[j] += i * elfPresentsPerHouse
			elfVisits++
			if maxElfVisits != 0 && elfVisits >= maxElfVisits {
				break
			}
		}
	}
	for houseNumber, presents := range presentsPerHouseNumber {
		if presents >= n {
			return houseNumber, nil
		}
	}
	return 0, fmt.Errorf("no house number found for >=%d presents", n)
}
