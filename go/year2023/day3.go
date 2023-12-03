package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(3, Day3Part1, Day3Part2)
}

func Day3Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	type xy struct {
		x, y int
	}
	cache := make(map[xy]int)
	uniques := make(map[xy]int)

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			c := line[x]
			if c == '.' { // Empty
			} else if c >= '0' && c <= '9' { // Number
			} else { // Symbol
				for nx := x - 1; nx <= x+1; nx++ {
					for ny := y - 1; ny <= y+1; ny++ {
						_, f := cache[xy{x: nx, y: ny}]
						if !f {
							f, n, fx := expandNumber(lines, nx, ny)
							if f {
								ns := strconv.Itoa(n)
								for cx := fx; cx < fx+len(ns); cx++ {
									cache[xy{x: cx, y: ny}] = n
								}
								uniques[xy{x: fx, y: ny}] = n
							}
						}
					}
				}
			}
		}
	}
	sum := 0
	for _, v := range uniques {
		sum += v
	}
	return strconv.Itoa(sum), nil
}

func expandNumber(lines []string, x, y int) (bool, int, int) {
	if y < 0 || x < 0 || y >= len(lines) || x >= len(lines[0]) {
		return false, 0, 0
	}
	f, c := getDigitChar(lines, x, y)
	if !f {
		return false, 0, 0
	}
	leftX := x
	s := fmt.Sprintf("%c", c)
	for nx := x - 1; nx >= 0; nx-- {
		f, nxc := getDigitChar(lines, nx, y)
		if f {
			s = string(nxc) + s
			leftX = nx
		} else {
			break
		}
	}
	for nx := x + 1; nx < len(lines[y]); nx++ {
		f, nxc := getDigitChar(lines, nx, y)
		if f {
			s = s + string(nxc)
		} else {
			break
		}
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err) // We read only digits, so this must succeed
	}
	return true, n, leftX
}

func getDigitChar(lines []string, x, y int) (bool, byte) {
	c := lines[y][x]
	if c < '0' || c > '9' {
		return false, 0
	}
	return true, c
}

func Day3Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 3 part 2 not implemented")
}
