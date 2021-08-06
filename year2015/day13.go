package year2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type happinessChangeMap map[string]map[string]int

func Day13Part1(input string) (int, error) {
	m, err := parseHappinessSpecs(strings.TrimSpace(input))
	if err != nil {
		return 0, err
	}
	return findMaximumHappinessChange(m), nil
}

func Day13Part2(input string) (int, error) {
	m, err := parseHappinessSpecs(strings.TrimSpace(input))
	if err != nil {
		return 0, err
	}

	injectYou(m)
	return findMaximumHappinessChange(m), nil
}

func injectYou(m happinessChangeMap) {
	names := collectAllNames(m)
	youName := "You"
	for _, spec := range m {
		spec[youName] = 0
	}
	you := make(map[string]int)
	m[youName] = you
	for _, name := range names {
		you[name] = 0
	}
}

func findMaximumHappinessChange(m happinessChangeMap) int {
	names := collectAllNames(m)
	maxHappinessChange := 0
	namePermutations := permutate(names)
	for _, perm := range namePermutations {
		happinessChange := calculateTotalHappinessChange(m, perm)
		if happinessChange > maxHappinessChange {
			maxHappinessChange = happinessChange
		}
	}

	return maxHappinessChange
}

func calculateTotalHappinessChange(m happinessChangeMap, seatings []string) int {
	total := 0
	for i, name := range seatings {
		leftI := i - 1
		if leftI == -1 {
			leftI = len(seatings) - 1
		}
		rightI := i + 1
		if rightI == len(seatings) {
			rightI = 0
		}

		left := seatings[leftI]
		right := seatings[rightI]

		spec := m[name]
		total += spec[left]
		total += spec[right]
	}
	return total
}

func collectAllNames(specs happinessChangeMap) []string {
	out := make([]string, len(specs))
	i := 0
	for name := range specs {
		out[i] = name
		i++
	}
	return out
}

var happinessSpecRegexp = regexp.MustCompile(`(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+).`)

func parseHappinessSpecs(input string) (happinessChangeMap, error) {
	allMatches := happinessSpecRegexp.FindAllStringSubmatch(input, -1)
	if len(allMatches) == 0 {
		return nil, fmt.Errorf("invalid happiness specs")
	}

	out := make(happinessChangeMap, len(allMatches))
	for _, matches := range allMatches {
		sign := 1
		if matches[2] == "lose" {
			sign = -1
		}
		happinessChangeValue, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, fmt.Errorf("invalid happiness number: %w", err)
		}
		happinessChange := happinessChangeValue * sign
		source := matches[1]
		target := matches[4]
		if sourceMap, found := out[source]; found {
			sourceMap[target] = happinessChange
		} else {
			out[source] = map[string]int{target: happinessChange}
		}
	}
	return out, nil
}

func permutate(input []string) [][]string {
	out := make([][]string, 0)
	if len(input) == 1 {
		return [][]string{input}
	}
	for i, s := range input {
		others := copySkipping(input, i)
		var othersPerm [][]string
		if len(others) == 1 {
			othersPerm = [][]string{others}
		} else {
			othersPerm = permutate(others)
		}
		for _, p := range othersPerm {
			out = append(out, prepend(p, s))
		}
	}
	return out
}

func prepend(slice []string, s string) []string {
	out := make([]string, len(slice)+1)
	out[0] = s
	copy(out[1:], slice)
	return out
}

func copySkipping(s []string, skipIndex int) []string {
	out := make([]string, len(s)-1)
	copy(out, s[:skipIndex])
	copy(out[skipIndex:], s[skipIndex+1:])
	return out
}
