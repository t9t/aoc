package year2018

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(6, Day6Part1, Day6Part2)
}

func Day6Part1(input string) (string, error) {
	type xAndY struct{ x, y int }
	type entry struct {
		distance, index int
		shared          bool
	}

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

	closests := make(map[xAndY]entry)
	for x := -maxDelta; x <= maxDelta; x++ {
		for y := -maxDelta; y <= maxDelta; y++ {
			pos := xAndY{x: x, y: y}
			closest, hasClosest := closests[pos]

			for i, coord := range coords {
				d := abs(x-coord.x) + abs(y-coord.y)
				if !hasClosest {
					closest = entry{distance: d, index: i, shared: false}
					closests[pos] = closest
					hasClosest = true
					continue
				}

				if d > closest.distance {
					continue
				}

				if d < closest.distance {
					closest = entry{distance: d, index: i, shared: false}
					closests[pos] = closest
					continue
				}

				if d == closest.distance {
					closest.shared = true
					closests[pos] = closest
				}
			}
		}
	}

	sizes := make([]int, len(coords))
	for i := range coords {
		sizes[i] = 0
	}

	edgeCoords := make(map[int]struct{})
	for c, e := range closests {
		if !e.shared {
			sizes[e.index]++
			if c.x == -maxDelta || c.x == maxDelta || c.y == -maxDelta || c.y == maxDelta {
				edgeCoords[e.index] = struct{}{}
			}
		}
	}

	largest := 0
	for i, s := range sizes {
		if _, bordersEdge := edgeCoords[i]; !bordersEdge && s > largest {
			largest = s
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
