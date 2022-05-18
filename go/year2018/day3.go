package year2018

import (
	"fmt"
	"regexp"
	"strconv"
)

func init() {
	mustRegisterPair(3, Day3Part1, Day3Part2)
}

func Day3Part1(input string) (string, error) {
	return day3(input, true)
}

func Day3Part2(input string) (string, error) {
	return day3(input, false)
}

func day3(input string, part1 bool) (string, error) {
	type xAndY struct{ x, y int }

	parseMatch := func(match []string) (x int, y int, w int, h int, err error) {
		if x, err = strconv.Atoi(match[2]); err != nil {
			return
		}
		if y, err = strconv.Atoi(match[3]); err != nil {
			return
		}
		if w, err = strconv.Atoi(match[4]); err != nil {
			return
		}
		h, err = strconv.Atoi(match[5])
		return
	}

	re := regexp.MustCompile(`(?m)#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	claims := make(map[xAndY]int)
	multiClaims := 0
	for _, match := range matches {
		cx, cy, cw, ch, err := parseMatch(match)
		if err != nil {
			return "", fmt.Errorf("invalid line %s: %w", match[0], err)
		}

		for x := cx; x < cx+cw; x++ {
			for y := cy; y < cy+ch; y++ {
				xy := xAndY{x: x, y: y}
				if c, found := claims[xy]; found {
					if c == 1 {
						multiClaims++
					}
					claims[xy] = c + 1
				} else {
					claims[xy] = 1
				}
			}
		}
	}

	if part1 {
		return strconv.Itoa(multiClaims), nil
	}

	for _, match := range matches {
		cx, cy, cw, ch, err := parseMatch(match)
		if err != nil {
			return "", fmt.Errorf("invalid line %s: %w", match[0], err)
		}

		if func() bool {
			for x := cx; x < cx+cw; x++ {
				for y := cy; y < cy+ch; y++ {
					if claims[xAndY{x: x, y: y}] != 1 {
						return false
					}
				}
			}
			return true
		}() {
			return match[1], nil
		}
	}
	return "", fmt.Errorf("no good")
}
