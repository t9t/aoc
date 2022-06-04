package year2018

import (
	"fmt"
	"strings"
)

func init() {
	mustRegisterPair(13, Day13Part1, Day13Part2)
}

func Day13Part1(input string) (string, error) {
	type cart struct{ current, next byte }

	carts := make([]*cart, 0)
	initialMap := make([][]byte, 0)
	cartMap := make([][]int, 0)

	for _, line := range strings.Split(input, "\n") {
		initialMapRow := make([]byte, 0)
		cartMapRow := make([]int, 0)
		for _, c := range []byte(line) {
			in := c
			if in == '>' || in == '<' {
				in = '-'
			} else if in == '^' || in == 'v' {
				in = '|'
			}
			initialMapRow = append(initialMapRow, in)

			if c != in {
				cartNum := len(carts)
				carts = append(carts, &cart{current: c, next: 'l'})
				cartMapRow = append(cartMapRow, cartNum)
			} else {
				cartMapRow = append(cartMapRow, -1)
			}
		}
		initialMap = append(initialMap, initialMapRow)
		cartMap = append(cartMap, cartMapRow)
	}

	copyCartMap := func() [][]int {
		newMap := make([][]int, len(cartMap))
		for y, line := range cartMap {
			newRow := make([]int, len(line))
			for x := range line {
				newRow[x] = -1
			}
			newMap[y] = newRow
		}
		return newMap
	}

	for i := 0; i < 1000; i++ {
		newCartMap := copyCartMap()
		for y, line := range initialMap {
			for x := range line {
				cartNum := cartMap[y][x]
				if cartNum == -1 {
					continue
				}
				cart := carts[cartNum]
				dx, dy := 0, 0
				if cart.current == '<' {
					dx = -1
				} else if cart.current == '>' {
					dx = 1
				} else if cart.current == '^' {
					dy = -1
				} else if cart.current == 'v' {
					dy = 1
				}
				ny, nx := y+dy, x+dx
				if newCartMap[ny][nx] != -1 || cartMap[ny][nx] != -1 {
					return fmt.Sprintf("%d,%d", nx, ny), nil
				}
				newCartMap[ny][nx] = cartNum

				nextC := initialMap[y+dy][x+dx]

				if nextC == '\\' {
					if cart.current == '<' {
						cart.current = '^'
					} else if cart.current == '>' {
						cart.current = 'v'
					} else if cart.current == '^' {
						cart.current = '<'
					} else if cart.current == 'v' {
						cart.current = '>'
					} else {
						panic("")
					}
				} else if nextC == '/' {
					if cart.current == '<' {
						cart.current = 'v'
					} else if cart.current == '>' {
						cart.current = '^'
					} else if cart.current == '^' {
						cart.current = '>'
					} else if cart.current == 'v' {
						cart.current = '<'
					} else {
						panic("")
					}
				} else if nextC == '+' {
					if cart.next == 'l' {
						cart.next = 's'
						if cart.current == '<' {
							cart.current = 'v'
						} else if cart.current == '>' {
							cart.current = '^'
						} else if cart.current == '^' {
							cart.current = '<'
						} else if cart.current == 'v' {
							cart.current = '>'
						} else {
							panic("")
						}
					} else if cart.next == 's' {
						cart.next = 'r'
						// do not turn
					} else if cart.next == 'r' {
						cart.next = 'l'
						if cart.current == '<' {
							cart.current = '^'
						} else if cart.current == '>' {
							cart.current = 'v'
						} else if cart.current == '^' {
							cart.current = '>'
						} else if cart.current == 'v' {
							cart.current = '<'
						} else {
							panic("")
						}
					} else {
						panic("")
					}
				} else if nextC != '-' && nextC != '|' {
					panic("")
				}
			}
		}
		cartMap = newCartMap
	}

	return "", fmt.Errorf("no answer found after 1000 iterations")
}

func Day13Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 13 part 2 not implemented")
}
