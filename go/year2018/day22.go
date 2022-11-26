package year2018

import (
	"fmt"
	"regexp"
	"strconv"
)

func init() {
	mustRegisterPair(22, Day22Part1, Day22Part2)
}

func Day22Part1(input string) (string, error) {
	nums := make([]int, 3)
	for i, s := range regexp.MustCompile(`(?m)depth: (\d+)\ntarget: (\d+),(\d+)`).FindStringSubmatch(input)[1:] {
		if n, err := strconv.Atoi(s); err != nil {
			return "", err
		} else {
			nums[i] = n
		}
	}
	depth, targetX, targetY := nums[0], nums[1], nums[2]

	d := &day22{targetX: targetX, targetY: targetX, depth: depth}
	d.cache = make(map[struct{ x, y int }]int)

	riskLevel := 0
	for x := 0; x <= targetX; x++ {
		for y := 0; y <= targetY; y++ {
			riskLevel += d.erosionLevel(x, y) % 3
		}
	}

	return strconv.Itoa(riskLevel), nil
}

func Day22Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 22 part 2 not implemented")
}

type day22 struct {
	targetX, targetY, depth int
	cache                   map[struct{ x, y int }]int
}

func (d *day22) geologicalIndex(x, y int) int {
	if x == 0 && y == 0 {
		// The region at 0,0 (the mouth of the cave) has a geologic index of 0.
		return 0
	} else if x == d.targetX && y == d.targetY {
		// The region at the coordinates of the target has a geologic index of 0.
		return 0
	} else if y == 0 {
		// If the region's Y coordinate is 0, the geologic index is its X coordinate times 16807.
		return x * 16807
	} else if x == 0 {
		// If the region's X coordinate is 0, the geologic index is its Y coordinate times 48271.
		return y * 48271
	}

	// Otherwise, the region's geologic index is the result of multiplying the erosion levels of the regions at X-1,Y and X,Y-1.
	return d.erosionLevel(x-1, y) * d.erosionLevel(x, y-1)
}

func (d *day22) erosionLevel(x, y int) int {
	xy := struct{ x, y int }{x: x, y: y}
	if el, found := d.cache[xy]; found {
		return el
	}

	el := (d.geologicalIndex(x, y) + d.depth) % 20183
	d.cache[xy] = el
	return el
}
