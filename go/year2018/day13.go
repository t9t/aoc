package year2018

import (
	"fmt"
	"sort"
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
	type cart struct {
		counter, current, next byte
		x, y                   int
		alive                  bool
	}

	carts := make([]*cart, 0)
	initialMap := make([][]byte, 0)
	counter := byte(1)

	for y, line := range strings.Split(input, "\n") {
		initialMapRow := make([]byte, 0)
		for x, c := range []byte(line) {
			in := c
			if in == '>' || in == '<' {
				in = '-'
			} else if in == '^' || in == 'v' {
				in = '|'
			}
			initialMapRow = append(initialMapRow, in)

			if c != in {
				carts = append(carts, &cart{counter: counter, current: c, next: 'l', x: x, y: y, alive: true})
				counter++
			}
		}
		initialMap = append(initialMap, initialMapRow)
	}

	findAt := func(carts []*cart, x, y int) (found bool, c *cart) {
		for _, c = range carts {
			if !c.alive {
				continue
			}
			if c.x == x && c.y == y {
				return true, c
			}
		}
		return false, c
	}

	for i := 0; i < 1_000_000; i++ {
		sort.Slice(carts, func(i, j int) bool {
			l, r := carts[i], carts[j]
			if l.y == r.y {
				return l.x < r.x
			}
			return l.y < r.y
		})
		for _, cart := range carts {
			if !cart.alive {
				continue
			}
			x, y := cart.x, cart.y
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
			if otherFound, other := findAt(carts, nx, ny); otherFound {
				fmt.Printf("~Crash @%d at %d,%d / %d,%d of %v and %v\n", i, x, y, nx, ny, cart, other)
				cart.alive = false
				other.alive = false
				continue
			}

			nextC := initialMap[ny][nx]
			cart.x = nx
			cart.y = ny

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
				panic(fmt.Sprintf("i: %d; nextC: #%c#; cart: %+v", i, nextC, cart))
			}
		}
		aliveCount := 0
		var aliveCart *cart
		for _, c := range carts {
			if c.alive {
				aliveCart = c
				aliveCount++
				if aliveCount > 1 {
					break
				}
			}
		}
		if aliveCount == 1 {
			return fmt.Sprintf("%d,%d", aliveCart.x, aliveCart.y), nil
		}
	}

	return "", fmt.Errorf("no answer found after 1000 iterations")
}
