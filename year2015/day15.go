package year2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day15Part1(input string) (int, error) {
	props, err := parseCookieProperties(input)
	if err != nil {
		return 0, err
	}
	return determineHighestScoringCookieScore(props), nil
}

func Day15Part2(input string) (int, error) {
	props, err := parseCookieProperties(input)
	if err != nil {
		return 0, err
	}
	return determineHighestScoringCookieScoreWithExactCalories(props, 500), nil

}

func determineHighestScoringCookieScore(props []cookieProperties) int {
	return determineHighestScoringCookieScoreWithExactCalories(props, 0)
}

func determineHighestScoringCookieScoreWithExactCalories(props []cookieProperties, matchCalories int) int {
	amounts := make(map[cookieProperties]int)
	return findHighestScoringCookieScore(props, 100, amounts, matchCalories)
}

func findHighestScoringCookieScore(props []cookieProperties, remaining int, amounts map[cookieProperties]int, matchCalories int) int {
	if len(props) == 0 {
		totalAmount := 0
		for _, n := range amounts {
			totalAmount += n
		}
		if totalAmount == 100 {
			if matchCalories != 0 && determineTotalCalories(amounts) != matchCalories {
				return 0
			}
			return determineTotalScore(amounts)
		} else {
			return 0
		}
	}

	max := 0
	for i := 0; i <= remaining; i++ {
		amounts[props[0]] = i
		score := findHighestScoringCookieScore(props[1:], remaining-i, amounts, matchCalories)
		if score > max {
			max = score
		}
	}
	return max
}

func determineTotalScore(amounts map[cookieProperties]int) int {
	var totalCapacity, totalDurability, totalFlavor, totalTexture int
	for props, amount := range amounts {
		totalCapacity += props.capacity * amount
		totalDurability += props.durability * amount
		totalFlavor += props.flavor * amount
		totalTexture += props.texture * amount
	}
	return negativeToZero(totalCapacity) * negativeToZero(totalDurability) * negativeToZero(totalFlavor) * negativeToZero(totalTexture)
}

func determineTotalCalories(amounts map[cookieProperties]int) int {
	total := 0
	for props, amount := range amounts {
		total += props.calories * amount
	}
	return total
}

func negativeToZero(i int) int {
	if i < 0 {
		return 0
	}
	return i
}

type cookieProperties struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

var cookiePropertiesRegexp = regexp.MustCompile(`(\w+): capacity (-?\d), durability (-?\d), flavor (-?\d), texture (-?\d), calories (-?\d)`)

func parseCookieProperties(input string) ([]cookieProperties, error) {
	allMatches := cookiePropertiesRegexp.FindAllStringSubmatch(strings.TrimSpace(input), -1)
	if len(allMatches) == 0 {
		return nil, fmt.Errorf("invalid cookie properties")
	}
	out := make([]cookieProperties, len(allMatches))
	for i, matches := range allMatches {
		if capacity, err := strconv.Atoi(matches[2]); err != nil {
			return nil, fmt.Errorf("invalid capacity: %w", err)
		} else if durability, err := strconv.Atoi(matches[3]); err != nil {
			return nil, fmt.Errorf("invalid durability: %w", err)
		} else if flavor, err := strconv.Atoi(matches[4]); err != nil {
			return nil, fmt.Errorf("invalid flavor: %w", err)
		} else if texture, err := strconv.Atoi(matches[5]); err != nil {
			return nil, fmt.Errorf("invalid texture: %w", err)
		} else if calories, err := strconv.Atoi(matches[6]); err != nil {
			return nil, fmt.Errorf("invalid calories: %w", err)
		} else {
			out[i] = cookieProperties{name: matches[1], capacity: capacity, durability: durability, flavor: flavor, texture: texture, calories: calories}
		}
	}

	return out, nil
}
