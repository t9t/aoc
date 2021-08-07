package year2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day16Part1(input string) (int, error) {
	senderProperties := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	return findMatchingAuntSue(input, senderProperties)
}

func findMatchingAuntSue(input string, senderProperties map[string]int) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		auntNumber, aunt, err := parseAuntSueLine(line)
		if err != nil {
			return 0, err
		}

		match := true
		for property, auntValue := range aunt {
			senderValue := senderProperties[property]
			if senderValue != auntValue {
				match = false
				break
			}
		}
		if match {
			return auntNumber, nil
		}
	}

	return 0, fmt.Errorf("no matching Aunt Sue found")
}

var auntSueLineRegexp = regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)

func parseAuntSueLine(s string) (int, map[string]int, error) {
	if strings.Count(s, ",") != 2 {
		return 0, nil, fmt.Errorf("unparseable aunt line %q", s)
	}

	matches := auntSueLineRegexp.FindStringSubmatch(s)
	if len(matches) == 0 {
		return 0, nil, fmt.Errorf("invalid aunt line %q", s)
	}

	if auntNumber, err := strconv.Atoi(matches[1]); err != nil {
		return 0, nil, fmt.Errorf("invalid aunt number in %q: %w", s, err)
	} else if v1, err := strconv.Atoi(matches[3]); err != nil {
		return 0, nil, fmt.Errorf("invalid number value in %q: %w", s, err)
	} else if v2, err := strconv.Atoi(matches[5]); err != nil {
		return 0, nil, fmt.Errorf("invalid number value in %q: %w", s, err)
	} else if v3, err := strconv.Atoi(matches[7]); err != nil {
		return 0, nil, fmt.Errorf("invalid number value in %q: %w", s, err)
	} else {
		return auntNumber, map[string]int{
			matches[2]: v1,
			matches[4]: v2,
			matches[6]: v3,
		}, nil
	}
}

func Day16Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
