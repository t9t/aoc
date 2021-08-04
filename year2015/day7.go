package year2015

import (
	"fmt"
	"strconv"
	"strings"
)

func Day7Part1(input string) (int, error) {
	wireValues, err := resolveSignals(input)
	if err != nil {
		return 0, err
	}
	a, found := wireValues["a"]
	if !found {
		return 0, fmt.Errorf("wire value \"a\" not found")
	}
	return int(a), nil
}

func resolveSignals(input string) (map[string]uint16, error) {
	wireMap, err := toWireMap(input)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve signals: %w", err)
	}

	wireValues := make(map[string]uint16)
	for wire, def := range wireMap {
		v, err := resolveWire(def, wireMap, wireValues)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve wire %q: %w", wire, err)
		}
		wireValues[wire] = v
	}

	return wireValues, nil
}

func resolveWire(def string, wireMap map[string]string, wireValues map[string]uint16) (uint16, error) {
	parts := strings.Split(def, " ")
	if len(parts) == 1 {
		// simple value
		return resolveOrParse(parts[0], wireMap, wireValues)
	} else if len(parts) == 2 {
		operator := strings.TrimSpace(parts[0])
		if !strings.HasPrefix(operator, "NOT") {
			return 0, fmt.Errorf("invalid operator %q in %q", operator, def)
		}
		i, err := resolveOrParse(parts[1], wireMap, wireValues)
		if err != nil {
			return 0, err
		}
		return ^i, nil
	} else if len(parts) == 3 {
		i1, err := resolveOrParse(parts[0], wireMap, wireValues)
		if err != nil {
			return 0, err
		}
		i2, err := resolveOrParse(parts[2], wireMap, wireValues)
		if err != nil {
			return 0, err
		}

		switch operator := strings.TrimSpace(parts[1]); operator {
		case "AND":
			return i1 & i2, nil
		case "OR":
			return i1 | i2, nil
		case "LSHIFT":
			return i1 << i2, nil
		case "RSHIFT":
			return i1 >> i2, nil
		default:
			return 0, fmt.Errorf("invalid operator %q in %q", operator, def)
		}
	}
	return 0, fmt.Errorf("invalid wire def %q", def)
}

func resolveOrParse(s string, wireMap map[string]string, wireValues map[string]uint16) (uint16, error) {
	i, err := parseUint16(s)
	if err == nil {
		return i, nil
	}
	refValue, found := wireValues[s]
	if found {
		return refValue, nil
	}

	refDef, found := wireMap[s]
	if !found {
		return 0, fmt.Errorf("invalid reference %q", s)
	}
	v, err := resolveWire(refDef, wireMap, wireValues)
	if err != nil {
		return 0, err
	}
	wireValues[s] = v
	return v, nil
}

func parseUint16(s string) (uint16, error) {
	i, err := strconv.ParseUint(strings.TrimSpace(s), 10, 16)
	return uint16(i), err
}

func toWireMap(input string) (map[string]string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	out := make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line %q: got %d parts instead of 2", line, len(parts))
		}
		wire := strings.TrimSpace(parts[1])
		if _, found := out[wire]; found {
			return nil, fmt.Errorf("duplicate definition for wire %q", wire)
		}
		out[wire] = strings.TrimSpace(parts[0])
	}
	return out, nil
}
