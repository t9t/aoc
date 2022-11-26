package year2018

import (
	"fmt"
	"regexp"
	"strconv"
)

func init() {
	mustRegisterPair(23, Day23Part1, Day23Part2)
}

func Day23Part1(input string) (string, error) {
	re := regexp.MustCompile(`pos=<(-?\d+),(-?\d+),(-?\d+)>, r=(\d+)`)
	type bot struct{ x, y, z, r int }
	bots := make([]bot, 0)
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		nums := make([]int, 4)
		for i, s := range match[1:] {
			if n, err := strconv.Atoi(s); err != nil {
				return "", err
			} else {
				nums[i] = n
			}
		}
		bots = append(bots, bot{x: nums[0], y: nums[1], z: nums[2], r: nums[3]})
	}

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	var strongest bot
	for _, b := range bots {
		if b.r > strongest.r {
			strongest = b
		}
	}

	inRange := 0
	for _, b := range bots {
		if abs(strongest.x-b.x)+abs(strongest.y-b.y)+abs(strongest.z-b.z) <= strongest.r {
			inRange++
		}
	}

	return strconv.Itoa(inRange), nil
}

func Day23Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 23 part 2 not implemented")
}
