package year2018

import (
	"fmt"
	"regexp"
	"strconv"
)

func init() {
	mustRegisterPair(9, Day9Part1, Day9Part2)
}

func Day9Part1(input string) (string, error) {
	type marble struct{ cw, ccw int }

	matches := regexp.MustCompile(`(\d+) players; last marble is worth (\d+) points`).FindStringSubmatch(input)

	var playerCount, highest int
	var err error
	if playerCount, err = strconv.Atoi(matches[1]); err != nil {
		return "", err
	} else if highest, err = strconv.Atoi(matches[2]); err != nil {
		return "", err
	}

	marbles := map[int]*marble{0: {cw: 0, ccw: 0}}
	currentNumber, playerNumber := 0, 1
	scores := make([]int, playerCount+1)

	for marbleNum := 1; marbleNum <= highest; marbleNum++ {
		if marbleNum%23 == 0 {
			scores[playerNumber] += marbleNum

			next := marbles[currentNumber]
			nextNum := currentNumber
			for i := 1; i <= 7; i++ {
				nextNum = next.ccw
				next = marbles[nextNum]
			}
			delete(marbles, nextNum)
			marbles[next.ccw].cw = next.cw
			marbles[next.cw].ccw = next.ccw
			scores[playerNumber] += nextNum
			currentNumber = next.cw
		} else {
			currentMarble := marbles[currentNumber]
			clockwiseMarble := marbles[currentMarble.cw]

			newMarble := &marble{cw: clockwiseMarble.cw, ccw: currentMarble.cw}
			clockwiseMarble.cw = marbleNum
			marbles[newMarble.cw].ccw = marbleNum
			marbles[marbleNum] = newMarble

			currentNumber = marbleNum
		}

		playerNumber++
		if playerNumber > playerCount {
			playerNumber = 1
		}
	}

	highScore := 0
	for _, score := range scores {
		if score > highScore {
			highScore = score
		}
	}

	return strconv.Itoa(highScore), nil
}

func Day9Part2(input string) (string, error) {
	return "", fmt.Errorf("Day 9 part 2 not implemented")
}
