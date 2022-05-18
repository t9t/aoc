package itest

import (
	"aoc/registry"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type result struct {
	year, day, part int
	result          string
}

func runTest(t *testing.T, year int) {
	resultData, err := os.ReadFile(fmt.Sprintf("../../input/%d/results.txt", year))
	if err != nil {
		t.Fatalf("could not read result data: %v", err)
	}

	results, err := parseResults(strings.TrimSpace(string(resultData)))
	if err != nil {
		t.Fatalf("could not parse result data: %v", err)
	}

	for _, r := range results {
		name := fmt.Sprintf("%d-%d-%d", r.year, r.day, r.part)
		t.Run(name, func(t *testing.T) {
			execution, found := registry.Get(r.year, r.day, r.part)
			if !found {
				t.Fatal("could not find execution")
			}

			inputFile := fmt.Sprintf("../../input/%d/%d.txt", r.year, r.day)
			inputData, err := os.ReadFile(inputFile)
			if err != nil {
				t.Fatalf("could not read input file %q: %v", inputFile, err)
			}
			input := strings.TrimSpace(string(inputData))

			result, err := execution(input)
			if err != nil {
				t.Fatalf("error executing: %v", err)
			} else if result != r.result {
				t.Fatalf("wrong result: %v != %v", r.result, result)
			}
		})
	}
}

func parseResults(s string) ([]result, error) {
	out := make([]result, 0)
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for _, line := range lines {
		r, err := parseLine(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, nil
}

var resultRegex = regexp.MustCompile(`(\d{4})-(\d{1,2})-(\d{1,2}): (.+)`)

func parseLine(line string) (r result, err error) {
	parts := resultRegex.FindStringSubmatch(line)
	if len(parts) != 5 {
		return r, fmt.Errorf("invalid line %q", line)
	}

	var year, day, part int
	if year, err = strconv.Atoi(parts[1]); err != nil {
		return r, fmt.Errorf("invalid line %q: %v", line, err)
	} else if day, err = strconv.Atoi(parts[2]); err != nil {
		return r, fmt.Errorf("invalid line %q: %v", line, err)
	} else if part, err = strconv.Atoi(parts[3]); err != nil {
		return r, fmt.Errorf("invalid line %q: %v", line, err)
	}

	return result{year: year, day: day, part: part, result: strings.TrimSpace(parts[4])}, nil
}
