package year2015

import (
	"fmt"
	"strconv"
	"strings"
)

type lightOperation int
type lightTransformationFunc func(cur int, op lightOperation) int

const (
	turnLightOn = iota
	turnLightOff
	toggleLight
)

const (
	lightsCount = 1000
)

type lightInstruction struct {
	op                         lightOperation
	startX, startY, endX, endY int
}

func Day6Part1(input string) (int, error) {
	instructions, err := parseLightInstructions(strings.Split(strings.TrimSpace(input), "\n"))
	if err != nil {
		return 0, err
	}
	grid := newLightGrid(lightsCount, applyTransformation)
	for _, instruction := range instructions {
		grid.applyInstruction(instruction)
	}
	return grid.countLit(), nil
}

func Day6Part2(input string) (int, error) {
	instructions, err := parseLightInstructions(strings.Split(strings.TrimSpace(input), "\n"))
	if err != nil {
		return 0, err
	}
	grid := newLightGrid(lightsCount, applyAncientNordicElvishTransformation)
	for _, instruction := range instructions {
		grid.applyInstruction(instruction)
	}
	return grid.totalBrightness(), nil
}

func newLightGrid(sideLength int, transformationFunc lightTransformationFunc) *lightGrid {
	grid := make([]int, sideLength*sideLength)
	return &lightGrid{grid: grid, sideLength: sideLength, transformationFunc: transformationFunc}
}

type lightGrid struct {
	grid               []int
	sideLength         int
	transformationFunc lightTransformationFunc
}

func (g *lightGrid) countLit() int {
	lit := 0
	for _, on := range g.grid {
		if on == 1 {
			lit++
		}
	}
	return lit
}

func (g *lightGrid) totalBrightness() int {
	brightness := 0
	for _, v := range g.grid {
		brightness += v
	}
	return brightness
}

func (g *lightGrid) applyInstruction(i lightInstruction) {
	for x := i.startX; x <= i.endX; x++ {
		for y := i.startY; y <= i.endY; y++ {
			idx := g.idx(x, y)
			g.grid[idx] = g.transformationFunc(g.grid[idx], i.op)
		}
	}
}

func applyTransformation(cur int, op lightOperation) int {
	switch op {
	case turnLightOn:
		return 1
	case turnLightOff:
		return 0
	case toggleLight:
		if cur == 0 {
			return 1
		} else {
			return 0
		}
	}
	panic(fmt.Sprintf("invalid operation %d", op))
}

func applyAncientNordicElvishTransformation(cur int, op lightOperation) int {
	switch op {
	case turnLightOn:
		return cur + 1
	case turnLightOff:
		new := cur - 1
		if new < 0 {
			return 0
		}
		return new
	case toggleLight:
		return cur + 2
	}
	panic(fmt.Sprintf("invalid operation %d", op))
}

func (g *lightGrid) idx(x, y int) int {
	return y*g.sideLength + x
}

func parseLightInstructions(lines []string) ([]lightInstruction, error) {
	out := make([]lightInstruction, len(lines))
	for i, line := range lines {
		instruction, err := parseLightInstruction(line)
		if err != nil {
			return nil, err
		}
		out[i] = instruction
	}
	return out, nil
}

func parseLightInstruction(s string) (r lightInstruction, err error) {
	through := strings.Split(strings.TrimSpace(s), " through ")
	if len(through) != 2 {
		return r, fmt.Errorf("invalid instruction %q: expected 2 parts but got %d", s, len(through))
	}
	endX, endY, err := parseCoords(through[1])
	if err != nil {
		return r, fmt.Errorf("invalid instruction %q: %w", s, err)
	}

	first := strings.Split(strings.TrimSpace(through[0]), " ")
	if len(first) != 2 && len(first) != 3 {
		return r, fmt.Errorf("invalid instruction %q: expected 2 or 3 parts but got %d", s, len(first))
	}

	startX, startY, err := parseCoords(first[len(first)-1])
	if err != nil {
		return r, fmt.Errorf("invalid instruction %q: %w", s, err)
	}

	var op lightOperation
	if len(first) == 2 && first[0] == "toggle" {
		op = toggleLight
	} else if len(first) == 3 && first[0] == "turn" {
		if first[1] == "on" {
			op = turnLightOn
		} else if first[1] == "off" {
			op = turnLightOff
		} else {
			return r, fmt.Errorf("invalid instruction %q: turn but not on/off: %q", s, first[1])
		}
	} else {
		return r, fmt.Errorf("invalid instruction %q: no turn/toggle but: %q", s, first[0])
	}

	return lightInstruction{op: op, startX: startX, startY: startY, endX: endX, endY: endY}, nil
}

func parseCoords(s string) (x int, y int, err error) {
	parts := strings.Split(strings.TrimSpace(s), ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid coords %q", s)
	}
	if x, err = strconv.Atoi(parts[0]); err != nil {
		return 0, 0, fmt.Errorf("invalid coords %q: %w", s, err)
	}
	if y, err = strconv.Atoi(parts[1]); err != nil {
		return 0, 0, fmt.Errorf("invalid coords %q: %w", s, err)
	}
	if x >= lightsCount {
		return 0, 0, fmt.Errorf("invalid coords %q: x %d out of range", s, x)
	}
	if y >= lightsCount {
		return 0, 0, fmt.Errorf("invalid coords %q: y %d out of range", s, y)
	}
	return x, y, nil
}
