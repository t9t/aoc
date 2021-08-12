package year2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day25Part1(input string) (int, error) {
	row, column, err := parseManualCodeInstruction(input)
	if err != nil {
		return 0, err
	}

	return findManualCode(row, column), nil
}

var manualCodeInstructionRegexp = regexp.MustCompile(`To continue, please consult the code grid in the manual\.  Enter the code at row (\d+), column (\d+)\.`)

func parseManualCodeInstruction(input string) (row int, column int, err error) {
	matches := manualCodeInstructionRegexp.FindStringSubmatch(strings.TrimSpace(input))
	if len(matches) != 3 {
		return 0, 0, fmt.Errorf("invalid manual code instruction")
	}
	if row, err = strconv.Atoi(matches[1]); err != nil {
		return 0, 0, fmt.Errorf("invalid row number %q in manual code instruction: %w", matches[1], err)
	}
	if column, err = strconv.Atoi(matches[2]); err != nil {
		return 0, 0, fmt.Errorf("invalid column number %q in manual code instruction: %w", matches[2], err)
	}
	return row, column, nil
}

func findManualCode(findRow, findColumn int) int {
	currentRow, currentCol := 1, 1
	totalNumbersInDiagonal := 1
	numbersLeftInDiagonal := totalNumbersInDiagonal
	diagonalStartingRow := 1

	n := 20151125
	for {
		if currentRow == findRow && currentCol == findColumn {
			return n
		}

		n = generateNextManualCode(n)

		numbersLeftInDiagonal--
		if numbersLeftInDiagonal == 0 {
			totalNumbersInDiagonal++
			numbersLeftInDiagonal = totalNumbersInDiagonal
			diagonalStartingRow++
			currentRow = diagonalStartingRow
			currentCol = 1
		} else {
			currentRow--
			currentCol++
		}
	}
}

func generateNextManualCode(previous int) int {
	return (previous * 252533) % 33554393
}
