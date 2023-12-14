package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(14, Day14Part1, Day14Part2)
}

func Day14Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	for x := 0; x < len(lines[0]); x += 1 {
		for y := 0; y < len(lines); y += 1 {
			line := lines[y]
			c := line[x]
			if c != 'O' || y == 0 || lines[y-1][x] != '.' {
				continue
			}

			targety := y - 1
			for oy := y - 1; oy >= 0; oy -= 1 {
				if lines[oy][x] == '.' {
					targety = oy
				} else {
					break // O or #
				}
			}
			lineBytes := []byte(line)
			lineBytes[x] = '.'
			lines[y] = string(lineBytes)

			lineBytes = []byte(lines[targety])
			lineBytes[x] = 'O'
			lines[targety] = string(lineBytes)
		}
	}

	sum := 0
	for y, line := range lines {
		load := len(lines) - y
		for _, r := range line {
			if r == 'O' {
				sum += load
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func Day14Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	lineCount, lineLen := len(lines), len(lines[0])

	north := func() {
		for x := 0; x < len(lines[0]); x += 1 {
			for y := 0; y < len(lines); y += 1 {
				line := lines[y]
				c := line[x]
				if c != 'O' || y == 0 || lines[y-1][x] != '.' {
					continue
				}

				targety := y - 1
				for oy := y - 1; oy >= 0; oy -= 1 {
					if lines[oy][x] == '.' {
						targety = oy
					} else {
						break // O or #
					}
				}
				lineBytes := []byte(line)
				lineBytes[x] = '.'
				lines[y] = string(lineBytes)

				lineBytes = []byte(lines[targety])
				lineBytes[x] = 'O'
				lines[targety] = string(lineBytes)
			}
		}
	}

	west := func() {
		for y := 0; y < lineCount; y += 1 {
			for x := 0; x < lineLen; x += 1 {
				line := lines[y]
				c := line[x]
				if c != 'O' || x == 0 || lines[y][x-1] != '.' {
					continue
				}

				targetx := x - 1
				for ox := x - 1; ox >= 0; ox -= 1 {
					if lines[y][ox] == '.' {
						targetx = ox
					} else {
						break // O or #
					}
				}
				lineBytes := []byte(line)
				lineBytes[x] = '.'
				lineBytes[targetx] = 'O'
				lines[y] = string(lineBytes)
			}
		}
	}

	south := func() {
		for x := 0; x < lineLen; x += 1 {
			for y := lineCount - 1; y >= 0; y -= 1 {
				line := lines[y]
				c := line[x]
				if c != 'O' || y == lineCount-1 || lines[y+1][x] != '.' {
					continue
				}

				targety := y + 1
				for oy := y + 1; oy < lineCount; oy += 1 {
					if lines[oy][x] == '.' {
						targety = oy
					} else {
						break // O or #
					}
				}
				lineBytes := []byte(line)
				lineBytes[x] = '.'
				lines[y] = string(lineBytes)

				lineBytes = []byte(lines[targety])
				lineBytes[x] = 'O'
				lines[targety] = string(lineBytes)
			}
		}
	}

	east := func() {
		for y := 0; y < lineCount; y += 1 {
			for x := lineLen - 1; x >= 0; x -= 1 {
				line := lines[y]
				c := line[x]
				if c != 'O' || x == lineLen-1 || lines[y][x+1] != '.' {
					continue
				}

				targetx := x + 1
				for ox := x + 1; ox < lineLen; ox += 1 {
					if lines[y][ox] == '.' {
						targetx = ox
					} else {
						break // O or #
					}
				}
				lineBytes := []byte(line)
				lineBytes[x] = '.'
				lineBytes[targetx] = 'O'
				lines[y] = string(lineBytes)
			}
		}
	}

	previous := make([]int, 0)
	// Assumptions: we can find the answer in < 1000 iterations, and a pattern is never longer than 100 integers long
	for i := 0; i < 1_000; i += 1 {
		north()
		west()
		south()
		east()

		sum := 0
		for y, line := range lines {
			load := len(lines) - y
			for _, r := range line {
				if r == 'O' {
					sum += load
				}
			}
		}
		previous = append(previous, sum)

		for l := 2; l <= 100; l += 1 {
			plen, twice := len(previous), l*2
			if len(previous) < twice {
				break
			}

			last := previous[plen-l:]
			beforeLast := previous[plen-l*2 : plen-l]
			same := true

			for j, a := range last {
				if a != beforeLast[j] {
					same = false
					break
				}
			}

			if same {
				offset := plen - l*2
				return strconv.Itoa(previous[offset+((1_000_000_000-offset)%l)-1]), nil
			}
		}
	}
	return "", fmt.Errorf("no answer found")
}
