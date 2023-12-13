package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(12, Day12Part1, Day12Part2)
}

func Day12Part1(input string) (string, error) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid line (expected 2 parts but got %d): %s", len(parts), line)
		}

		nrs := make([]int, 0)
		for _, nr := range strings.Split(parts[1], ",") {
			n, err := strconv.Atoi(nr)
			if err != nil {
				return "", fmt.Errorf("invalid line %s: %w", line, err)
			}
			nrs = append(nrs, n)
		}

		// TODO: running day1 without cache is 100x faster than with cache
		sum += day12line(parts[0], nrs)
	}

	return strconv.Itoa(sum), nil
}

func Day12Part2(input string) (string, error) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid line (expected 2 parts but got %d): %s", len(parts), line)
		}

		nrs := make([]int, 0)
		for _, nr := range strings.Split(strings.Repeat(","+parts[1], 5)[1:], ",") {
			n, err := strconv.Atoi(nr)
			if err != nil {
				return "", fmt.Errorf("invalid line %s: %w", line, err)
			}
			nrs = append(nrs, n)
		}

		sum += day12line(strings.Repeat("?"+parts[0], 5)[1:], nrs)
	}

	return strconv.Itoa(sum), nil
}

type day12CacheKey struct {
	s   string
	nrs [30]int // hopefully there are no inputs with more than 30 numbers
}

var day12cache = make(map[day12CacheKey]int)

func day12line(s string, nrs []int) int {
	k := day12CacheKey{s: s}
	for i, n := range nrs {
		k.nrs[i] = n
	}
	v, f := day12cache[k]
	if !f {
		v = day12line_(s, nrs)
		day12cache[k] = v
	}
	return v
}

func day12line_(s string, nrs []int) int {
	if len(s) == 0 && len(nrs) == 0 {
		return 1
	} else if len(nrs) == 0 {
		for _, c := range s {
			if c == '#' {
				return 0
			}
		}
		return 1
	} else if len(s) == 0 {
		return 0
	} else if s[0] == '.' {
		return day12line(s[1:], nrs)
	}

	nr := nrs[0]
	if nr > len(s) {
		return 0
	}

	chunk := s[:nr]
	if strings.Contains(chunk, ".") {
		if chunk[0] == '#' {
			return 0
		} else {
			return day12line(s[1:], nrs)
		}
	}

	after := s[nr:]
	afterC := byte(0)
	if len(after) > 0 {
		afterC = after[0]
	}

	newNrs := make([]int, len(nrs)-1)
	copy(newNrs, nrs[1:])

	next := ""
	if len(after) > 0 {
		next = after[1:]
	}

	if chunk[0] == '#' {
		if afterC == '#' {
			return 0
		}

		return day12line(next, newNrs)
	}

	if afterC == '#' {
		return day12line(s[1:], nrs)
	}

	a := day12line(s[1:], nrs)
	b := day12line(next, newNrs)
	return a + b
}
