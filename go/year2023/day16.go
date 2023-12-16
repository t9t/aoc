package year2023

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	mustRegisterPair(16, Day16Part1, Day16Part2)
}

func Day16Part1(input string) (string, error) {
	type xy struct{ x, y int }
	type direction int
	const (
		up direction = iota
		right
		down
		left
	)
	lightMap := make(map[xy]map[direction]struct{})
	lines := strings.Split(input, "\n")
	xLen, yLen := len(lines[0]), len(lines)

	type beam struct {
		x, y int
		dir  direction
	}

	beams := []beam{{x: 0, y: 0, dir: right}}
	for len(beams) != 0 {
		b := beams[0]
		beams = beams[1:]

		for {
			bx, by, bdir, bxy := b.x, b.y, b.dir, xy{x: b.x, y: b.y}
			if _, found := lightMap[bxy][bdir]; found {
				break
			}
			if _, found := lightMap[bxy]; !found {
				lightMap[bxy] = make(map[direction]struct{})
			}
			lightMap[bxy][bdir] = struct{}{}

			tile := lines[by][bx]
			newdir := bdir
			if tile == '/' {
				switch bdir {
				case up:
					newdir = right
				case right:
					newdir = up
				case down:
					newdir = left
				case left:
					newdir = down
				}
			} else if tile == '\\' {
				switch bdir {
				case up:
					newdir = left
				case right:
					newdir = down
				case down:
					newdir = right
				case left:
					newdir = up
				}
			} else if tile == '-' {
				if bdir == up || bdir == down {
					// split up; this one goes right, add a new one going left
					newdir = right
					beams = append(beams, beam{x: bx, y: by, dir: left})
				} // else: pass through
			} else if tile == '|' {
				if bdir == left || bdir == right {
					// split up; this one goes up, add a new one going down
					newdir = up
					beams = append(beams, beam{x: bx, y: by, dir: down})
				} // else pass through
			}

			var dx, dy int
			switch newdir {
			case up:
				dx, dy = 0, -1
			case right:
				dx, dy = 1, 0
			case down:
				dx, dy = 0, 1
			case left:
				dx, dy = -1, 0
			}

			nx, ny := bx+dx, by+dy
			if nx < 0 || ny < 0 || nx == xLen || ny == yLen {
				break
			}
			b.x, b.y, b.dir = nx, ny, newdir
		}
	}

	return strconv.Itoa(len(lightMap)), nil
}

func Day16Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 16 part 2 not implemented")
}
