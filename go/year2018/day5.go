package year2018

import (
	"fmt"
	"strconv"
)

func init() {
	mustRegisterPair(5, Day5Part1, Day5Part2)
}

func Day5Part1(input string) (string, error) {
	diff := byte('a') - byte('A')
	for {
		anyReduced := false
		for i := 0; i < len(input)-1; i++ {
			left := input[i]
			right := input[i+1]

			if left+diff == right || left-diff == right {
				input = input[:i] + input[i+2:]
				anyReduced = true
				i++
			}
		}
		if !anyReduced {
			break
		}
	}

	return strconv.Itoa(len(input)), nil
}

func Day5Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 5 part 2 not implemented")
}
