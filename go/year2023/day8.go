package year2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(8, Day8Part1, Day8Part2)
}

func Day8Part1(input string) (string, error) {
	chunks := strings.Split(input, "\n\n")
	if len(chunks) != 2 {
		return "", fmt.Errorf("invalid input")
	}

	type node struct{ left, right string }

	instructions := chunks[0]
	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	nodes := make(map[string]node, 0)
	for _, m := range re.FindAllStringSubmatch(chunks[1], -1) {
		nodes[m[1]] = node{left: m[2], right: m[3]}
	}

	cur := nodes["AAA"]
	index, steps := 0, 0
	for {
		steps += 1
		r := instructions[index]
		index += 1
		if index == len(instructions) {
			index = 0
		}
		next := cur.left
		if r == 'R' {
			next = cur.right
		}
		if next == "ZZZ" {
			break
		}
		cur = nodes[next]
	}

	return strconv.Itoa(steps), nil
}

func Day8Part2(input string) (string, error) {
	chunks := strings.Split(input, "\n\n")
	if len(chunks) != 2 {
		return "", fmt.Errorf("invalid input")
	}

	type node struct{ left, right string }

	instructions := chunks[0]
	re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	nodes := make(map[string]node, 0)
	for _, m := range re.FindAllStringSubmatch(chunks[1], -1) {
		nodes[m[1]] = node{left: m[2], right: m[3]}
	}

	curs := make([]node, 0)
	for el, node := range nodes {
		if strings.HasSuffix(el, "A") {
			curs = append(curs, node)
		}
	}

	minSteps := make(map[int]int)
	index, steps := 0, 0
	for {
		steps += 1
		r := instructions[index]
		index += 1
		if index == len(instructions) {
			index = 0
		}

		nexts := make([]node, len(curs))
		for i, cur := range curs {
			next := cur.left
			if r == 'R' {
				next = cur.right
			}
			nexts[i] = nodes[next]
			if strings.HasSuffix(next, "Z") {
				if _, f := minSteps[i]; !f {
					minSteps[i] = steps
				}
			}
		}

		if len(minSteps) == len(nexts) {
			break
		}

		curs = nexts
	}

	p := minSteps[0]
	n := p
	for i := 1; i < len(minSteps); i++ {
		k := minSteps[i]
		for {
			if n%k == 0 {
				p = n
				break
			} else {
				n += p
			}
		}
	}

	return strconv.Itoa(n), nil
}
