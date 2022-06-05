package year2018

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(14, Day14Part1, Day14Part2)
}

func Day14Part1(input string) (string, error) {
	numberOfRecipesTarget, err := strconv.Atoi(input)
	if err != nil {
		return "", err
	}

	scoreboard := []int{3, 7}
	elf1, elf2 := 0, 1

	breakTarget := numberOfRecipesTarget + 10

	for len(scoreboard) < breakTarget {
		elf1recipe, elf2recipe := scoreboard[elf1], scoreboard[elf2]
		newRecipe := elf1recipe + elf2recipe
		digit1, digit2 := newRecipe/10, newRecipe%10
		if digit1 != 0 {
			scoreboard = append(scoreboard, digit1)
		}
		scoreboard = append(scoreboard, digit2)
		elf1 = (elf1 + 1 + elf1recipe) % len(scoreboard)
		elf2 = (elf2 + 1 + elf2recipe) % len(scoreboard)
	}

	var out strings.Builder
	for _, n := range scoreboard[numberOfRecipesTarget:breakTarget] {
		out.WriteString(strconv.Itoa(n))
	}
	return out.String(), nil
}

func Day14Part2(input string) (string, error) {
	scoreboard := make([]int, 0, 25_000_000)
	scoreboard = append(scoreboard, 3)
	scoreboard = append(scoreboard, 7)
	elf1, elf2 := 0, 1

	needle := make([]int, 0)
	for _, b := range []byte(input) {
		needle = append(needle, int(b-'0'))
	}

	maxIter := 100_000_000
	for round := 0; round < maxIter; round++ {
		elf1recipe, elf2recipe := scoreboard[elf1], scoreboard[elf2]
		for _, c := range []byte(strconv.Itoa(elf1recipe + elf2recipe)) {
			scoreboard = append(scoreboard, int(c-'0'))
			if len(scoreboard) >= len(needle) {
				allMatch := true
				for i, n := range needle {
					if n != scoreboard[len(scoreboard)-len(needle)+i] {
						allMatch = false
						break
					}
				}
				if allMatch {
					return strconv.Itoa(len(scoreboard) - len(needle)), nil
				}
			}
		}

		elf1 = (elf1 + 1 + elf1recipe) % len(scoreboard)
		elf2 = (elf2 + 1 + elf2recipe) % len(scoreboard)
	}

	return "", fmt.Errorf("no answer found after %d iterations", maxIter)
}
