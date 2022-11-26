package year2018

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(19, Day19Part1, Day19Part2)
}

func Day19Part1(input string) (string, error) {
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

	operations := map[string]operation{
		"addr": addr,
		"addi": addi,
		"mulr": mulr,
		"muli": muli,
		"banr": banr,
		"bani": bani,
		"borr": borr,
		"bori": bori,
		"setr": setr,
		"seti": seti,
		"gtir": gtir,
		"gtri": gtri,
		"gtrr": gtrr,
		"eqir": eqir,
		"eqri": eqri,
		"eqrr": eqrr,
	}

	lines := strings.Split(input, "\n")
	registers := make([]int, 6)
	var ipRegister int
	var err error

	// Assumption: only one #ip line in the input, and it's the first line
	if ipRegister, err = strconv.Atoi(strings.Split(lines[0], "#ip ")[1]); err != nil {
		return "", fmt.Errorf("invalid '#ip' line %q: %w", lines[0], err)
	}
	lines = lines[1:]

	ip := 0
	for {
		if ip >= len(lines) {
			return strconv.Itoa(registers[0]), nil
		}
		line := lines[ip]

		parts := strings.Split(line, " ")
		op := parts[0]
		inputs := make([]int, 3)
		for i, s := range parts[1:] {
			if n, err := strconv.Atoi(s); err != nil {
				return "", fmt.Errorf("invalid line %q: %w", line, err)
			} else {
				inputs[i] = n
			}
		}

		registers[ipRegister] = ip
		operations[op](inputs[0], inputs[1], inputs[2], registers)
		ip = registers[ipRegister]

		ip += 1
	}
}

func Day19Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 19 part 2 not implemented")
}
