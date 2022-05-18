package year2018

import (
	"aoc/registry"
)

func mustRegisterPair(day int, part1 registry.Execution, part2 registry.Execution) {
	registry.MustRegisterPair(2018, day, part1, part2)
}
