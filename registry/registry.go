package registry

import "fmt"

type Execution func(input string) (string, error)

type selector struct {
	year, day, part int
}

var registry map[selector]Execution = make(map[selector]Execution)

func MustRegister(year, day, part int, function Execution) {
	err := Register(year, day, part, function)
	if err != nil {
		panic(err)
	}
}

func Register(year, day, part int, function Execution) error {
	selector := selector{year: year, day: day, part: part}
	if _, ok := registry[selector]; ok {
		return fmt.Errorf("%d/%d/%d already registered", year, day, part)
	}
	registry[selector] = function
	return nil
}

func Get(year, day, part int) (Execution, bool) {
	selector := selector{year: year, day: day, part: part}
	function, ok := registry[selector]
	return function, ok
}
