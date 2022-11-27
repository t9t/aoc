package year2018

import (
	"aoc/registry"
	"regexp"
	"strconv"
)

func init() {
	registry.MustRegister(2018, 25, 1, Day25Part1)
}

func Day25Part1(input string) (string, error) {
	type point []int

	re := regexp.MustCompile(`(-?\d+),(-?\d+),(-?\d+),(-?\d+)`)
	points := make([]point, 0)
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		nums := make(point, 4)
		for i, s := range match[1:] {
			if n, err := strconv.Atoi(s); err != nil {
				return "", err
			} else {
				nums[i] = n
			}
		}
		points = append(points, nums)
	}

	constellations := make([]*[]point, 0)

	for iter := 0; len(points) != 0; iter++ {
		currentConstellation := []point{points[0]}
		constellations = append(constellations, &currentConstellation)
		currentMatch := 0
		points = points[1:]
		if len(points) == 0 {
			break
		}

		for {
			leftPoints := make([]point, 0)
			match := currentConstellation[currentMatch]
			for _, other := range points {
				dist := 0
				for i, l := range match {
					d := l - other[i]
					if d < 0 {
						d = -d
					}
					dist += d
				}
				if dist <= 3 {
					currentConstellation = append(currentConstellation, other)
				} else {
					leftPoints = append(leftPoints, other)
				}
			}
			points = leftPoints
			currentMatch += 1
			if currentMatch >= len(currentConstellation) {
				break
			}
		}
	}

	return strconv.Itoa(len(constellations)), nil
}
