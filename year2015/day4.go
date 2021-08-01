package year2015

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Day4Part1(input string) (int, error) {
	return findSumStartingWith(input, "00000")
}

func Day4Part2(input string) (int, error) {
	return findSumStartingWith(input, "000000")
}

func findSumStartingWith(input, prefix string) (int, error) {
	const max = 10_000_000
	for i := 0; i < max; i++ {
		sum := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))

		if strings.HasPrefix(sum, prefix) {
			return i, nil
		}
	}

	return 0, fmt.Errorf("no hash starting with %q found after %d iterations", prefix, max)
}
