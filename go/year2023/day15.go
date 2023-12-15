package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(15, Day15Part1, Day15Part2)
}

func Day15Part1(input string) (string, error) {
	parts := strings.Split(strings.ReplaceAll(input, "\n", ""), ",")
	sum := 0
	for _, part := range parts {
		hash := 0
		for _, r := range part {
			hash += int(r)
			hash *= 17
			hash %= 256
		}
		sum += hash
	}

	return strconv.Itoa(sum), nil
}

func Day15Part2(input string) (string, error) {
	steps := strings.Split(strings.ReplaceAll(input, "\n", ""), ",")
	type lens struct {
		label       string
		focalLength int
	}
	var boxes [256][]lens

	for _, step := range steps {
		var label string
		var operation byte
		var focalLength int
		if strings.HasSuffix(step, "-") {
			operation = '-'
			label = strings.TrimSuffix(step, "-")
		} else {
			operation = '='
			parts := strings.Split(step, "=")
			if len(parts) != 2 {
				return "", fmt.Errorf("invalid step %s (expected 2 parts but got %d)", step, len(parts))
			}
			label = parts[0]
			var err error
			if focalLength, err = strconv.Atoi(parts[1]); err != nil {
				return "", fmt.Errorf("invalid step %s: %w", step, err)
			}
		}

		hash := 0
		for _, r := range label {
			hash += int(r)
			hash *= 17
			hash %= 256
		}
		box := boxes[hash]
		if operation == '-' {
			newBox := make([]lens, 0)
			for _, item := range box {
				if item.label != label {
					newBox = append(newBox, item)
				}
			}
			boxes[hash] = newBox
		} else if operation == '=' {
			newBox, found := make([]lens, 0), false
			for _, item := range box {
				if item.label == label {
					newBox = append(newBox, lens{label: label, focalLength: focalLength})
					found = true
				} else {
					newBox = append(newBox, item)
				}
			}
			if !found {
				newBox = append(newBox, lens{label: label, focalLength: focalLength})
			}
			boxes[hash] = newBox
		}
	}

	sum := 0
	for i, box := range boxes {
		for slotNum, lens := range box {
			sum += (i + 1) * (slotNum + 1) * lens.focalLength
		}
	}

	return strconv.Itoa(sum), nil
}
