package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(19, Day19Part1, Day19Part2)
}

func Day19Part1(input string) (string, error) {
	workflowsChunk, partsChunk, f := strings.Cut(input, "\n\n")
	if !f {
		return "", fmt.Errorf("invalid input")
	}

	type step struct {
		condition func(map[byte]int) bool
		target    string
	}

	workflows := make(map[string][]step)
	for _, line := range strings.Split(workflowsChunk, "\n") {
		name, rest, f := strings.Cut(line, "{")
		if !f {
			return "", fmt.Errorf("invalid workflow line: %s\n", line)
		}
		workflows[name] = make([]step, 0)
		for _, stepStr := range strings.Split(rest[:len(rest)-1], ",") {
			first, second, f := strings.Cut(stepStr, ":")
			var s step
			if !f {
				s = step{condition: func(map[byte]int) bool { return true }, target: first}
			} else {
				category := first[0]
				op := first[1]
				if op != '>' && op != '<' {
					return "", fmt.Errorf("invalid workflow line (2nd char not < or >): %s", line)
				}
				n, err := strconv.Atoi(first[2:])
				if err != nil {
					return "", fmt.Errorf("invalid workflow line %s: %w", line, err)
				}
				s = step{condition: func(part map[byte]int) bool {
					rating := part[category]
					if op == '>' {
						return rating > n
					} else {
						return rating < n
					}
				}, target: second}
			}
			workflows[name] = append(workflows[name], s)
		}
	}

	acceptedSum := 0
	for _, line := range strings.Split(partsChunk, "\n") {
		part := make(map[byte]int)
		for _, item := range strings.Split(line[1:len(line)-1], ",") {
			category, ratingStr, f := strings.Cut(item, "=")
			if !f {
				return "", fmt.Errorf("invalid part line: %s", line)
			}
			rating, err := strconv.Atoi(ratingStr)
			if err != nil {
				return "", fmt.Errorf("invalid part line %s: %w", line, err)
			}
			if len(category) != 1 {
				return "", fmt.Errorf("invalid part line (expected 1 char but got %d in %s): %s", len(category), category, line)
			}
			part[category[0]] = rating
		}

		workflow := workflows["in"]
	workflowLoop:
		for {
			for _, step := range workflow {
				if step.condition(part) {
					if step.target == "A" { // Accepted
						for _, v := range part {
							acceptedSum += v
						}
						break workflowLoop
					} else if step.target == "R" { // Rejected
						break workflowLoop
					} else {
						workflow = workflows[step.target]
						continue workflowLoop
					}
				}
			}
			panic(fmt.Sprintf("part did not complete workflow: %+v", part))
		}
	}

	return strconv.Itoa(acceptedSum), nil
}

func Day19Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 19 part 2 not implemented")
}
