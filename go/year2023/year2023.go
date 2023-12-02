package year2023

import (
	"aoc/registry"
)

func mustRegisterPair(day int, part1 registry.Execution, part2 registry.Execution) {
	registry.MustRegisterPair(2023, day, part1, part2)
}
