package year2018

import (
	"fmt"
	"strconv"
)

func init() {
	mustRegisterPair(11, Day11Part1, Day11Part2)
}

func Day11Part1(input string) (string, error) {
	gridSerialNumber, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}

	var day11 day11
	maxPowerLevel, maxX, maxY := 0, 0, 0
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {
			pl := 0
			for dx := 0; dx <= 2; dx++ {
				for dy := 0; dy <= 2; dy++ {
					pl += day11.calculatePowerLevel(x+dx, y+dy, gridSerialNumber)
				}
			}
			if pl > maxPowerLevel {
				maxPowerLevel, maxX, maxY = pl, x, y
			}
		}
	}
	return fmt.Sprintf("%d,%d", maxX, maxY), nil
}

func Day11Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 11 part 2 not implemented")
}

type day11 struct{}

func (day11) calculatePowerLevel(x, y, gridSerialNumber int) int {
	rackId := x + 10
	return (((((rackId * y) + gridSerialNumber) * rackId) / 100) % 10) - 5
}
