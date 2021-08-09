package year2015

import (
	"fmt"
	"strings"
)

func Day19Part1(input string) (int, error) {
	molecule, replacements, err := parseMoleculeAndAtomReplacements(input)
	if err != nil {
		return 0, err
	}

	return countPossibleAtomReplacements(molecule, replacements), nil
}

func Day19Part2(input string) (int, error) {
	molecule, _, err := parseMoleculeAndAtomReplacements(input)
	if err != nil {
		return 0, err
	}

	rn, ar, y := 0, 0, 0
	for _, atom := range molecule {
		if atom == "Rn" {
			rn++
		} else if atom == "Ar" {
			ar++
		} else if atom == "Y" {
			y++
		}
	}
	return len(molecule) - rn - ar - 2*y - 1, nil
}

type atomReplacement struct {
	from string
	to   []string
}

var b4 = make(map[string]bool)

func bla(molecule []string, replacements []atomReplacement) {
	//c := 0
	if len(molecule) == 1 {
		fmt.Printf("Len 1: %v\n", molecule)
		return
	}

	for i := 0; i < len(molecule); i++ {
		for _, replacement := range replacements {
			v := molecule[i:]
			if stringSliceStartsWith(v, replacement.to) {
				//c++
				new := combineStringSlices(molecule[:i], []string{replacement.from}, molecule[i+len(replacement.to):])
				//fmt.Printf("Molecule: %v\n    Slice: %v\n    Matches: %#v\n    New: %v\n", molecule, v, replacement, new)
				//return
				s := strings.Join(new, "")
				if !b4[s] {
					b4[s] = true
					bla(new, replacements)
				}
			}
		}
	}
	//fmt.Printf("c: %d\n", c)
}

func combineStringSlices(slices ...[]string) []string {
	out := make([]string, 0)
	for _, slice := range slices {
		out = append(out, slice...)
	}
	return out
}

func stringSliceStartsWith(toCheck []string, startsWith []string) bool {
	if len(toCheck) < len(startsWith) {
		return false
	}
	for i, v := range startsWith {
		if toCheck[i] != v {
			return false
		}
	}
	return true
}

func countPossibleAtomReplacements(molecule []string, replacements []atomReplacement) int {
	results := make(map[string]bool)
	for i, atom := range molecule {
		for _, replacement := range replacements {
			if atom == replacement.from {
				var sb strings.Builder
				for _, orig := range molecule[:i] {
					sb.WriteString(orig)
				}
				for _, r := range replacement.to {
					sb.WriteString(r)
				}
				for _, orig := range molecule[i+1:] {
					sb.WriteString(orig)
				}
				results[sb.String()] = true
			}
		}
	}
	return len(results)
}

func parseMoleculeAndAtomReplacements(input string) ([]string, []atomReplacement, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	molecule := lines[len(lines)-1]
	replacements := make([]atomReplacement, 0)
	for i := 0; i < len(lines)-2; i++ {
		replacement, err := parseAtomReplacement(lines[i])
		if err != nil {
			return nil, nil, err
		}
		replacements = append(replacements, replacement)
	}

	return untangleMolecule(molecule), replacements, nil
}

func parseAtomReplacement(s string) (atomReplacement, error) {
	parts := strings.Split(s, " => ")
	if len(parts) != 2 {
		return atomReplacement{}, fmt.Errorf("invalid atom replacement %q", s)
	}
	to := strings.TrimSpace(parts[1])
	return atomReplacement{from: strings.TrimSpace(parts[0]), to: untangleMolecule(to)}, nil
}

func untangleMolecule(molecule string) []string {
	out := make([]string, 0)
	for len(molecule) > 0 {
		atom := readNextAtom(molecule)
		out = append(out, atom)
		molecule = molecule[len(atom):]
	}
	return out
}

func readNextAtom(s string) string {
	if len(s) == 1 {
		return s
	}
	for i := 1; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			return s[:i]
		}
	}
	return s
}
