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
	return "", fmt.Errorf("Day 8 part 2 not implemented")
}
