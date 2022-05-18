package year2018

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func init() {
	mustRegisterPair(4, Day4Part1, Day4Part2)
}

func Day4Part1(input string) (string, error) {
	return day4(input, true)
}

func Day4Part2(input string) (string, error) {
	return day4(input, false)
}

func day4(input string, part1 bool) (string, error) {
	timeRe := regexp.MustCompile(`(?m)\[1518-(\d+)-(\d+) (\d+):(\d+)\] .+`)
	actionRe := regexp.MustCompile(`(?m).+\] (Guard #(\d+) begins shift|falls asleep|wakes up)`)

	parseTimeParts := func(match []string) (d, m, h, i int, err error) {
		if d, err = strconv.Atoi(match[1]); err != nil {
			return
		}
		if m, err = strconv.Atoi(match[2]); err != nil {
			return
		}
		if h, err = strconv.Atoi(match[3]); err != nil {
			return
		}
		i, err = strconv.Atoi(match[4])
		return
	}

	toTime := func(matches []string) (time.Time, error) {
		m, d, h, i, err := parseTimeParts(matches)
		if err != nil {
			return time.Time{}, err
		}
		return time.Date(1518, time.Month(m), d, h, i, 0, 0, time.UTC), nil
	}

	lines := strings.Split(input, "\n")
	sort.Slice(lines, func(i, j int) bool { return lines[i] < lines[j] })

	timeTables, totals := make(map[int]map[int]int), make(map[int]int)
	currentGuard := 0
	var fellAsleep time.Time
	maxOnAnyMinute, maxOnAnyMinuteGuardId, maxOnAnyMinuteMinute := 0, 0, 0
	maxSleeping, maxSleepingGuardId := 0, 0
	for _, line := range lines {
		timeMatches := timeRe.FindStringSubmatch(line)
		actionMatches := actionRe.FindStringSubmatch(line)

		if actionMatches[2] != "" {
			var err error
			currentGuard, err = strconv.Atoi(actionMatches[2])
			if err != nil {
				return "", fmt.Errorf("cannot parse guard number in %s: %w", line, err)
			}

			if _, hasTimeTable := timeTables[currentGuard]; !hasTimeTable {
				timeTable := make(map[int]int)
				for m := 0; m < 60; m++ {
					timeTable[m] = 0
				}
				timeTables[currentGuard] = timeTable
				totals[currentGuard] = 0
			}
			continue
		}

		timestamp, err := toTime(timeMatches)
		if err != nil {
			return "", fmt.Errorf("cannot parse timestamp in %s :%w", line, err)
		}

		if actionMatches[1] == "falls asleep" {
			fellAsleep = timestamp
		} else if actionMatches[1] == "wakes up" {
			timeTable := timeTables[currentGuard]
			for !fellAsleep.Equal(timestamp) {
				m := fellAsleep.Minute()
				minuteSleepCount := timeTable[m] + 1
				timeTable[m] = minuteSleepCount
				if minuteSleepCount > maxOnAnyMinute {
					maxOnAnyMinute = minuteSleepCount
					maxOnAnyMinuteGuardId = currentGuard
					maxOnAnyMinuteMinute = m
				}

				guardTotalSleepCount := totals[currentGuard] + 1
				totals[currentGuard] = guardTotalSleepCount
				if guardTotalSleepCount > maxSleeping {
					maxSleeping = guardTotalSleepCount
					maxSleepingGuardId = currentGuard
				}
				fellAsleep = fellAsleep.Add(time.Minute)
			}
		} else {
			return "", fmt.Errorf("invalid line: %s", line)
		}

	}

	if part1 {
		maxCount, mostSleepingMinute := 0, 0
		for minute, count := range timeTables[maxSleepingGuardId] {
			if count > maxCount {
				maxCount = count
				mostSleepingMinute = minute
			}
		}

		return strconv.Itoa(maxSleepingGuardId * mostSleepingMinute), nil
	}

	return strconv.Itoa(maxOnAnyMinuteMinute * maxOnAnyMinuteGuardId), nil
}
