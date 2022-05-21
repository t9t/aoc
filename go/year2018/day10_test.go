package year2018

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var day10input = `
position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>
`

var day10LetterMap = map[string]string{
	"#...#.\n#...#.\n#...#.\n#####.\n#...#.\n#...#.\n#...#.\n#...#.\n": "H",
	"###\n.#.\n.#.\n.#.\n.#.\n.#.\n.#.\n###\n":                         "I",
}

func Test_Day10Part1(t *testing.T) {
	actual, _, err := day104test()
	require.Nil(t, err)
	assert.Equal(t, "HI", actual)
}

func Test_Day10Part2(t *testing.T) {
	_, actual, err := day104test()
	require.Nil(t, err)
	assert.Equal(t, 3, actual)
}

func day104test() (string, int, error) {
	return day10(strings.TrimSpace(day10input), 7, 8, day10LetterMap)
}
