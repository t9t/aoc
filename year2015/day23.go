package year2015

import (
	"fmt"
	"strconv"
	"strings"
)

func Day23Part1(input string) (int, error) {
	return day23(input, 0)
}

func Day23Part2(input string) (int, error) {
	return day23(input, 1)
}

func day23(input string, regAStartingValue int) (int, error) {
	instructions, err := parseInstructions(input)
	if err != nil {
		return 0, err
	}
	return runProgram(instructions, regAStartingValue), nil
}

func parseInstructions(input string) ([]instruction, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	out := make([]instruction, len(lines))
	for i, line := range lines {
		ins, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}
		out[i] = ins
	}
	return out, nil
}

func parseInstruction(s string) (out instruction, err error) {
	parts := strings.Split(strings.TrimSpace(strings.ReplaceAll(s, ",", "")), " ")
	if len(parts) < 2 || len(parts) > 3 {
		return out, fmt.Errorf("invalid instruction %q: %d parts", s, len(parts))
	}
	ins := parts[0]
	arg := 0
	if strings.HasPrefix(ins, "ji") {
		if len(parts) != 3 {
			return out, fmt.Errorf("invalid jie/jio instruction %q", s)
		} else if arg, err = strconv.Atoi(parts[2]); err != nil {
			return out, fmt.Errorf("invalid jie/jio instruction arg %q: %w", s, err)
		}
	} else {
		if len(parts) != 2 {
			return out, fmt.Errorf("invalid instruction %q", s)
		}
	}
	regString := ""
	if ins == "jmp" {
		if arg, err = strconv.Atoi(parts[1]); err != nil {
			return out, fmt.Errorf("invalid jmp instruction arg %q: %w", s, err)
		}
	} else {
		regString = parts[1]
		if regString != "a" && regString != "b" {
			return out, fmt.Errorf("invalid instruction %q: unknown register %q", s, regString)
		}
	}

	return instruction{instruction: ins, regA: regString == "a", arg: arg}, nil
}

type instruction struct {
	instruction string
	regA        bool
	arg         int
}

func runProgram(instructions []instruction, regAStartingValue int) int {
	a, b := regAStartingValue, 0

	p := 0
	l := len(instructions)
	iters := 0
	for p < l {
		iters++
		instruction := instructions[p]
		regA := instruction.regA
		arg := instruction.arg
		jump := 1
		switch instruction.instruction {
		case "hlf":
			if regA {
				a = a / 2
			} else {
				b = b / 2
			}
		case "tpl":
			if regA {
				a = a * 3
			} else {
				b = b * 3
			}
		case "inc":
			if regA {
				a++
			} else {
				b++
			}
		case "jmp":
			jump = arg
		case "jie":
			v := b
			if regA {
				v = a
			}
			if v%2 == 0 {
				jump = arg
			}
		case "jio":
			v := b
			if regA {
				v = a
			}
			if v == 1 {
				jump = arg
			}
		}
		p += jump
	}
	return b
}
