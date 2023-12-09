package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(9, Day9Part1, Day9Part2)
}

func Day9Part1(input string) (string, error) {
	return day9(input, false)
}

func Day9Part2(input string) (string, error) {
	return day9(input, true)
}

func day9(input string, left bool) (string, error) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		var numbers []int
		for _, part := range strings.Split(line, " ") {
			n, err := strconv.Atoi(part)
			if err != nil {
				return "", fmt.Errorf("invalid line %s: %w", line, err)
			}
			numbers = append(numbers, n)
		}

		getNum := func() int {
			if left {
				return numbers[0]
			}
			return numbers[len(numbers)-1]
		}
		lastNumbers := []int{getNum()}
		for {
			var next []int
			allZeroes := true
			for i := 0; i < len(numbers)-1; i++ {
				n, k := numbers[i], numbers[i+1]
				j := k - n
				next = append(next, j)
				if j != 0 {
					allZeroes = false
				}
			}

			if allZeroes {
				break
			}
			numbers = next
			lastNumbers = append(lastNumbers, getNum())
		}

		prev := 0
		for i := len(lastNumbers) - 1; i >= 0; i-- {
			n := lastNumbers[i]
			if left {
				prev = n - prev
			} else {
				prev += n
			}
		}
		sum += prev
	}

	return strconv.Itoa(sum), nil
}
