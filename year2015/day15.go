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
	return determineHighestScoringCookieScore4(props), nil
}

func Day15Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func determineHighestScoringCookieScore2(props []cookieProperties) int {
	if len(props) != 2 {
		panic(fmt.Sprintf("%d != %d", len(props), 2))
	}
	max := 0
	for i1 := 0; i1 <= 100; i1++ {
		for i2 := 0; i2 <= 100; i2++ {
			if i1+i2 != 100 {
				continue
			}
			amounts := map[cookieProperties]int{
				props[0]: i1,
				props[1]: i2,
			}
			score := determineTotalScore(amounts)
			if score > max {
				max = score
			}
		}
	}
	return max
}

func determineHighestScoringCookieScore4(props []cookieProperties) int {
	if len(props) != 4 {
		panic(fmt.Sprintf("%d != %d", len(props), 4))
	}
	max := 0
	for i1 := 0; i1 <= 100; i1++ {
		for i2 := 0; i2 <= 100; i2++ {
			for i3 := 0; i3 <= 100; i3++ {
				for i4 := 0; i4 <= 100; i4++ {
					if i1+i2+i3+i4 != 100 {
						continue
					}
					amounts := map[cookieProperties]int{
						props[0]: i1,
						props[1]: i2,
						props[2]: i3,
						props[3]: i4,
					}
					score := determineTotalScore(amounts)
					if score > max {
						max = score
					}
				}
			}
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
