package year2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day12Part1(t *testing.T) {
	basicMultiTest(t, Day12Part1, []testInput{
		{`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`, "21"},
	})
}

func Test_day12line(t *testing.T) {
	tests := []struct {
		s    string
		nrs  []int
		want int
	}{
		{"???.###", []int{1, 1, 3}, 1},
		{".??..??...?##.", []int{1, 1, 3}, 4},
		{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 1},
		{"????.#...#...", []int{4, 1, 1}, 1},
		{"????.######..#####.", []int{1, 6, 5}, 4},
		{"?###????????", []int{3, 2, 1}, 10},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %v", tt.s, tt.nrs), func(t *testing.T) {
			got := day12line(tt.s, tt.nrs, 0)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_Day12Part2(t *testing.T) {
	basicMultiTest(t, Day12Part2, []testInput{
		{"", ""},
	})
}
