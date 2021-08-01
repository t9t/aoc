package year2015

import "fmt"

type coords struct {
	x, y int
}

func Day3Part1(input string) (int, error) {
	visits, err := traverse(input)
	if err != nil {
		return 0, err
	}
	return len(visits), nil
}

func traverse(instructions string) (map[coords]int, error) {
	// Starting point 0,0 always gets a visit at the start
	visits := map[coords]int{{x: 0, y: 0}: 1}
	x, y := 0, 0
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		switch instruction {
		case '>':
			x += 1
		case '<':
			x -= 1
		case '^':
			y -= 1
		case 'v':
			y += 1
		default:
			return nil, fmt.Errorf("invalid instruction %s", string(instruction))
		}
		c := coords{x: x, y: y}
		n := visits[c]
		visits[c] = n + 1
	}
	return visits, nil
}
