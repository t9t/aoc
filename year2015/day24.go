package year2015

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day24Part1(input string) (int, error) {
	return findQuantumEntanglementOfSmallestEquallySummedGroup(input, 3)
}

func Day24Part2(input string) (int, error) {
	return findQuantumEntanglementOfSmallestEquallySummedGroup(input, 4)
}

func findQuantumEntanglementOfSmallestEquallySummedGroup(input string, groupSize int) (int, error) {
	packages, err := parsePackages(input)
	if err != nil {
		return 0, err
	}

	sum := sumInts(packages)
	if sum%groupSize != 0 {
		return 0, fmt.Errorf("sum %d of packages not equally divisible by %d", sum, groupSize)
	}
	groupSum := sum / groupSize

	smallest := findSmallestSubsetSummingTo(groupSum, packages)
	if smallest == nil {
		return 0, fmt.Errorf("could not group packages to group sum %d", groupSum)
	}

	return calculateQuantumEntanglement(smallest), nil
}

func parsePackages(input string) ([]int, error) {
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

func sumInts(ints []int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return sum
}

func calculateQuantumEntanglement(ints []int) int {
	if len(ints) == 0 {
		return 0
	}
	power := ints[0]
	for i := 1; i < len(ints); i++ {
		next := power * ints[i]
		if next < power {
			// overflow
			return math.MaxInt64
		}
		power = next
	}
	return power
}

func findSmallestSubsetSummingTo(target int, intsOrig []int) []int {
	ints := make([]int, len(intsOrig))
	copy(ints, intsOrig)
	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})

	for i := 1; i <= 100; i++ {
		if r := findFirstSubsetSummingTo(target, ints, 1, i); r != nil {
			return r
		}
	}

	return nil
}

func findFirstSubsetSummingTo(target int, ints []int, depth, maxDepth int) []int {
	for i, n := range ints {
		if n == target {
			return []int{n}
		} else if n > target {
			continue
		}
		if depth < maxDepth {
			nextTarget := target - n
			if r := findFirstSubsetSummingTo(nextTarget, ints[i+1:], depth+1, maxDepth); r != nil {
				return prependInt(n, r)
			}
		}
	}
	return nil
}

func prependInt(n int, other []int) []int {
	out := make([]int, len(other)+1)
	out[0] = n
	copy(out[1:], other)
	return out
}
