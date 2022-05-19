package year2018

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(8, Day8Part1, Day8Part2)
}

/*
2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2
A----------------------------------
    B----------- C-----------
                     D-----
*/
func Day8Part1(input string) (string, error) {
	parts := strings.Split(input, " ")
	numbers := make([]int, len(parts))
	for i, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			return "", err
		}
		numbers[i] = n
	}

	sum, _ := day8{}.sumNextNode(numbers, 0)
	return strconv.Itoa(sum), nil
}

func Day8Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 8 part 2 not implemented")
}

type day8 struct{}

func (d day8) sumNextNode(numbers []int, depth int) (int, []int) {
	childCount, metadataCount := numbers[0], numbers[1]
	numbers = numbers[2:]

	sum := 0
	for ; childCount > 0; childCount-- {
		nextSum, nextNumbers := d.sumNextNode(numbers, depth+1)
		sum += nextSum
		numbers = nextNumbers
	}
	for ; metadataCount > 0; metadataCount-- {
		sum += numbers[0]
		numbers = numbers[1:]
	}
	return sum, numbers
}
