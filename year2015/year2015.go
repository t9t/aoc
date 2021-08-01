package year2015

import "aoc/registry"

func RegisterAll() {
	registry.Register(2015, 1, 1, Day1Part1)
	registry.Register(2015, 1, 2, Day1Part2)
	registry.Register(2015, 2, 1, Day2Part1)
	registry.Register(2015, 2, 2, Day2Part2)
	registry.Register(2015, 3, 1, Day3Part1)
	registry.Register(2015, 3, 2, Day3Part2)
}
