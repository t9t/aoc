package year2018

import (
	"strings"
	"testing"
)

func Test_Day12Part1(t *testing.T) {
	input := `
initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
`
	basicTest(t, Day12Part1, strings.TrimSpace(input), "325")
}
