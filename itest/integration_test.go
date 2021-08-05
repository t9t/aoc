// +build itest

package itest

import (
	"aoc/registry"
	"aoc/year2015"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

type result struct {
	year, day, part int
	result          string
}

func Test_allTheThings(t *testing.T) {
	year2015.RegisterAll()

	resultData, err := os.ReadFile("../input/results.txt")
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

			inputFile := fmt.Sprintf("../input/%d-%d.txt", r.year, r.day)
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

func parseLine(line string) (result, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return result{}, fmt.Errorf("invalid line %q", line)
	}

	inputParts := strings.Split(strings.TrimSpace(parts[0]), "-")
	if len(inputParts) != 3 {
		return result{}, fmt.Errorf("invalid line %q", line)
	}

	var year, day, part int
	var err error
	if year, err = strconv.Atoi(inputParts[0]); err != nil {
		return result{}, fmt.Errorf("invalid line %q: %v", line, err)
	}
	if day, err = strconv.Atoi(inputParts[1]); err != nil {
		return result{}, fmt.Errorf("invalid line %q: %v", line, err)
	}
	if part, err = strconv.Atoi(inputParts[2]); err != nil {
		return result{}, fmt.Errorf("invalid line %q: %v", line, err)
	}

	return result{year: year, day: day, part: part, result: strings.TrimSpace(parts[1])}, nil
}
