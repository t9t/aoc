package year2018

import (
	"math"
	"regexp"
	"strconv"
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
	_, totalReached, err := day17(input)
	return strconv.Itoa(totalReached), err
}

func Day17Part2(input string) (string, error) {
	waterSettled, _, err := day17(input)
	return strconv.Itoa(waterSettled), err
}

func day17(input string) (waterSettled, totalReached int, err error) {
	gridMinY, gridMaxY := math.MaxInt32, math.MinInt32

	var re = regexp.MustCompile(`(?m)(\w)=(\d+), (\w)=(\d+)\.\.(\d+)`)
	grid := make(map[xAndY]byte)

	for _, match := range re.FindAllStringSubmatch(input, -1) {
		var leftNum, rangeStart, rangeEnd int
		if leftNum, err = strconv.Atoi(match[2]); err != nil {
			return
		} else if rangeStart, err = strconv.Atoi(match[4]); err != nil {
			return
		} else if rangeEnd, err = strconv.Atoi(match[5]); err != nil {
			return
		}

		var minX, maxX, minY, maxY int
		if match[1] == "x" {
			minX, maxX = leftNum, leftNum
			minY, maxY = rangeStart, rangeEnd
		} else {
			minX, maxX = rangeStart, rangeEnd
			minY, maxY = leftNum, leftNum
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

	for xy, c := range grid {
		if xy.y >= gridMinY && xy.y <= gridMaxY {
			if c == water {
				waterSettled += 1
			}
			if c == water || c == reached {
				totalReached += 1
			}
		}
	}
	return
}

type xAndY struct {
	x int
	y int
}
