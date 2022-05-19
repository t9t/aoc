package year2018

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(7, Day7Part1, Day7Part2)
}

func Day7Part1(input string) (string, error) {
	isBlockedBy := day7{}.buildIsBlockedBy(input)

	order := strings.Builder{}
	for len(isBlockedBy) > 0 {
		next := byte(0)
		for step, dependencies := range isBlockedBy {
			if len(dependencies) == 0 && (next == 0 || step < next) {
				next = step
			}
		}
		order.WriteByte(next)
		delete(isBlockedBy, next)
		for _, deps := range isBlockedBy {
			delete(deps, next)
		}
	}

	return order.String(), nil
}

func Day7Part2(input string) (string, error) {
	return day7Part2(input, 5, 60)
}

func day7Part2(input string, workerCount, extraStepDuration int) (string, error) {
	type worker struct {
		step     byte
		timeLeft int
	}

	isBlockedBy := day7{}.buildIsBlockedBy(input)

	workers := make([]*worker, workerCount)
	for i := range workers {
		workers[i] = &worker{}
	}

	time := 0
	anyWorking := true
	for len(isBlockedBy) > 0 || anyWorking {
		available := make([]byte, 0)
		for step, deps := range isBlockedBy {
			if len(deps) == 0 {
				available = append(available, step)
			}
		}
		sort.Slice(available, func(i, j int) bool { return available[i] < available[j] })
		anyWorking = false
		for _, worker := range workers {
			if worker.timeLeft > 0 {
				worker.timeLeft--
			}

			if worker.timeLeft == 0 {
				if worker.step != 0 {
					for _, deps := range isBlockedBy {
						delete(deps, worker.step)
					}
				}

				if len(available) != 0 {
					next := available[0]
					available = available[1:]
					worker.timeLeft = extraStepDuration + int(next-'A')
					worker.step = next
					delete(isBlockedBy, next)
				}
			}

			if worker.timeLeft > 0 {
				anyWorking = true
			}
		}

		time++
	}

	return strconv.Itoa(time), nil
}

type day7 struct{}

func (day7) buildIsBlockedBy(input string) map[byte]map[byte]struct{} {
	re := regexp.MustCompile(`(?m)Step (\w) must be finished before step (\w) can begin\.`)

	isBlockedBy := make(map[byte]map[byte]struct{})
	ensureEntry := func(step byte) {
		if _, found := isBlockedBy[step]; !found {
			isBlockedBy[step] = make(map[byte]struct{})
		}
	}

	for _, match := range re.FindAllStringSubmatch(input, -1) {
		first, then := match[1][0], match[2][0]
		ensureEntry(first)
		ensureEntry(then)
		isBlockedBy[then][first] = struct{}{}
	}
	return isBlockedBy
}
