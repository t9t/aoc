package year2015

import (
	"crypto/md5"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func Day4Part1(input string) (int, error) {
	return findSumStartingWithParallel(input, "00000")
}

func Day4Part2(input string) (int, error) {
	return findSumStartingWithParallel(input, "000000")
}

//nolint:deadcode,unused
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

func findSumStartingWithParallel(input, prefix string) (int, error) {
	workerCount := runtime.NumCPU()
	batchSize := 30_000
	resultChan := make(chan int, workerCount)
	defer func() {
		close(resultChan)
	}()

	i := 0
	for {
		for w := 0; w < workerCount; w++ {
			go func(start, end int) {
				if result, found := findSumStartingWithBetweenStartAndEnd(input, prefix, start, end); found {
					resultChan <- result
				} else {
					resultChan <- -1
				}
			}(i, i+batchSize)
			i += batchSize
		}

		lowestResult := -1
		for w := 0; w < workerCount; w++ {
			if r := <-resultChan; r != -1 {
				if lowestResult == -1 || r < lowestResult {
					lowestResult = r
				}
			}
		}

		if lowestResult != -1 {
			return lowestResult, nil
		}
	}
}

func findSumStartingWithBetweenStartAndEnd(input, prefix string, start, end int) (int, bool) {
	for i := start; i < end; i++ {
		sum := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))

		if strings.HasPrefix(sum, prefix) {
			return i, true
		}
	}

	return 0, false
}
