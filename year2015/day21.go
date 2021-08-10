package year2015

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Day21Part1(input string) (int, error) {
	boss, err := parseBossInput(input)
	if err != nil {
		return 0, err
	}

	weapons := []rpgItem{
		{g: 0, dmg: 0},  // Unarmed
		{g: 8, dmg: 4},  // Dagger
		{g: 10, dmg: 5}, // Shortsword
		{g: 25, dmg: 6}, // Warhammer
		{g: 40, dmg: 7}, // Longsword
		{g: 74, dmg: 8}, // Greataxe
	}
	armor := []rpgItem{
		{g: 0, ac: 0},   // Naked
		{g: 13, ac: 1},  // Leather
		{g: 31, ac: 2},  // Chainmail
		{g: 53, ac: 3},  // Splintmail
		{g: 75, ac: 4},  // Bandedmail
		{g: 102, ac: 5}, // Platemail
	}
	rings := []rpgItem{
		{g: 0, dmg: 0, ac: 0},   // Poor fellow!
		{g: 0, dmg: 0, ac: 0},   // Poor fellow!!
		{g: 25, dmg: 1, ac: 0},  // Damage +1
		{g: 50, dmg: 2, ac: 0},  // Damage +2
		{g: 100, dmg: 3, ac: 0}, // Damage +3
		{g: 20, dmg: 0, ac: 1},  // Defense +1
		{g: 40, dmg: 0, ac: 2},  // Defense +2
		{g: 80, dmg: 0, ac: 3},  // Defense +3
	}
	minCost := math.MaxInt32
	for _, weapon := range weapons {
		for _, armorItem := range armor {
			for i, leftRing := range rings {
				for j := i; j < len(rings); j++ {
					rightRing := rings[j]
					totalDmg := weapon.dmg + armorItem.dmg + leftRing.dmg + rightRing.dmg
					totalAc := weapon.ac + armorItem.ac + leftRing.ac + rightRing.ac
					inPlayer := rpgCharacter{hp: 100, dmg: totalDmg, ac: totalAc}
					_, outBoss, _ := playGame(inPlayer, boss)
					if playerWon := outBoss.hp <= 0; playerWon {
						if totalG := weapon.g + armorItem.g + leftRing.g + rightRing.g; totalG < minCost {
							minCost = totalG
						}
					}
				}
			}
		}
	}
	return minCost, nil
}

func Day21Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func parseBossInput(input string) (boss rpgCharacter, err error) {
	re := regexp.MustCompile(`Hit Points: (\d+)\nDamage: (\d+)\nArmor: (\d+)`)
	matches := re.FindStringSubmatch(strings.TrimSpace(input))
	if len(matches) != 4 {
		return boss, fmt.Errorf("invalid boss input")
	}
	if boss.hp, err = strconv.Atoi(matches[1]); err != nil {
		return boss, fmt.Errorf("invalid boss Hit Points in input: %w", err)
	} else if boss.dmg, err = strconv.Atoi(matches[2]); err != nil {
		return boss, fmt.Errorf("invalid boss Damage in input: %w", err)
	} else if boss.ac, err = strconv.Atoi(matches[3]); err != nil {
		return boss, fmt.Errorf("invalid boss Armor in input: %w", err)
	}
	return boss, nil
}

type rpgCharacter struct {
	hp, dmg, ac int
}

type rpgItem struct {
	g, dmg, ac int
}

func playRound(player, boss rpgCharacter) (rpgCharacter, rpgCharacter) {
	boss.hp -= atLeastOne(player.dmg - boss.ac)
	if boss.hp > 0 {
		player.hp -= atLeastOne(boss.dmg - player.ac)
	}
	return player, boss
}

func atLeastOne(n int) int {
	if n < 0 {
		return 1
	}
	return n
}

func playGame(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, int) {
	turns := 0
	for boss.hp > 0 && player.hp > 0 {
		turns++
		player, boss = playRound(player, boss)
	}
	return player, boss, turns
}
