package year2015

import (
	"encoding/json"
	"fmt"
)

func Day12Part1(input string) (int, error) {
	return unmarshalJsonAndCountNumbersIgnoringObjectsContainingCertainValues(input, "")
}

func Day12Part2(input string) (int, error) {
	return unmarshalJsonAndCountNumbersIgnoringObjectsContainingCertainValues(input, "red")
}

func unmarshalJsonAndCountNumbersIgnoringObjectsContainingCertainValues(input, ignoredValue string) (int, error) {
	var obj interface{}
	if err := json.Unmarshal([]byte(input), &obj); err != nil {
		return 0, err
	}

	return countNumbersRecursivelyIgnoringObjectsContainingCertainValues(obj, ignoredValue), nil
}

func countNumbersRecursivelyIgnoringObjectsContainingCertainValues(obj interface{}, ignoredValue string) int {
	if i, ok := obj.(float64); ok {
		return int(i)
	}

	if array, ok := obj.([]interface{}); ok {
		total := 0
		for _, element := range array {
			total += countNumbersRecursivelyIgnoringObjectsContainingCertainValues(element, ignoredValue)
		}
		return total
	}

	if m, ok := obj.(map[string]interface{}); ok {
		total := 0
		for _, v := range m {
			if ignoredValue != "" && v == ignoredValue {
				return 0
			}
			total += countNumbersRecursivelyIgnoringObjectsContainingCertainValues(v, ignoredValue)
		}
		return total
	}

	if _, ok := obj.(string); ok {
		return 0
	}

	panic(fmt.Sprintf("impossible json object %#v (type %T)", obj, obj))
}
