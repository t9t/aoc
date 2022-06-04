package year2018

import (
	"strings"
	"testing"
)

func Test_Day13Part1(t *testing.T) {
	input := `
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
	\------/ 
`
	basicTest(t, Day13Part1, strings.TrimSpace(input), "7,3")
}

func Test_Day13Part2(t *testing.T) {
	input := `
/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/
`
	basicTest(t, Day13Part2, strings.TrimSpace(input), "6,4")
}
