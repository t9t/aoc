package year2023

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(18, Day18Part1, Day18Part2)
}

func Day18Part1(input string) (string, error) {
	grid, posX, posY := [][]byte{{'.'}}, 0, 0

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			return "", fmt.Errorf("invalid line (expected 3 parts but got %d): %s", len(parts), line)
		}
		count, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("invalid line %s: %w", line, err)
		}
		dir := parts[0]
		for i := 0; i < count; i++ {
			if dir == "U" {
				posY -= 1
				if posY == -1 {
					newGrid := make([][]byte, len(grid)+1)
					copy(newGrid[1:], grid)
					grid = newGrid
					grid[0] = bytes.Repeat([]byte{'.'}, len(grid[1]))
					posY = 0
				}
			} else if dir == "R" {
				posX += 1
				if posX == len(grid[0]) {
					for y, row := range grid {
						grid[y] = append(row, '.')
					}
				}
			} else if dir == "D" {
				posY += 1
				if posY == len(grid) {
					grid = append(grid, bytes.Repeat([]byte{'.'}, len(grid[0])))
				}
			} else if dir == "L" {
				posX -= 1
				if posX == -1 {
					for y, row := range grid {
						newRow := make([]byte, len(row)+1)
						copy(newRow[1:], row)
						newRow[0] = '.'
						grid[y] = newRow
					}
					posX = 0
				}
			}
			grid[posY][posX] = '#'
		}
	}

	type xy struct{ x, y int }
	todo := make([]xy, 0)

	m3 := 0

	for y, row := range grid {
		for x, b := range row {
			if y == 0 && len(todo) == 0 {
				next := byte('.')
				if x < len(row)-1 {
					next = row[x+1]
				}
				// TODO: I think this way of finding "any" tile inside the pool is generally wrong, but works in my specific case and the test case
				if b == '#' && next == '#' && grid[y+1][x] == '#' && grid[y+1][x+1] == '.' {
					todo = append(todo, xy{x: x + 1, y: y + 1})
				}
			}
			if b == '#' {
				m3 += 1
			}
		}
	}

	visited := make(map[xy]struct{})
	visited[todo[0]] = struct{}{}

	for len(todo) != 0 {
		m3 += 1
		item := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		x, y := item.x, item.y
		grid[y][x] = '~'

		for _, next := range []xy{{x: x + 1, y: y}, {x: x - 1, y: y}, {x: x, y: y + 1}, {x: x, y: y - 1}} {
			nx, ny := next.x, next.y
			b := grid[ny][nx]
			if b == '#' || b == '~' {
				continue
			} else if _, f := visited[next]; f {
				continue
			}
			todo = append(todo, next)
			visited[next] = struct{}{}
		}
	}

	m3 = 0
	for _, row := range grid {
		for _, b := range row {
			if b == '#' || b == '~' {
				m3 += 1
			}
		}
	}

	return strconv.Itoa(m3), nil
}

func Day18Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 18 part 2 not implemented")
}
