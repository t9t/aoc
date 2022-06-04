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
	return day13(input, true)
}

func Day13Part2(input string) (string, error) {
	return day13(input, false)
}

func day13(input string, returnFirstCrash bool) (string, error) {
	type cart struct {
		current, next byte
		x, y          int
	}

	carts := make([]*cart, 0)
	tracks := make([][]byte, 0)

	for y, line := range strings.Split(input, "\n") {
		row := []byte(line)
		for x, c := range row {
			if c == '>' || c == '<' {
				row[x] = '-'
			} else if c == '^' || c == 'v' {
				row[x] = '|'
			}
			if c != row[x] {
				carts = append(carts, &cart{current: c, next: 'l', x: x, y: y})
			}
		}
		tracks = append(tracks, row)
	}

	backSlash := map[byte]byte{'<': '^', '>': 'v', '^': '<', 'v': '>'}
	forwardSlash := map[byte]byte{'<': 'v', '>': '^', '^': '>', 'v': '<'}
	rightTurns := map[byte]byte{'<': '^', '>': 'v', '^': '>', 'v': '<'}
	leftTurns := map[byte]byte{'<': 'v', '>': '^', '^': '<', 'v': '>'}

	maxIter := 100_000
	for i := 0; i < maxIter; i++ {
		oldCarts := carts
		carts = carts[0:0]
		for _, cart := range oldCarts {
			if cart != nil {
				carts = append(carts, cart)
			}
		}

		if len(carts) == 1 {
			return fmt.Sprintf("%d,%d", carts[0].x, carts[0].y), nil
		}

		sort.Slice(carts, func(i, j int) bool {
			l, r := carts[i], carts[j]
			if l.y == r.y {
				return l.x < r.x
			}
			return l.y < r.y
		})

		for ci, cart := range carts {
			if cart == nil {
				continue
			}

			nx, ny := cart.x, cart.y
			if cart.current == '>' {
				nx++
			} else if cart.current == '<' {
				nx--
			} else if cart.current == 'v' {
				ny++
			} else if cart.current == '^' {
				ny--
			}

			for oi, other := range carts {
				if other != nil && other.x == nx && other.y == ny {
					carts[ci], carts[oi] = nil, nil
					if returnFirstCrash {
						return fmt.Sprintf("%d,%d", nx, ny), nil
					}
				}
			}
			if carts[ci] == nil {
				continue
			}
			cart.x, cart.y = nx, ny

			c := tracks[ny][nx]
			if c == '\\' {
				cart.current = backSlash[cart.current]
			} else if c == '/' {
				cart.current = forwardSlash[cart.current]
			} else if c == '+' {
				if cart.next == 'l' {
					cart.next = 's'
					cart.current = leftTurns[cart.current]
				} else if cart.next == 's' {
					cart.next = 'r'
				} else if cart.next == 'r' {
					cart.next = 'l'
					cart.current = rightTurns[cart.current]
				}
			} else if c != '-' && c != '|' {
				panic(fmt.Sprintf("wrong turn c: %v", c))
			}
		}
	}

	return "", fmt.Errorf("no answer found after %d iterations", maxIter)
}
