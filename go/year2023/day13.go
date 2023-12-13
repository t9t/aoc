package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(13, Day13Part1, Day13Part2)
}

func Day13Part1(input string) (string, error) {
	chunks := strings.Split(input, "\n\n")

	summary := 0
	for _, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		numRows, numCols := len(lines), len(lines[0])

		row := func(y int) string {
			return lines[y]
		}

		col := func(x int) string {
			var sb strings.Builder
			for _, line := range lines {
				sb.WriteByte(line[x])
			}
			return sb.String()
		}

		xMatch := false
	verticalReflection:
		for x := 1; x < numCols; x++ {
			a, b := col(x-1), col(x)
			if a == b {
				offset := 1
				for {
					leftX, rightX := x-1-offset, x+offset
					if leftX < 0 || rightX == numCols {
						summary += x
						xMatch = true
						break verticalReflection
					}
					if left, right := col(leftX), col(rightX); left == right {
						offset += 1
						continue
					}
					break
				}
			}
		}

		if xMatch {
			// Assumption: a chunk cannot have both a horizontal and vertical reflection
			continue
		}

	horizontalReflection:
		for y := 1; y < numRows; y++ {
			a, b := row(y-1), row(y)
			if a == b {
				offset := 1
				for {
					topY, bottomY := y-1-offset, y+offset
					if topY < 0 || bottomY == numRows {
						summary += y * 100
						break horizontalReflection
					}
					if top, bottom := row(topY), row(bottomY); top == bottom {
						offset += 1
						continue
					}
					break
				}
			}
		}
	}

	return strconv.Itoa(summary), nil
}

func Day13Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 13 part 2 not implemented")
}
