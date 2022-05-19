package year2018

import (
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(8, Day8Part1, Day8Part2)
}

func Day8Part1(input string) (string, error) {
	d := &day8{}
	numbers, err := d.parseInput(input)
	if err != nil {
		return "", err
	}
	sum, _ := d.sumNextNode(numbers)
	return strconv.Itoa(sum), nil
}

func Day8Part2(input string) (string, error) {
	d := &day8{}
	numbers, err := d.parseInput(input)
	if err != nil {
		return "", err
	}
	sum, _ := d.sumNextNode2(numbers)
	return strconv.Itoa(sum), nil
}

type day8 struct{}

func (d *day8) parseInput(input string) ([]int, error) {
	parts := strings.Split(input, " ")
	numbers := make([]int, len(parts))
	for i, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		numbers[i] = n
	}
	return numbers, nil
}

func (d *day8) sumNextNode(numbers []int) (int, []int) {
	childCount, metadataCount := numbers[0], numbers[1]
	numbers = numbers[2:]

	sum := 0
	for ; childCount > 0; childCount-- {
		nextSum, nextNumbers := d.sumNextNode(numbers)
		sum += nextSum
		numbers = nextNumbers
	}
	for ; metadataCount > 0; metadataCount-- {
		sum += numbers[0]
		numbers = numbers[1:]
	}
	return sum, numbers
}

func (d *day8) sumNextNode2(numbers []int) (int, []int) {
	childCount, metadataCount := numbers[0], numbers[1]
	numbers = numbers[2:]

	if childCount == 0 {
		sum := 0
		for ; metadataCount > 0; metadataCount-- {
			sum += numbers[0]
			numbers = numbers[1:]
		}
		return sum, numbers
	}

	childSums := make([]int, childCount)
	for i := 0; i < childCount; i++ {
		nextSum, nextNumbers := d.sumNextNode2(numbers)
		numbers = nextNumbers
		childSums[i] = nextSum
	}

	sum := 0
	for ; metadataCount > 0; metadataCount-- {
		childIndex := numbers[0] - 1
		numbers = numbers[1:]
		if childIndex >= 0 && childIndex < len(childSums) {
			sum += childSums[childIndex]
		}
	}
	return sum, numbers
}
