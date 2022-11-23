package year2018

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(17, Day17Part1, Day17Part2)
}

const (
	clay        = '#'
	sand        = '.'
	reached     = '|'
	water       = '~'
	outOfBounds = 'x'
)

func Day17Part1(input string) (string, error) {
	gridMinY, gridMaxY := math.MaxInt32, math.MinInt32

	grid := make(map[xAndY]byte)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ", ")
		leftParts := strings.Split(parts[0], "=")
		rightParts := strings.Split(parts[1], "=")
		rightNums := strings.Split(rightParts[1], "..")

		minX, minY := math.MaxInt32, math.MaxInt32
		maxX, maxY := math.MinInt32, math.MinInt32
		var l, lowerR, upperR int
		var err error
		if l, err = strconv.Atoi(leftParts[1]); err != nil {
			return "", err
		} else if lowerR, err = strconv.Atoi(rightNums[0]); err != nil {
			return "", err
		} else if upperR, err = strconv.Atoi(rightNums[1]); err != nil {
			return "", err
		}
		if leftParts[0] == "x" {
			minX, maxX = l, l
			minY, maxY = lowerR, upperR
		} else {
			minX, maxX = lowerR, upperR
			minY, maxY = l, l
		}
		if minY < gridMinY {
			gridMinY = minY
		}
		if maxY > gridMaxY {
			gridMaxY = maxY
		}
		for y := minY; y <= maxY; y += 1 {
			for x := minX; x <= maxX; x += 1 {
				grid[xAndY{x: x, y: y}] = clay
			}
		}
	}

	get := func(x, y int) byte {
		if y > gridMaxY {
			return outOfBounds
		}
		c, found := grid[xAndY{x: x, y: y}]
		if found {
			return c
		} else {
			return sand
		}
	}

	findEdge := func(startX, dx int, y int) (int, bool) {
		for otherX := startX; ; otherX += dx {
			down := get(otherX, y+1)
			if down == sand || down == reached {
				// We fall down, so edge doesn't exist
				return otherX, false
			}
			if get(otherX, y) == clay {
				return otherX, true
			}
		}
	}

	dripDown := func(x, y int) []xAndY {
		for {
			this := get(x, y)
			if this == sand {
				grid[xAndY{x: x, y: y}] = reached
			}

			// Fall down as far as we can
			below := get(x, y+1)
			if below == outOfBounds {
				// Falling down outside the grid, we can stop
				return []xAndY{}
			} else if below == reached {
				// We've been here before, we can stop
				return []xAndY{}
			} else if below != sand {
				break
			}
			// Go one down
			y += 1
		}

		// Below is clay or water, so check left and right edges

		leftEdgeX, leftEdgeFound := findEdge(x, -1, y)
		rightEdgeX, rightEdgeFound := findEdge(x, 1, y)
		inBasin := leftEdgeFound && rightEdgeFound
		fill := byte(reached)
		if inBasin {
			fill = water
		}
		for otherX := leftEdgeX + 1; otherX < rightEdgeX; otherX++ {
			grid[xAndY{x: otherX, y: y}] = fill
		}

		continuations := make([]xAndY, 0)
		if inBasin {
			continuations = append(continuations, xAndY{x: x, y: y - 1})
		}
		if !leftEdgeFound {
			continuations = append(continuations, xAndY{x: leftEdgeX, y: y})
		}
		if !rightEdgeFound {
			continuations = append(continuations, xAndY{x: rightEdgeX, y: y})
		}
		return continuations
	}

	coords := []xAndY{{x: 500, y: 0}}
	for len(coords) != 0 {
		next := make([]xAndY, 0)
		for _, xy := range coords {
			for _, contd := range dripDown(xy.x, xy.y) {
				next = append(next, contd)
			}
		}
		coords = next
	}

	reachable := 0
	for xy, c := range grid {
		if xy.y >= gridMinY && xy.y <= gridMaxY && (c == water || c == reached) {
			reachable++
		}
	}

	return strconv.Itoa(reachable), nil
}

func Day17Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 17 part 2 not implemented")
}

type xAndY struct {
	x int
	y int
}

func drawGrid(grid map[xAndY]byte) {
	fmt.Printf(buildGrid(grid))
}

func buildGrid(grid map[xAndY]byte) string {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32
	for item := range grid {
		if item.x < minX {
			minX = item.x
		}
		if item.x > maxX {
			maxX = item.x
		}
		if item.y < minY {
			minY = item.y
		}
		if item.y > maxY {
			maxY = item.y
		}
	}
	var sb strings.Builder
	for y := minY - 1; y <= maxY+1; y += 1 {
		for x := minX - 1; x <= maxX+1; x += 1 {
			v, found := grid[xAndY{x: x, y: y}]
			if !found {
				v = sand
			}
			if x == 500 && y == 0 {
				v = '+'
			}
			sb.WriteByte(v)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
