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
	gridSerialNumber, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}

	var day11 day11
	// https://en.wikipedia.org/wiki/Summed-area_table
	summedAreaTable := make([][]int, 301)
	summedAreaTable[0] = make([]int, 301)
	for y := 1; y <= 300; y++ {
		summedAreaTable[y] = make([]int, 301)
		for x := 1; x <= 300; x++ {
			pl := day11.calculatePowerLevel(x, y, gridSerialNumber)
			t2 := summedAreaTable[y-1][x]
			t3 := summedAreaTable[y][x-1]
			t4 := summedAreaTable[y-1][x-1]
			summedAreaTable[y][x] = pl + t2 + t3 - t4
		}
	}

	maxPowerLevel, maxX, maxY, matchingSquareSize := 0, 0, 0, 0
	for squareSize := 1; squareSize <= 300; squareSize++ {
		for y := 1; y <= 300-squareSize; y++ {
			for x := 1; x <= 300-squareSize; x++ {
				pl := summedAreaTable[y+squareSize][x+squareSize] + summedAreaTable[y][x] - summedAreaTable[y+squareSize][x] - summedAreaTable[y][x+squareSize]

				if pl > maxPowerLevel {
					maxPowerLevel, maxX, maxY, matchingSquareSize = pl, x, y, squareSize
				}
			}
		}
	}

	// TODO: I don't understand why I have to do +1 :(
	return fmt.Sprintf("%d,%d,%d", maxX+1, maxY+1, matchingSquareSize), nil
}

type day11 struct{}

func (day11) calculatePowerLevel(x, y, gridSerialNumber int) int {
	rackId := x + 10
	return (((((rackId * y) + gridSerialNumber) * rackId) / 100) % 10) - 5
}
