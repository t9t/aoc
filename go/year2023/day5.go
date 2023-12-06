package year2023

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(5, Day5Part1, Day5Part2)
}

func Day5Part1(input string) (string, error) {
	chunks := strings.Split(input, "\n\n")
	converters := make(map[string]almanacTypeConverter)
	seeds := make([]int, 0)

	for _, c := range chunks {
		if strings.HasPrefix(c, "seeds:") {
			for _, item := range strings.Split(strings.TrimSpace(strings.TrimPrefix(c, "seeds: ")), " ") {
				n, err := strconv.Atoi(item)
				if err != nil {
					return "", fmt.Errorf("invalid seeds line %s: %w", c, err)
				}
				seeds = append(seeds, n)
			}
			continue
		}

		converter, err := parseAlmanacChunk(c)
		if err != nil {
			return "", err
		}
		converters[converter.inType] = converter
	}

	lowcation := math.MaxInt

	for _, seed := range seeds {
		c := converters["seed"]
		for {
			_, seed = c.converter(seed)
			if c.outType == "location" {
				if seed < lowcation {
					lowcation = seed
				}
				break
			}
			c = converters[c.outType]
		}
	}

	return strconv.Itoa(lowcation), nil
}

func parseAlmanacChunk(chunk string) (c almanacTypeConverter, err error) {
	lines := strings.Split(strings.TrimSpace(chunk), "\n")
	mapTypeLine := strings.TrimSuffix(strings.TrimSpace(lines[0]), " map:")
	mapTypeParts := strings.Split(mapTypeLine, "-")
	if len(mapTypeParts) != 3 {
		return c, fmt.Errorf("unparseable map type line: %s", chunk)
	}
	inType, outType := mapTypeParts[0], mapTypeParts[2]
	converters := make([]almanacConverterFunc, 0)

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			return c, fmt.Errorf("invalid line: %s", line)
		}
		var destinationStart, sourceStart, rangeLength int
		var err error
		if destinationStart, err = strconv.Atoi(parts[0]); err != nil {
			return c, fmt.Errorf("unparseable line %s: %w", line, err)
		} else if sourceStart, err = strconv.Atoi(parts[1]); err != nil {
			return c, fmt.Errorf("unparseable line %s: %w", line, err)
		} else if rangeLength, err = strconv.Atoi(parts[2]); err != nil {
			return c, fmt.Errorf("unparseable line %s: %w", line, err)
		}

		converters = append(converters, almanacConverter(destinationStart, sourceStart, rangeLength))
	}

	convert := func(in int) (bool, int) {
		for _, converter := range converters {
			match, out := converter(in)
			if match {
				return true, out
			}
		}
		return true, in
	}

	return almanacTypeConverter{
		inType:    inType,
		outType:   outType,
		converter: convert,
	}, nil
}

type almanacTypeConverter struct {
	inType, outType string
	converter       almanacConverterFunc
}

type almanacConverterFunc func(in int) (bool, int)

func almanacConverter(destinationStart, sourceStart, rangeLength int) almanacConverterFunc {
	maxSource := sourceStart + rangeLength
	return func(in int) (bool, int) {
		if in >= sourceStart && in <= maxSource {
			return true, destinationStart + in - sourceStart
		}
		return false, 0
	}
}

func Day5Part2(input string) (string, error) {
	chunks := strings.Split(input, "\n\n")

	type tup2 struct{ a, b int }
	type tup3 struct{ a, b, c int }

	type seedRange struct{ start, end int }
	type mapRange struct{ sourceStart, sourceEnd, diff int }
	var seedRanges []seedRange

	var maps [][]mapRange
	var names []string

	var fs [][]tup3
	var pairs []tup2

	for i, chunk := range chunks {
		if i == 0 {
			if !strings.HasPrefix(chunk, "seeds:") {
				return "", fmt.Errorf("invalid first chunk (doesn't start with 'seeds:'): %s", chunk)
			}
			start := 0
			for j, s := range strings.Split(strings.TrimSpace(strings.TrimPrefix(chunk, "seeds: ")), " ") {
				n, err := strconv.Atoi(s)
				if err != nil {
					return "", fmt.Errorf("invalid seeds chunk %s: %w", chunk, err)
				}
				if j%2 == 0 {
					start = n
				} else {
					seedRanges = append(seedRanges, seedRange{start: start, end: start + n})
					pairs = append(pairs, tup2{a: start, b: n})
				}
			}
			fmt.Printf("Seed ranges: %+v\n", seedRanges)
			continue
		}

		var m []mapRange
		var f []tup3
		lines := strings.Split(chunk, "\n")
		names = append(names, strings.TrimSpace(lines[0]))
		for j := 1; j < len(lines); j++ {
			line := lines[j]
			parts := strings.Split(line, " ")
			if len(parts) != 3 {
				return "", fmt.Errorf("invalid line (expected 3 parts but got %d): %s", len(parts), line)
			}
			var destStart, sourceStart, length int
			var err error
			if destStart, err = strconv.Atoi(parts[0]); err != nil {
				return "", fmt.Errorf("invalid line %s: %w", line, err)
			} else if sourceStart, err = strconv.Atoi(parts[1]); err != nil {
				return "", fmt.Errorf("invalid line %s: %w", line, err)
			} else if length, err = strconv.Atoi(parts[2]); err != nil {
				return "", fmt.Errorf("invalid line %s: %w", line, err)
			}
			m = append(m, mapRange{sourceStart: sourceStart, sourceEnd: sourceStart + length, diff: destStart - sourceStart})
			f = append(f, tup3{destStart, sourceStart, length})
		}
		maps = append(maps, m)
		fs = append(fs, f)
	}

	for _, m := range maps {
		fmt.Printf("Map:\n")
		for _, r := range m {
			fmt.Printf("  Range: %+v\n", r)
		}
	}

	for _, f := range fs {
		fmt.Printf("f:\n")
		for _, t := range f {
			fmt.Printf("\t%+v\n", t)
		}
	}

	a := func(ranges []seedRange, m []mapRange) []seedRange {
		var matched []seedRange
		for _, transform := range m {
			var unmatched []seedRange
			for _, r := range ranges {
				before := seedRange{start: r.start, end: min(r.end, transform.sourceStart)}
				inside := seedRange{start: max(r.start, transform.sourceStart), end: min(r.end, transform.sourceEnd)}
				after := seedRange{start: max(r.start, transform.sourceEnd), end: r.end}

				if before.end > before.start {
					unmatched = append(unmatched, before)
				}
				if inside.end > inside.start {
					matched = append(matched, seedRange{start: inside.start + transform.diff, end: inside.end + transform.diff})
				}
				if after.end > after.start {
					unmatched = append(unmatched, after)
				}
			}

			ranges = unmatched
		}
		for _, r := range matched {
			ranges = append(ranges, r)
		}
		return ranges
	}

	// Assumptions: maps are listed in processing order in the input, and map ranges don't overlap
	lowcation := math.MaxInt
	for _, inputRange := range seedRanges {
		ranges := []seedRange{inputRange}
		for _, m := range maps {
			ranges = a(ranges, m) // APPLY
		}
		for _, r := range ranges {
			if r.start < lowcation {
				lowcation = r.start
			}
		}
	}

	fmt.Printf("lowcation: %d\n", lowcation)

	apply_range := func(tuples []tup3, r []tup2) []tup2 {
		var a []tup2
		for _, _t := range tuples {
			dest, src, sz := _t.a, _t.b, _t.c
			src_end := src + sz
			var nr []tup2
			for _, _r := range r {
				st, ed := _r.a, _r.b
				before := tup2{st, min(ed, src)}
				inter := tup2{max(st, src), min(src_end, ed)}
				after := tup2{max(src_end, st), ed}
				if before.b > before.a {
					nr = append(nr, before)
				}
				if inter.b > inter.a {
					a = append(a, tup2{inter.a - src + dest, inter.b - src + dest})
				}
				if after.b > after.a {
					nr = append(nr, after)
				}
			}
			r = nr
		}
		fmt.Printf("\t\ta: %+v; r: %+v\n", a, r)
		var n []tup2
		for _, t := range a {
			n = append(n, t)
		}
		for _, t := range r {
			n = append(n, t)
		}
		return n
	}

	var p2 []int
	for _, _p := range pairs {
		st, sz := _p.a, _p.b
		r := []tup2{{st, st + sz}}
		for _, f := range fs {
			_f := f
			fmt.Printf("in: %+v\n", r)
			r = apply_range(_f, r)
			fmt.Printf("\tout: %+v\n", r)
		}
		_min := math.MaxInt
		for _, t := range r {
			fmt.Printf("%+v\n", t)
			if t.a < _min {
				_min = t.a
			}
		}
		p2 = append(p2, _min)
	}

	_min := math.MaxInt
	for _, n := range p2 {
		if n < _min {
			_min = n
		}
	}
	fmt.Printf("P2: %d\n", _min)

	return "", fmt.Errorf("Day 5 part 2 not implemented")
}
