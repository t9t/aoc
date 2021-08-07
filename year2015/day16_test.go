package year2015

import (
	"reflect"
	"testing"
)

func Test_parseAuntSueLine(t *testing.T) {
	tests := []struct {
		input   string
		wantN   int
		wantMap map[string]int
	}{
		{"Sue 1: cars: 7, akitas: 2, goldfish: 0", 1, map[string]int{"cars": 7, "akitas": 2, "goldfish": 0}},
		{"Sue 135: vizslas: 9, cats: 3, trees: 10", 135, map[string]int{"vizslas": 9, "cats": 3, "trees": 10}},
		{"Sue 431: children: 0, samoyeds: 1, pomeranians: 11", 431, map[string]int{"children": 0, "samoyeds": 1, "pomeranians": 11}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, got1, err := parseAuntSueLine(tt.input)
			if err != nil {
				t.Errorf("parseAuntSueLine() error = %v", err)
				return
			}
			if got != tt.wantN {
				t.Errorf("parseAuntSueLine() gotN = %v, want %v", got, tt.wantN)
			}
			if !reflect.DeepEqual(got1, tt.wantMap) {
				t.Errorf("parseAuntSueLine() gotMap = %v, want %v", got1, tt.wantMap)
			}
		})
	}
}

func Test_findMatchingAuntSue(t *testing.T) {
	props := map[string]int{"children": 3, "cats": 7, "akitas": 11, "trees": 8}
	input := `
Sue 7: children: 3, cats: 7, akitas: 0
Sue 13: children: 3, cats: 7, trees: 8
Sue 251: children: 5, cats: 9, akias: 11
	`
	want := 13
	got, err := findMatchingAuntSue(input, props)
	if err != nil {
		t.Errorf("findMatchingAuntSue() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("findMatchingAuntSue() = %v, want %v", got, want)
	}
}
