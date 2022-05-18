package year2018

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(6, Day6Part1, Day6Part2)
}

func Day6Part1(input string) (string, error) {
	type xAndY struct{ x, y int }

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	setMinMax := func(min, max *int, n int) {
		if n > *max {
			*max = n
		}
		if n < *min {
			*min = n
		}
	}

	coords := make([]xAndY, 0)
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ", ")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", fmt.Errorf("invalid line %s: %w", line, err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("invalid line %s: %w", line, err)
		}

		coords = append(coords, xAndY{x: x, y: y})
		setMinMax(&minX, &maxX, x)
		setMinMax(&minY, &maxY, y)
	}

	sizes := make([]int, len(coords))
	for i := range coords {
		sizes[i] = 0
	}

	largest := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			distance, index, shared := -1, -1, false

			for i, coord := range coords {
				d := abs(x-coord.x) + abs(y-coord.y)
				if distance == -1 {
					distance = d
					index = i
				} else if d < distance {
					distance = d
					index = i
					shared = false
					continue
				} else if d == distance {
					shared = true
				}
			}
			if !shared {
				size := sizes[index] + 1
				sizes[index] = size
				if size > largest {
					largest = size
				}
			}
		}
	}

	return strconv.Itoa(largest), nil
}

func Day6Part2(input string) (string, error) {
	return day6Part2(input, 10_000)
}

func day6Part2(input string, maxTotalDistance int) (string, error) {
	type xAndY struct{ x, y int }

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	coords := make([]xAndY, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ", ")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", fmt.Errorf("invalid line %s: %w", line, err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("invalid line %s: %w", line, err)
		}
		coords = append(coords, xAndY{x: x, y: y})
	}

	maxDelta := 0
	for i, a := range coords {
		for j, b := range coords {
			if i == j {
				continue
			}
			d := abs(a.x-b.x) + abs(a.y-b.y)
			if d > maxDelta {
				maxDelta = d
			}
		}
	}

	regionSize := 0
	for x := -maxDelta; x <= maxDelta; x++ {
		for y := -maxDelta; y <= maxDelta; y++ {
			totalDistance := 0
			for _, coord := range coords {
				totalDistance += abs(x-coord.x) + abs(y-coord.y)
			}
			if totalDistance < maxTotalDistance {
				regionSize++
			}
		}
	}

	return strconv.Itoa(regionSize), nil
}
