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
	return "", fmt.Errorf("Day 5 part 2 not implemented")
}
