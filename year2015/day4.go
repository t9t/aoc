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
	const max = 10_000_000
	workerCount := runtime.NumCPU()
	batchSize := 10_000

	resultChan := make(chan int, workerCount)
	batchChans := make([]chan []int, workerCount)
	for i := 0; i < workerCount; i++ {
		batchChans[i] = make(chan []int)
		startWorker(input, prefix, batchChans[i], resultChan)
	}
	defer func() {
		for _, c := range batchChans {
			close(c)
		}
	}()

	for i := 0; i < max; i++ {
		for b := 0; b < workerCount; b++ {
			batch := make([]int, batchSize)
			for j := 0; j < len(batch); j++ {
				batch[j] = i
				i++
			}
			batchChans[b] <- batch
		}
		lowestResult := -1
		for j := 0; j < workerCount; j++ {
			r := <-resultChan
			if r != -1 {
				if lowestResult == -1 || r < lowestResult {
					lowestResult = r
				}
			}
		}

		if lowestResult != -1 {
			return lowestResult, nil
		}
	}
	return 0, fmt.Errorf("no hash starting with %q found after %d iterations", prefix, max)
}

func startWorker(input, prefix string, batchChan chan []int, resultChan chan int) {
	go func() {
		for {
			batch, ok := <-batchChan
			if !ok {
				break
			}
			if result, found := findSumStartingWithInBatch(input, prefix, batch); found {
				resultChan <- result
			} else {
				resultChan <- -1
			}
		}
	}()
}

func findSumStartingWithInBatch(input, prefix string, batch []int) (int, bool) {
	for _, i := range batch {
		sum := fmt.Sprintf("%x", md5.Sum([]byte(input+strconv.Itoa(i))))

		if strings.HasPrefix(sum, prefix) {
			return i, true
		}
	}

	return 0, false
}
