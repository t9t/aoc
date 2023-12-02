package year2023

import (
	"testing"
)

func Test_Day1Part1(t *testing.T) {
	basicMultiTest(t, Day1Part1, []testInput{
		{`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`, "142"},
	})
}

func Test_Day1Part2(t *testing.T) {
	basicMultiTest(t, Day1Part2, []testInput{
		{`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`, "281"},
	})
}
