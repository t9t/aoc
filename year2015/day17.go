package year2015

import (
	"fmt"
	"strconv"
	"strings"
)

func Day17Part1(input string) (int, error) {
	containers, err := parseContainers(input)
	if err != nil {
		return 0, err
	}
	return findNumberOfCombinationsOfContainersToFitEggnog(containers, 150), nil
}

func parseContainers(input string) ([]int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	out := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		out[i] = n
	}
	return out, nil
}

func findNumberOfCombinationsOfContainersToFitEggnog(containers []int, eggnog int) int {
	totalWays := 0
	for i := 0; i < (1 << len(containers)); i++ {
		step := i
		total := 0
		for _, container := range containers {
			if step%2 == 1 {
				total += container
			}
			step = step / 2
		}
		if total == eggnog {
			totalWays += 1
		}
	}
	return totalWays
}

func Day17Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
