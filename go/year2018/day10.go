package year2018

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(10, Day10Part1, Day10Part2)
}

func Day10Part1(input string) (string, error) {
	s, _, err := day104real(input)
	return s, err
}

func Day10Part2(input string) (string, error) {
	_, t, err := day104real(input)
	if err != nil {
		return "", err
	} else {
		return strconv.Itoa(t), nil
	}
}

func day104real(input string) (string, int, error) {
	letterMap := map[string]string{
		"######\n#.....\n#.....\n#.....\n#####.\n#.....\n#.....\n#.....\n#.....\n######\n": "E",
		"#....#\n#...#.\n#..#..\n#.#...\n##....\n##....\n#.#...\n#..#..\n#...#.\n#....#\n": "K",
		"#.....\n#.....\n#.....\n#.....\n#.....\n#.....\n#.....\n#.....\n#.....\n######\n": "L",
		"#....#\n##...#\n##...#\n#.#..#\n#.#..#\n#..#.#\n#..#.#\n#...##\n#...##\n#....#\n": "N",
		"#####.\n#....#\n#....#\n#....#\n#####.\n#.....\n#.....\n#.....\n#.....\n#.....\n": "P",
		"#####.\n#....#\n#....#\n#....#\n#####.\n#..#..\n#...#.\n#...#.\n#....#\n#....#\n": "R",
		"#....#\n#....#\n.#..#.\n.#..#.\n..##..\n..##..\n.#..#.\n.#..#.\n#....#\n#....#\n": "X",
	}

	return day10(input, 8, 10, letterMap)
}

func day10(input string, presumedLetterWidth, presumedLetterHeight int, letterMap map[string]string) (string, int, error) {
	type xAndY struct{ x, y int }

	parseXAndY := func(s string) (r xAndY, err error) {
		parts := strings.Split(s, ",")
		if r.x, err = strconv.Atoi(strings.TrimSpace(parts[0])); err != nil {
			return
		}
		r.y, err = strconv.Atoi(strings.TrimSpace(parts[1]))
		return
	}

	positions := make([]*xAndY, 0)
	velocities := make([]xAndY, 0)
	for _, line := range strings.Split(input, "\n") {
		lr := strings.Split(line, "> velocity=<")
		pos, err := parseXAndY(strings.ReplaceAll(lr[0], "position=<", ""))
		if err != nil {
			return "", 0, fmt.Errorf("invalid line %v: %w", line, err)
		}
		velo, err := parseXAndY(strings.ReplaceAll(lr[1], ">", ""))
		if err != nil {
			return "", 0, fmt.Errorf("invalid line %v: %w", line, err)
		}
		positions = append(positions, &pos)
		velocities = append(velocities, velo)
	}

	findingHeight := presumedLetterHeight - 1
	for s := 1; s <= 1_000_000; s++ {
		for i, pos := range positions {
			velo := velocities[i]
			pos.x += velo.x
			pos.y += velo.y
		}

		isMaxLetterHeight := func() bool {
			maximumDy := 0
			for i, l := range positions {
				for j, r := range positions {
					if i != j {
						dy := l.y - r.y
						if dy < 0 {
							dy = -dy
						}
						if dy > maximumDy {
							maximumDy = dy
						}
						if maximumDy > findingHeight {
							return false
						}
					}
				}
			}
			return true
		}()
		if isMaxLetterHeight {
			min, max := xAndY{x: math.MaxInt, y: math.MaxInt}, xAndY{x: math.MinInt, y: math.MinInt}
			posMap := make(map[xAndY]struct{})
			for _, p := range positions {
				posMap[*p] = struct{}{}
				if p.x < min.x {
					min.x = p.x
				}
				if p.y < min.y {
					min.y = p.y
				}
				if p.x > max.x {
					max.x = p.x
				}
				if p.y > max.y {
					max.y = p.y
				}
			}

			out := make([][][]string, (max.x-min.x)/presumedLetterWidth+1)
			for i := 0; i < len(out); i++ {
				out[i] = make([][]string, presumedLetterHeight)
				for j := 0; j < len(out[i]); j++ {
					out[i][j] = make([]string, presumedLetterWidth)
				}
			}
			for y := min.y; y <= max.y; y++ {
				for x := min.x; x <= max.x; x++ {
					char := "."
					if _, set := posMap[xAndY{x: x, y: y}]; set {
						char = "#"
					}
					adjustedX, adjustedY := x-min.x, y-min.y
					out[adjustedX/presumedLetterWidth][adjustedY][adjustedX%presumedLetterWidth] = char
				}
			}

			var word strings.Builder
			for _, letter := range out {
				var letterBuilder strings.Builder
				for _, line := range letter {
					letterBuilder.WriteString(strings.Join(line[:6], ""))
					letterBuilder.WriteString("\n")
				}
				lg := letterBuilder.String()
				if m, found := letterMap[lg]; found {
					word.WriteString(m)
				} else {
					return "", 0, fmt.Errorf("letter could not be mapped: %s", lg)
				}
			}

			return word.String(), s, nil
		}
	}

	return "", 0, fmt.Errorf("no solution found")
}
