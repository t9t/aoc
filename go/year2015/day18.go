package year2015

import (
	"fmt"
	"strings"
)

func Day18Part1(input string) (int, error) {
	lights, err := parseLightsGrid(input)
	if err != nil {
		return 0, err
	}

	return process100StepsAndCountLitLights(lights, false), nil
}

func Day18Part2(input string) (int, error) {
	lights, err := parseLightsGrid(input)
	if err != nil {
		return 0, err
	}

	max := len(lights) - 1
	lights[0][0] = true
	lights[max][0] = true
	lights[0][max] = true
	lights[max][max] = true

	return process100StepsAndCountLitLights(lights, true), nil
}

func process100StepsAndCountLitLights(lights [][]bool, cornerLightsAreStuck bool) int {
	for i := 0; i < 100; i++ {
		lights = processLightsAnimationStep(lights, cornerLightsAreStuck)
	}

	return countLitLights(lights)
}

func parseLightsGrid(input string) ([][]bool, error) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	lineCount := len(lines)
	out := make([][]bool, lineCount)
	for i, line := range lines {
		if len(line) != lineCount {
			return nil, fmt.Errorf("invalid line %q in lights grid: line length %d != line count %d", line, len(line), lineCount)
		}
		row := make([]bool, lineCount)
		for j, c := range line {
			var on bool
			if c == '.' {
				on = false
			} else if c == '#' {
				on = true
			} else {
				return nil, fmt.Errorf("invalid char '%c' in lights grid line %q", c, line)
			}
			row[j] = on
		}
		out[i] = row
	}

	return out, nil
}

func processLightsAnimationStep(lights [][]bool, cornerLightsAreStuck bool) [][]bool {
	l := len(lights)
	out := make([][]bool, l)
	for y, line := range lights {
		out[y] = make([]bool, l)
		copy(out[y], line)
	}

	for y, line := range lights {
		for x, on := range line {
			if cornerLightsAreStuck && isCornerLight(y, x, l) {
				// No touchy
				continue
			}

			n := countOnNeighborLights(lights, y, x)
			if on {
				if n != 2 && n != 3 {
					out[y][x] = false
				}
			} else {
				if n == 3 {
					out[y][x] = true
				}
			}
		}
	}
	return out
}

func isCornerLight(y, x, total int) bool {
	max := total - 1
	return (y == 0 && x == 0) || (y == max && x == max) || (y == max && x == 0) || (y == 0 && x == max)
}

func countLitLights(lights [][]bool) int {
	n := 0
	for _, line := range lights {
		for _, on := range line {
			if on {
				n++
			}
		}
	}
	return n
}

func countOnNeighborLights(lights [][]bool, y, x int) int {
	neighborsOn := 0
	l := len(lights)
	for dy := y - 1; dy <= y+1; dy++ {
		if dy < 0 || dy >= l {
			continue
		}
		for dx := x - 1; dx <= x+1; dx++ {
			if dx < 0 || dx >= l || (dx == x && dy == y) {
				continue
			}

			if lights[dy][dx] {
				neighborsOn++
			}
		}
	}
	return neighborsOn
}
