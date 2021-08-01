package year2015

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day2Part1(input string) (int, error) {
	presents := strings.Split(input, "\n")
	total := 0
	for _, present := range presents {
		l, w, h, err := parseDimensions(present)
		if err != nil {
			return 0, fmt.Errorf("invalid input: %w", err)
		}
		total += wrappingPaperNeeded(l, w, h)
	}
	return total, nil
}

func Day2Part2(input string) (int, error) {
	presents := strings.Split(input, "\n")
	total := 0
	for _, present := range presents {
		l, w, h, err := parseDimensions(present)
		if err != nil {
			return 0, fmt.Errorf("invalid input: %w", err)
		}
		total += ribbonNeeded(l, w, h)
	}
	return total, nil
}

func parseDimensions(s string) (l int, w int, h int, err error) {
	parts := strings.Split(s, "x")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("expected 3 elements but got %d in %q", len(parts), s)
	}
	if l, err = strconv.Atoi(parts[0]); err != nil {
		return 0, 0, 0, fmt.Errorf("cannot parse %q as dimensions: %w", s, err)
	}
	if w, err = strconv.Atoi(parts[1]); err != nil {
		return 0, 0, 0, fmt.Errorf("cannot parse %q as dimensions: %w", s, err)
	}
	if h, err = strconv.Atoi(parts[2]); err != nil {
		return 0, 0, 0, fmt.Errorf("cannot parse %q as dimensions: %w", s, err)
	}
	return l, w, h, nil
}

func parseDimension(s string) (int, error) {
	if i, err := strconv.Atoi(s); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

func wrappingPaperNeeded(l, w, h int) int {
	area := surfaceArea(l, w, h)
	slack := l * w
	if w*h < slack {
		slack = w * h
	}
	if l*h < slack {
		slack = l * h
	}
	return area + slack
}

func surfaceArea(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

func ribbonNeeded(l, w, h int) int {
	dims := []int{l, w, h}
	sort.Slice(dims, func(i, j int) bool {
		return dims[i] < dims[j]
	})
	s1, s2 := dims[0], dims[1]
	shortestPathAround := 2*s1 + 2*s2
	bow := l * w * h
	return shortestPathAround + bow
}
