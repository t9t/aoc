package year2015

import (
	"fmt"
	"strconv"
	"strings"
)

type lightOperation int

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
	grid := make([]bool, lightsCount*lightsCount)
	for _, i := range instructions {
		applyInstruction(grid, i)
	}
	lit := 0
	for _, on := range grid {
		if on {
			lit++
		}
	}
	return lit, nil
}

func applyInstruction(grid []bool, i lightInstruction) {
	for x := i.startX; x <= i.endX; x++ {
		for y := i.startY; y <= i.endY; y++ {
			idx := lightIndex(x, y)
			cur := grid[idx]
			new := cur
			switch i.op {
			case turnLightOn:
				new = true
			case turnLightOff:
				new = false
			case toggleLight:
				new = !cur
			}
			grid[idx] = new
		}
	}
}

func lightIndex(x, y int) int {
	return y*lightsCount + x
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
