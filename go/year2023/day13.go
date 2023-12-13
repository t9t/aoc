package year2023

import (
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(13, Day13Part1, Day13Part2)
}

func Day13Part1(input string) (string, error) {
	return day13(input, 0)
}

func Day13Part2(input string) (string, error) {
	return day13(input, 1)
}

func day13(input string, maxDiff int) (string, error) {
	chunks := strings.Split(input, "\n\n")
	summary := 0

	findReflectionLine := func(maxDim1, maxDim2 int, getc func(int, int) byte) (bool, int) {
	outer:
		for n := 0; n < maxDim1-1; n += 1 {
			differences := 0
			for dn := 0; dn < maxDim1; dn += 1 {
				mind, plusd := n-dn, n+dn+1
				if mind < 0 || plusd >= maxDim1 {
					break
				}
				for k := 0; k < maxDim2; k += 1 {
					if getc(mind, k) != getc(plusd, k) {
						differences += 1
						if differences > maxDiff {
							continue outer
						}
					}
				}
			}
			if differences == maxDiff {
				return true, n
			}
		}
		return false, 0
	}

	for _, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		numRows, numCols := len(lines), len(lines[0])

		if f, x := findReflectionLine(numCols, numRows, func(x, y int) byte {
			return lines[y][x]
		}); f {
			summary += x + 1
		}

		if f, y := findReflectionLine(numRows, numCols, func(y, x int) byte {
			return lines[y][x]
		}); f {
			summary += 100 * (y + 1)
		}
	}

	return strconv.Itoa(summary), nil
}
