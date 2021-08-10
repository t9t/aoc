package year2015

import (
	"aoc/registry"
	"strconv"
)

func fromInt(f func(input string) (int, error)) registry.Execution {
	return func(input string) (string, error) {
		i, err := f(input)
		return strconv.Itoa(i), err
	}
}

func RegisterAll() {
	registry.MustRegister(2015, 1, 1, fromInt(Day1Part1))
	registry.MustRegister(2015, 1, 2, fromInt(Day1Part2))
	registry.MustRegister(2015, 2, 1, fromInt(Day2Part1))
	registry.MustRegister(2015, 2, 2, fromInt(Day2Part2))
	registry.MustRegister(2015, 3, 1, fromInt(Day3Part1))
	registry.MustRegister(2015, 3, 2, fromInt(Day3Part2))
	registry.MustRegister(2015, 4, 1, fromInt(Day4Part1))
	registry.MustRegister(2015, 4, 2, fromInt(Day4Part2))
	registry.MustRegister(2015, 5, 1, fromInt(Day5Part1))
	registry.MustRegister(2015, 5, 2, fromInt(Day5Part2))
	registry.MustRegister(2015, 6, 1, fromInt(Day6Part1))
	registry.MustRegister(2015, 6, 2, fromInt(Day6Part2))
	registry.MustRegister(2015, 7, 1, fromInt(Day7Part1))
	registry.MustRegister(2015, 7, 2, fromInt(Day7Part2))
	registry.MustRegister(2015, 8, 1, fromInt(Day8Part1))
	registry.MustRegister(2015, 8, 2, fromInt(Day8Part2))
	registry.MustRegister(2015, 9, 1, fromInt(Day9Part1))
	registry.MustRegister(2015, 9, 2, fromInt(Day9Part2))
	registry.MustRegister(2015, 10, 1, fromInt(Day10Part1))
	registry.MustRegister(2015, 10, 2, fromInt(Day10Part2))
	registry.MustRegister(2015, 11, 1, Day11Part1)
	registry.MustRegister(2015, 11, 2, Day11Part2)
	registry.MustRegister(2015, 12, 1, fromInt(Day12Part1))
	registry.MustRegister(2015, 12, 2, fromInt(Day12Part2))
	registry.MustRegister(2015, 13, 1, fromInt(Day13Part1))
	registry.MustRegister(2015, 13, 2, fromInt(Day13Part2))
	registry.MustRegister(2015, 14, 1, fromInt(Day14Part1))
	registry.MustRegister(2015, 14, 2, fromInt(Day14Part2))
	registry.MustRegister(2015, 15, 1, fromInt(Day15Part1))
	registry.MustRegister(2015, 15, 2, fromInt(Day15Part2))
	registry.MustRegister(2015, 16, 1, fromInt(Day16Part1))
	registry.MustRegister(2015, 16, 2, fromInt(Day16Part2))
	registry.MustRegister(2015, 17, 1, fromInt(Day17Part1))
	registry.MustRegister(2015, 17, 2, fromInt(Day17Part2))
	registry.MustRegister(2015, 18, 1, fromInt(Day18Part1))
	registry.MustRegister(2015, 18, 2, fromInt(Day18Part2))
	registry.MustRegister(2015, 19, 1, fromInt(Day19Part1))
	registry.MustRegister(2015, 19, 2, fromInt(Day19Part2))
	registry.MustRegister(2015, 20, 1, fromInt(Day20Part1))
	registry.MustRegister(2015, 20, 2, fromInt(Day20Part2))
}
