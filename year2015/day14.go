package year2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day14Part1(input string) (int, error) {
	return findMaxDistanceTraveled(input, 2503)
}

func findMaxDistanceTraveled(input string, seconds int) (int, error) {
	reindeer, err := parseReindeerDescriptions(input)
	if err != nil {
		return 0, err
	}

	max := 0
	for _, r := range reindeer {
		dist := r.distanceTraveledAfter(seconds)
		if dist > max {
			max = dist
		}
	}

	return max, nil
}

func Day14Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

type reindeerDescription struct {
	flySpeed, flyTime, restTime int
}

func (r reindeerDescription) distanceTraveledAfter(seconds int) int {
	cycleTime := r.flyTime + r.restTime
	fullCycles := seconds / cycleTime
	fullCyclesDistance := r.flySpeed * r.flyTime * fullCycles

	remainderTime := seconds % cycleTime
	remainderFlyTime := remainderTime
	if remainderFlyTime > r.flyTime {
		remainderFlyTime = r.flyTime
	}
	remainderDistance := remainderFlyTime * r.flySpeed

	return fullCyclesDistance + remainderDistance
}

var reindeerDescriptionRegexp = regexp.MustCompile(`\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

func parseReindeerDescriptions(input string) ([]reindeerDescription, error) {
	allMatches := reindeerDescriptionRegexp.FindAllStringSubmatch(strings.TrimSpace(input), -1)
	if len(allMatches) == 0 {
		return nil, fmt.Errorf("invalid reindeer descriptions input")
	}
	out := make([]reindeerDescription, len(allMatches))
	for i, matches := range allMatches {
		flySpeed, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("unparseable reindeer fly speed: %w", err)
		}
		flyTime, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("unparseable reindeer fly time: %w", err)
		}
		restTime, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, fmt.Errorf("unparseable reindeer rest time: %w", err)
		}
		out[i] = reindeerDescription{flySpeed: flySpeed, flyTime: flyTime, restTime: restTime}
	}

	return out, nil
}
