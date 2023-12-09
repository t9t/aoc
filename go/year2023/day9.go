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
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		fmt.Printf("Line: %s\n", line)
		var numbers []int
		for _, part := range strings.Split(line, " ") {
			n, err := strconv.Atoi(part)
			if err != nil {
				return "", fmt.Errorf("invalid line %s: %w", line, err)
			}
			numbers = append(numbers, n)
		}
		fmt.Printf("\tnumbers: %#v\n", numbers)

		lastNumbers := []int{numbers[len(numbers)-1]}
		for {
			var next []int
			allZeroes := true
			for i := 0; i < len(numbers)-1; i++ {
				n := numbers[i]
				k := numbers[i+1]
				j := k - n
				next = append(next, j)
				if j != 0 {
					allZeroes = false
				}
			}

			fmt.Printf("\tNext: %#v (allZeroes: %t)\n", next, allZeroes)
			if allZeroes {
				break
			}
			lastNumbers = append(lastNumbers, next[len(next)-1])
			numbers = next
		}

		fmt.Printf("\tlastNumbers: %#v\n", lastNumbers)
		prev := 0
		for i := len(lastNumbers) - 1; i >= 0; i-- {
			n := lastNumbers[i]
			prev += n
		}
		fmt.Printf("\tPrev: %v\n", prev)
		sum += prev
	}

	fmt.Printf("Sum: %d\n", sum)

	return "", fmt.Errorf("Day 9 part 1 not implemented")
}

func Day9Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 9 part 2 not implemented")
}
