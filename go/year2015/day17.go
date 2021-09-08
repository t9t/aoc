package year2015

import (
	"strconv"
	"strings"
)

func Day17Part1(input string) (int, error) {
	containers, err := parseContainers(input)
	if err != nil {
		return 0, err
	}
	n, _ := findNumberOfCombinationsOfContainersToFitEggnog(containers, 150)
	return n, nil
}

func Day17Part2(input string) (int, error) {
	containers, err := parseContainers(input)
	if err != nil {
		return 0, err
	}
	_, n := findNumberOfCombinationsOfContainersToFitEggnog(containers, 150)
	return n, nil
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

func findNumberOfCombinationsOfContainersToFitEggnog(containers []int, eggnog int) (int, int) {
	totalWays := 0
	minWays := 0
	minContainerCount := len(containers)
	allCombinationsBitMask := 1 << len(containers)
	for i := 1; i <= allCombinationsBitMask; i++ {
		total := 0
		containerCount := 0

		for c, container := range containers {
			if (i>>c)&1 != 0 {
				total += container
				containerCount++
			}
		}

		if total != eggnog {
			continue
		}

		totalWays += 1
		if containerCount < minContainerCount {
			minContainerCount = containerCount
			minWays = 1
		} else if minContainerCount == containerCount {
			minWays++
		}
	}
	return totalWays, minWays
}
