package year2018

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(18, Day18Part1, Day18Part2)
}

func Day18Part1(input string) (string, error) {
	const (
		open       = '.'
		trees      = '|'
		lumberyard = '#'
	)

	type xAndY struct{ x, y int }
	type grid [][]byte

	lines := strings.Split(input, "\n")
	landscape := make(grid, len(lines))
	for y, line := range lines {
		row := make([]byte, len(line))
		for x, c := range line {
			row[x] = byte(c)
		}
		landscape[y] = row
	}

	getAround := func(g grid, x, y int) (opens int, treess int, lumberyards int) {
		for ny := y - 1; ny <= y+1; ny++ {
			if ny < 0 || ny >= len(g) {
				continue
			}
			for nx := x - 1; nx <= x+1; nx++ {
				row := g[ny]
				if nx < 0 || nx >= len(row) || (ny == y && nx == x) {
					continue
				}
				c := row[nx]
				if c == open {
					opens += 1
				} else if c == trees {
					treess += 1
				} else if c == lumberyard {
					lumberyards += 1
				}
			}
		}
		return
	}

	transform := func(g grid) grid {
		out := make(grid, len(g))
		for y, row := range g {
			out[y] = make([]byte, len(row))
			for x, c := range row {
				_, treesAround, lumberyardsAround := getAround(g, x, y)
				out[y][x] = c
				if c == open {
					if treesAround >= 3 {
						out[y][x] = trees
					}
				} else if c == trees {
					if lumberyardsAround >= 3 {
						out[y][x] = lumberyard
					}
				} else if c == lumberyard {
					if lumberyardsAround >= 1 && treesAround >= 1 {
						// Stays lumberyard
					} else {
						out[y][x] = open
					}
				}
			}
		}
		return out
	}

	for minute := 1; minute <= 10; minute++ {
		landscape = transform(landscape)
	}

	woodedAcres, lumberyards := 0, 0
	for _, row := range landscape {
		for _, c := range row {
			if c == trees {
				woodedAcres += 1
			} else if c == lumberyard {
				lumberyards += 1
			}
		}
	}

	return strconv.Itoa(woodedAcres * lumberyards), nil
}

func Day18Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 18 part 2 not implemented")
}
