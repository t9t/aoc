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
	type xAndY struct{ x, y int }
	gridSerialNumber, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}

	powerLevelMap := make(map[xAndY]int)

	var day11 day11
	maxPowerLevel, maxX, maxY, theSquareSize := 0, 0, 0, 0
	// TODO: valid square sizes are from 1x1 to 300x300
	for squareSize := 10; squareSize <= 16; squareSize++ {
		for x := 1; x <= 301-squareSize; x++ {
			for y := 1; y <= 301-squareSize; y++ {
				pl := 0
				for dx := 0; dx <= squareSize-1; dx++ {
					for dy := 0; dy <= squareSize-1; dy++ {
						xy := xAndY{x: x + dx, y: y + dy}
						if spl, found := powerLevelMap[xy]; found {
							pl += spl
						} else {
							spl := day11.calculatePowerLevel(x+dx, y+dy, gridSerialNumber)
							powerLevelMap[xy] = spl
							pl += spl
						}
					}
				}
				if pl > maxPowerLevel {
					maxPowerLevel, maxX, maxY, theSquareSize = pl, x, y, squareSize
				}
			}
		}
	}
	return fmt.Sprintf("%d,%d,%d", maxX, maxY, theSquareSize), nil
}

type day11 struct{}

func (day11) calculatePowerLevel(x, y, gridSerialNumber int) int {
	rackId := x + 10
	return (((((rackId * y) + gridSerialNumber) * rackId) / 100) % 10) - 5
}
