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

	type entry struct{ destStart, sourceStart, rangeLength int }
	type seedRange struct{ start, length int }

	seedRanges := make([]seedRange, 0)
	mappings := make([][]entry, 0)

	for _, chunk := range chunks {
		if strings.HasPrefix(chunk, "seeds:") {
			parts := strings.Split(strings.TrimPrefix(chunk, "seeds: "), " ")
			start := 0

			for i, part := range parts {
				n, err := strconv.Atoi(part)
				if err != nil {
					return "", fmt.Errorf("invalid seeds line %s: %w", strings.TrimSpace(chunk), err)
				} else if i%2 == 0 {
					start = n
				} else {
					seedRanges = append(seedRanges, seedRange{start: start, length: n})
					start = 0
				}
			}

			fmt.Printf("seed ranges: %+v\n", seedRanges)
			continue
		}
		chunkEntry := make([]entry, 0)

		fmt.Printf("Chunk: %v\n", chunk)
		lines := strings.Split(strings.TrimSpace(chunk), "\n")
		for i := 1; i < len(lines); i++ {
			line := lines[i]
			fmt.Printf("    Line: %v\n", line)
			parts := strings.Split(line, " ")
			var e entry
			var err error
			if e.destStart, err = strconv.Atoi(parts[0]); err != nil {
				return "", fmt.Errorf("invalid line %s: %w", line, err)
			} else if e.sourceStart, err = strconv.Atoi(parts[1]); err != nil {
				return "", fmt.Errorf("invalid line %s: %w", line, err)
			} else if e.rangeLength, err = strconv.Atoi(parts[2]); err != nil {
				return "", fmt.Errorf("invalid line %s: %w", line, err)
			}
			chunkEntry = append(chunkEntry, e)
		}
		mappings = append(mappings, chunkEntry)
	}

	for _, a := range mappings {
		fmt.Printf("~~\n")
		for _, b := range a {
			fmt.Printf("    %+v\n", b)
		}
	}

	// Assumption: the mappings are listed in the input in processing order
	// Assumption: mapping ranges don't overlap

	a := func(r seedRange, e entry) (matched []seedRange, notMatched []seedRange) {
		minSeed, maxSeed := r.start, r.start+r.length-1
		if minSeed > maxSeed {
			panic(fmt.Sprintf("minSeed %d > maxSeed %d", minSeed, maxSeed))
		}
		minSrc, maxSrc := e.sourceStart, e.sourceStart+e.rangeLength-1
		destDiff := e.destStart - e.sourceStart
		if maxSeed < minSrc || minSeed > maxSrc {
			// seed range is entirely "before" or "after" matching range, no overlap
			return nil, []seedRange{r}
		}
		if minSeed < minSrc && maxSeed > maxSrc {
			// seed range starts "before" matching range and ends "after", meaning the seed range fully encloses the matching range
			before := seedRange{start: minSeed, length: minSrc - minSeed}
			inside := seedRange{start: e.sourceStart, length: e.rangeLength}
			after := seedRange{start: maxSrc + 1, length: maxSeed - maxSrc}
			return []seedRange{inside}, []seedRange{before, after}
		}
		if minSeed >= minSeed && maxSeed <= maxSeed {
			// seed range is entirely "inside" the matching range, convert the full thing
			return []seedRange{{start: minSeed + destDiff, length: r.length}}, nil
		}
		if minSeed < minSrc && maxSeed <= maxSrc {
			// seed range starts "before" matching range, and ends "inside"
			before := []seedRange{{start: minSeed, length: minSrc - minSeed}}
			inside := []seedRange{{start: minSrc + destDiff, length: maxSeed - minSrc + 1}}
			return inside, before
		}
		if minSeed <= maxSrc && maxSeed > maxSrc {
			// seed range starts "inside" and ends "after" matching range
			inside := []seedRange{{start: minSeed + destDiff, length: maxSrc - minSeed + 1}}
			after := []seedRange{{start: maxSrc + 1, length: maxSeed - maxSrc}}
			return inside, after
		}
		panic("woops")
	}

	b := func(ranges []seedRange, e entry) (allMatched, allUnmatched []seedRange) {
		for _, r := range ranges {
			matched, unmatched := a(r, e)
			for _, m := range matched {
				allMatched = append(allMatched, m)
			}
			for _, m := range unmatched {
				allUnmatched = append(allUnmatched, m)
			}
		}
		return allMatched, allUnmatched
	}

	c := func(ranges []seedRange, entries []entry) []seedRange {
		var allMatched []seedRange
		allUnmatched := ranges
		for _, e := range entries {
			matched, unmatched := b(allUnmatched, e)
			for _, m := range matched {
				allMatched = append(allMatched, m)
			}
			allUnmatched = unmatched
		}
		// All unmatched are just taken as-is
		for _, m := range allUnmatched {
			allMatched = append(allMatched, m)
		}
		return allMatched
	}

	//	seedRanges = []seedRange{{start: 82, length: 1}}
	for _, mapping := range mappings {
		_s := c(seedRanges, mapping)
		fmt.Printf("Mapped %d to %d (%+v -> %+v)\n", len(seedRanges), len(_s), seedRanges, _s)
		seedRanges = _s
	}

	lowcation := math.MaxInt
	for _, r := range seedRanges {
		fmt.Printf("Out seedrange: %+v\n", r)
		if r.start < lowcation {
			lowcation = r.start
		}
	}

	fmt.Printf("%d\n", lowcation)

	return "", fmt.Errorf("Day 5 part 2 not implemented")
}
