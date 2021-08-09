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
	return 0, fmt.Errorf("not implemented")
}

type atomReplacement struct {
	from, to string
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
				sb.WriteString(replacement.to)
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
	return atomReplacement{from: strings.TrimSpace(parts[0]), to: strings.TrimSpace(parts[1])}, nil
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
