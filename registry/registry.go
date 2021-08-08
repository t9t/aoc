package registry

import (
	"fmt"
	"sort"
)

type Execution func(input string) (string, error)

type Selector struct {
	Year, Day, Part int
}

var Map map[Selector]Execution = make(map[Selector]Execution)

func MustRegister(year, day, part int, function Execution) {
	err := Register(year, day, part, function)
	if err != nil {
		panic(err)
	}
}

func Register(year, day, part int, function Execution) error {
	selector := Selector{Year: year, Day: day, Part: part}
	if _, ok := Map[selector]; ok {
		return fmt.Errorf("%d/%d/%d already registered", year, day, part)
	}
	Map[selector] = function
	return nil
}

func Get(year, day, part int) (Execution, bool) {
	selector := Selector{Year: year, Day: day, Part: part}
	function, ok := Map[selector]
	return function, ok
}

func AllSelectorsSorted() []Selector {
	all := make([]Selector, len(Map))
	i := 0
	for selector := range Map {
		all[i] = selector
		i++
	}

	sort.Slice(all, func(i, j int) bool {
		a, b := all[i], all[j]
		if a.Year == b.Year {
			if a.Day == b.Day {
				return a.Part < b.Part
			}
			return a.Day < b.Day
		}
		return a.Year < b.Year
	})
	return all
}
