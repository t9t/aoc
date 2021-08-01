package year2015

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Day4Part1(input string) (int, error) {
	const max = 10_000_000
	for i := 0; i < max; i++ {
		sum := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))

		if strings.HasPrefix(sum, "00000") {
			return i, nil
		}
	}

	return 0, fmt.Errorf("no answer found after %d iterations", max)
}
