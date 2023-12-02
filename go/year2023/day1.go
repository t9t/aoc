package year2023

import (
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(1, Day1Part1, Day1Part2)
}

func Day1Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		first, last := 0, 0
		for _, c := range line {
			if c >= 49 && c <= 57 {
				last = int(c - 48)
				if first == 0 {
					first = last
				}
			}
		}
		sum += first*10 + last
	}

	return strconv.Itoa(sum), nil
}

func Day1Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		first, last := 0, 0
		for i := range line {
			c := line[i]
			n := 0
			if c >= 49 && c <= 57 {
				n = int(c - 48)
			} else {
				sub := line[i:]
				switch {
				case strings.HasPrefix(sub, "one"):
					n = 1
				case strings.HasPrefix(sub, "two"):
					n = 2
				case strings.HasPrefix(sub, "three"):
					n = 3
				case strings.HasPrefix(sub, "four"):
					n = 4
				case strings.HasPrefix(sub, "five"):
					n = 5
				case strings.HasPrefix(sub, "six"):
					n = 6
				case strings.HasPrefix(sub, "seven"):
					n = 7
				case strings.HasPrefix(sub, "eight"):
					n = 8
				case strings.HasPrefix(sub, "nine"):
					n = 9
				}
			}
			if n != 0 {
				last = n
				if first == 0 {
					first = n
				}
			}
		}
		sum += first*10 + last
	}

	return strconv.Itoa(sum), nil
}
