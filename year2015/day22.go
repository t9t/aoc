package year2015

import (
	"fmt"
	"math"
	"strings"
)

func Day22Part1(input string) (int, error) {
	return determineMinimumManaToWin(input, false)
}

func Day22Part2(input string) (int, error) {
	return determineMinimumManaToWin(input, true)
}

func determineMinimumManaToWin(input string, hardMode bool) (int, error) {
	boss, err := parseBossInputWithoutArmor(input)
	if err != nil {
		return 0, err
	}

	min := &manaRegistration{mana: math.MaxInt32}
	tryNextWizardMove(rpgCharacter{hp: 50, mana: 500}, boss, hardMode, min)
	if min.mana == math.MaxInt32 {
		return 0, fmt.Errorf("player did not win")
	}
	return min.mana, nil
}

func parseBossInputWithoutArmor(input string) (rpgCharacter, error) {
	patched := strings.TrimSpace(input) + "\nArmor: 0"
	return parseBossInput(patched)
}

type wizardSpell int

const (
	spellMagicMissile wizardSpell = iota
	spellDrain
	spellShield
	spellPoison
	spellRecharge
)

var wizardSpells = []wizardSpell{spellMagicMissile, spellDrain, spellShield, spellPoison, spellRecharge}

const (
	manaCostMagicMissile = 53
	manaCostDrain        = 73
	manaCostShield       = 113
	manaCostPoison       = 173
	manaCostRecharge     = 229
)

type manaRegistration struct {
	mana int
}

func (r *manaRegistration) update(c rpgCharacter) {
	if c.totalManaSpent < r.mana {
		r.mana = c.totalManaSpent
	}
}

func tryNextWizardMove(inPlayer rpgCharacter, inBoss rpgCharacter, hardMode bool, minMana *manaRegistration) {
	if hardMode {
		inPlayer.hp -= 1
		if inPlayer.isDead() {
			return
		}
	}

	if inPlayer.mana < manaCostMagicMissile {
		return
	}

	inPlayer, inBoss = processAllEffects(inPlayer, inBoss)

	if inBoss.isDead() {
		// Boss died from a status effect (eg. poison)
		minMana.update(inPlayer)
		return
	}

	for _, spell := range wizardSpells {
		branchPlayer, branchBoss := inPlayer, inBoss

		var cast bool
		branchPlayer, branchBoss, cast = tryCastSpell(spell, branchPlayer, branchBoss)
		if !cast {
			continue
		}

		if branchPlayer.totalManaSpent >= minMana.mana {
			// No hope in this branch
			continue
		}

		if branchBoss.isDead() {
			// Boss died from the attack
			minMana.update(branchPlayer)
			continue
		}

		// Boss turn
		branchPlayer, branchBoss = processAllEffects(branchPlayer, branchBoss)
		if branchBoss.isDead() {
			// Boss died from a status effect (eg. poison)
			minMana.update(branchPlayer)
			continue
		}

		branchPlayer, branchBoss = bossAttack(branchPlayer, branchBoss)
		if branchPlayer.isDead() {
			// RIP
			continue
		}

		// Still alive, we need to go deeper
		tryNextWizardMove(branchPlayer, branchBoss, hardMode, minMana)
	}
}

func tryCastSpell(spell wizardSpell, player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	switch spell {
	case spellMagicMissile:
		return tryCastMagicMissile(player, boss)
	case spellDrain:
		return tryCastDrain(player, boss)
	case spellShield:
		return tryCastShield(player, boss)
	case spellPoison:
		return tryCastPoison(player, boss)
	case spellRecharge:
		return tryCastRecharge(player, boss)
	}
	return player, boss, false
}

func tryCastMagicMissile(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if player.mana < manaCostMagicMissile {
		return player, boss, false
	}
	player.mana -= manaCostMagicMissile
	player.totalManaSpent += manaCostMagicMissile
	boss.hp -= 4
	return player, boss, true
}

func tryCastDrain(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if player.mana < manaCostDrain {
		return player, boss, false
	}
	player.mana -= manaCostDrain
	player.totalManaSpent += manaCostDrain
	player.hp += 2
	boss.hp -= 2
	return player, boss, true
}

func tryCastShield(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if player.mana < manaCostShield || player.shieldTurns > 0 {
		return player, boss, false
	}
	player.mana -= manaCostShield
	player.totalManaSpent += manaCostShield
	player.shieldTurns = 6
	player.ac = 7
	return player, boss, true
}

func tryCastPoison(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if player.mana < manaCostPoison || boss.poisonTurns > 0 {
		return player, boss, false
	}
	player.mana -= manaCostPoison
	player.totalManaSpent += manaCostPoison
	boss.poisonTurns = 6
	return player, boss, true
}

func tryCastRecharge(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if player.mana < manaCostRecharge || player.rechargeTurns > 0 {
		return player, boss, false
	}
	player.mana -= manaCostRecharge
	player.totalManaSpent += manaCostRecharge
	player.rechargeTurns = 5
	return player, boss, true
}

func bossAttack(player, boss rpgCharacter) (rpgCharacter, rpgCharacter) {
	player.hp -= atLeastOne(boss.dmg - player.ac)
	return player, boss
}

func processEffects(c rpgCharacter) rpgCharacter {
	if c.poisonTurns > 0 {
		c.poisonTurns--
		c.hp -= 3
	}
	if c.shieldTurns > 0 {
		c.shieldTurns--
		if c.shieldTurns == 0 {
			c.ac = 0
		}
	}
	if c.rechargeTurns > 0 {
		c.rechargeTurns--
		c.mana += 101
	}
	return c
}

func processAllEffects(player, boss rpgCharacter) (rpgCharacter, rpgCharacter) {
	player, boss = processEffects(player), processEffects(boss)
	return player, boss
}
