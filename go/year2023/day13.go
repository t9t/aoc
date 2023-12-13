package year2023

import (
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(13, Day13Part1, Day13Part2)
}

func Day13Part1(input string) (string, error) {
	chunks := strings.Split(input, "\n\n")

	checkReflection := func(maxNum int, f func(int) string) (bool, int) {
		for x := 1; x < maxNum; x++ {
			if a, b := f(x-1), f(x); a == b {
				offset := 1
				for {
					leftX, rightX := x-1-offset, x+offset
					if leftX < 0 || rightX == maxNum {
						return true, x
					}
					if left, right := f(leftX), f(rightX); left == right {
						offset += 1
						continue
					}
					break
				}
			}
		}
		return false, 0
	}

	findReflectionXy := func(chunk string) (int, int) {
		lines := strings.Split(chunk, "\n")

		xMatch, xReflection := checkReflection(len(lines[0]), func(x int) string {
			var sb strings.Builder
			for _, line := range lines {
				sb.WriteByte(line[x])
			}
			return sb.String()
		})

		if xMatch {
			// Assumption: a chunk cannot have both a horizontal and vertical reflection
			return xReflection, 0
		}

		_, yReflection := checkReflection(len(lines), func(y int) string {
			return lines[y]
		})
		return 0, yReflection
	}

	summary := 0
	for _, chunk := range chunks {
		x, y := findReflectionXy(chunk)
		summary += x
		summary += y * 100
	}

	return strconv.Itoa(summary), nil
}

func Day13Part2(input string) (string, error) {
	chunks := strings.Split(input, "\n\n")

	checkReflection := func(maxNum int, notN int, f func(int) string) (bool, int) {
		for x := 1; x < maxNum; x++ {
			if x == notN {
				continue
			}
			if a, b := f(x-1), f(x); a == b {
				offset := 1
				for {
					leftX, rightX := x-1-offset, x+offset
					if leftX < 0 || rightX == maxNum {
						return true, x
					}
					if left, right := f(leftX), f(rightX); left == right {
						offset += 1
						continue
					}
					break
				}
			}
		}
		return false, 0
	}

	findReflectionXy := func(lines []string, notX, notY int) (int, int, bool) {
		yMatch, yReflection := checkReflection(len(lines), notY, func(y int) string {
			return lines[y]
		})

		if yMatch {
			// Assumption: a chunk cannot have both a horizontal and vertical reflection
			return 0, yReflection, true
		}

		xMatch, xReflection := checkReflection(len(lines[0]), notX, func(x int) string {
			var sb strings.Builder
			for _, line := range lines {
				sb.WriteByte(line[x])
			}
			return sb.String()
		})

		if xMatch {
			return xReflection, 0, true
		}
		return 0, 0, false
	}
	summary := 0
kek:
	for _, chunk := range chunks {

		origLines := strings.Split(chunk, "\n")
		origX, origY, origMatch := findReflectionXy(origLines, -1, -1)
		if !origMatch {
			panic("")
		}

		for sy := 0; sy < len(origLines); sy += 1 {
			for sx := 0; sx < len(origLines[0]); sx += 1 {
				copyLines := make([]string, len(origLines))
				copy(copyLines, origLines)
				line := []byte(copyLines[sy])
				if line[sx] == '.' {
					line[sx] = '#'
				} else {
					line[sx] = '.'
				}
				copyLines[sy] = string(line)

				fixedX, fixedY, match := findReflectionXy(copyLines, origX, origY)
				diff := chunk != strings.TrimSpace(strings.Join(copyLines, "\n"))
				if match && diff && (fixedX != origX || fixedY != origY) {
					func(int, int) {}(origX, origY)

					summary += fixedX
					summary += fixedY * 100
					continue kek
				}
			}
		}
		panic("no answer")
	}

	return strconv.Itoa(summary), nil
}
