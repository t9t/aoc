package year2018

import (
	"regexp"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(16, Day16Part1, Day16Part2)
}

func Day16Part1(input string) (string, error) {
	return day16(input, false)
}

func Day16Part2(input string) (string, error) {
	return day16(input, true)
}

func day16(input string, part2 bool) (string, error) {
	type operation func(a, b, c int, in []int)
	addr := func(a, b, c int, in []int) {
		in[c] = in[a] + in[b]
	}

	addi := func(a, b, c int, in []int) {
		in[c] = in[a] + b
	}

	mulr := func(a, b, c int, in []int) {
		in[c] = in[a] * in[b]
	}

	muli := func(a, b, c int, in []int) {
		in[c] = in[a] * b
	}

	banr := func(a, b, c int, in []int) {
		in[c] = in[a] & in[b]
	}

	bani := func(a, b, c int, in []int) {
		in[c] = in[a] & b
	}

	borr := func(a, b, c int, in []int) {
		in[c] = in[a] | in[b]
	}

	bori := func(a, b, c int, in []int) {
		in[c] = in[a] | b
	}

	setr := func(a, b, c int, in []int) {
		in[c] = in[a]
	}

	seti := func(a, b, c int, in []int) {
		in[c] = a
	}

	gtir := func(a, b, c int, in []int) {
		if a > in[b] {
			in[c] = 1
		} else {
			in[c] = 0
		}
	}

	gtri := func(a, b, c int, in []int) {
		if in[a] > b {
			in[c] = 1
		} else {
			in[c] = 0
		}
	}

	gtrr := func(a, b, c int, in []int) {
		if in[a] > in[b] {
			in[c] = 1
		} else {
			in[c] = 0
		}
	}

	eqir := func(a, b, c int, in []int) {
		if a == in[b] {
			in[c] = 1
		} else {
			in[c] = 0
		}
	}

	eqri := func(a, b, c int, in []int) {
		if in[a] == b {
			in[c] = 1
		} else {
			in[c] = 0
		}
	}

	eqrr := func(a, b, c int, in []int) {
		if in[a] == in[b] {
			in[c] = 1
		} else {
			in[c] = 0
		}
	}

	operations := []operation{
		addr, addi,
		mulr, muli,
		banr, bani,
		borr, bori,
		setr, seti,
		gtir, gtri, gtrr,
		eqir, eqri, eqrr,
	}

	equalArrays := func(l, r []int) bool {
		for i, lv := range l {
			if r[i] != lv {
				return false
			}
		}
		return true
	}

	re := regexp.MustCompile(`.*?(\d+),? (\d+),? (\d+),? (\d+)`)
	parseNumbers := func(line string) ([]int, error) {
		matches := re.FindStringSubmatch(line)
		nrs := make([]int, 4)
		for i := range nrs {
			if n, err := strconv.Atoi(matches[i+1]); err != nil {
				return nil, err
			} else {
				nrs[i] = n
			}
		}
		return nrs, nil
	}

	sections := strings.Split(input, "\n\n\n\n")
	samples := strings.Split(sections[0], "\n\n")

	behaveLikeThreeOrMoreCount := 0
	operationMap := make(map[int]operation)
	for len(operations) > 0 {
		for _, sample := range samples {
			var err error
			var before, instruction, after []int
			lines := strings.Split(sample, "\n")
			if before, err = parseNumbers(lines[0]); err != nil {
				return "", err
			} else if instruction, err = parseNumbers(lines[1]); err != nil {
				return "", err
			} else if after, err = parseNumbers(lines[2]); err != nil {
				return "", err
			}

			if _, found := operationMap[instruction[0]]; part2 && found {
				continue
			}

			matchingOperations := 0
			var matchingOperation operation
			var matchingIndex int
			for i, operation := range operations {
				out := make([]int, 4)
				copy(out, before)
				operation(instruction[1], instruction[2], instruction[3], out)
				if equalArrays(out, after) {
					matchingOperations++
					if part2 && matchingOperations > 1 {
						break
					}
					matchingOperation, matchingIndex = operation, i
				}
			}

			if part2 && matchingOperations == 1 {
				operationMap[instruction[0]] = matchingOperation
				operations = append(operations[:matchingIndex], operations[matchingIndex+1:]...)
			} else if !part2 && matchingOperations >= 3 {
				behaveLikeThreeOrMoreCount++
			}
		}
		if !part2 {
			return strconv.Itoa(behaveLikeThreeOrMoreCount), nil
		}
	}

	registers := make([]int, 4)
	for _, line := range strings.Split(sections[1], "\n") {
		instruction, err := parseNumbers(line)
		if err != nil {
			return "", err
		}
		operationMap[instruction[0]](instruction[1], instruction[2], instruction[3], registers)
	}

	return strconv.Itoa(registers[0]), nil
}
