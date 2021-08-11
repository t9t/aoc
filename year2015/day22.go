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

const (
	manaCostMagicMissile = 53
	manaCostDrain        = 73
	manaCostShield       = 113
	manaCostPoison       = 173
	manaCostRecharge     = 229
)
const shieldAc = 7

var wizardSpells = []wizardSpell{spellMagicMissile, spellDrain, spellShield, spellPoison, spellRecharge}

type manaRegistration struct {
	mana int
}

func (r *manaRegistration) update(c rpgCharacter) {
	if c.manaSpent < r.mana {
		r.mana = c.manaSpent
	}
}

func tryNextWizardMove(inPlayer rpgCharacter, inBoss rpgCharacter, hardMode bool, minMana *manaRegistration) {
	if hardMode {
		inPlayer.hp -= 1
		if inPlayer.isDead() {
			return
		}
	}

	inPlayer, inBoss = processEffects(inPlayer), processEffects(inBoss)

	if inBoss.isDead() {
		// Boss died from a status effect (eg. poison)
		minMana.update(inPlayer)
		return
	}

	// TODO: implement better way to termine min mana spell
	if !inPlayer.canAfford(manaCostMagicMissile) {
		return
	}

	for _, spell := range wizardSpells {
		branchPlayer, branchBoss := inPlayer, inBoss

		switch spell {
		case spellMagicMissile:
			if !branchPlayer.canAfford(manaCostMagicMissile) {
				continue
			}
			branchPlayer, branchBoss = castMagicMissile(branchPlayer, branchBoss)
		case spellDrain:
			if !branchPlayer.canAfford(manaCostDrain) {
				continue
			}
			branchPlayer, branchBoss = castDrain(branchPlayer, branchBoss)
		case spellShield:
			if !branchPlayer.canAfford(manaCostShield) || branchPlayer.shieldTurns > 0 {
				continue
			}
			branchPlayer = applyShield(branchPlayer)
		case spellPoison:
			if !branchPlayer.canAfford(manaCostPoison) || branchBoss.poisonTurns > 0 {
				continue
			}
			branchPlayer, branchBoss = applyPoison(branchPlayer, branchBoss)
		case spellRecharge:
			if !branchPlayer.canAfford(manaCostRecharge) || branchPlayer.rechargeTurns > 0 {
				continue
			}
			branchPlayer = applyRecharge(branchPlayer)
		}

		if inPlayer.manaSpent >= minMana.mana {
			// No hope in this branch
			continue
		}

		if branchBoss.isDead() {
			// Boss died from the attack
			minMana.update(branchPlayer)
			continue
		}

		// Boss turn
		branchPlayer, branchBoss = processEffects(branchPlayer), processEffects(branchBoss)
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

// Magic Missile costs 53 mana. It instantly does 4 damage.
func castMagicMissile(src, dest rpgCharacter) (rpgCharacter, rpgCharacter) {
	if src.mana < manaCostMagicMissile {
		panic(fmt.Sprintf("Player %+v does not have >= %d mana for Magic Missile", src, manaCostMagicMissile))
	}
	src.mana -= manaCostMagicMissile
	src.manaSpent += manaCostMagicMissile
	dest.hp -= 4
	return src, dest
}

// Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
func castDrain(player, boss rpgCharacter) (rpgCharacter, rpgCharacter) {
	if player.mana < manaCostDrain {
		panic(fmt.Sprintf("Player %+v does not have >= %d mana for Drain", player, manaCostDrain))
	}
	player.mana -= manaCostDrain
	player.manaSpent += manaCostDrain
	boss.hp -= 2
	player.hp += 2
	return player, boss
}

// Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
func applyShield(c rpgCharacter) rpgCharacter {
	if c.mana < manaCostShield {
		panic(fmt.Sprintf("Player %+v does not have >= %d mana for Shield", c, manaCostShield))
	}
	if c.shieldTurns > 0 {
		panic(fmt.Sprintf("Shield already active on %+v", c))
	}
	c.mana -= manaCostShield
	c.manaSpent += manaCostShield
	c.shieldTurns = 6
	c.ac += shieldAc
	return c
}

// Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
func applyPoison(src, dest rpgCharacter) (rpgCharacter, rpgCharacter) {
	if src.mana < manaCostPoison {
		panic(fmt.Sprintf("Player %+v does not have >= %d mana for Poison", src, manaCostPoison))
	}
	if dest.poisonTurns > 0 {
		panic(fmt.Sprintf("Poison already active on %+v", dest))
	}
	src.mana -= manaCostPoison
	src.manaSpent += manaCostPoison
	dest.poisonTurns = 6
	return src, dest
}

// Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.
func applyRecharge(c rpgCharacter) rpgCharacter {
	if c.mana < manaCostRecharge {
		panic(fmt.Sprintf("Player %+v does not have >= %d mana for Recharge", c, manaCostRecharge))
	}
	if c.rechargeTurns > 0 {
		panic(fmt.Sprintf("Recharge already active on %+v", c))
	}
	c.mana -= manaCostRecharge
	c.manaSpent += manaCostRecharge
	c.rechargeTurns = 5
	return c
}

func bossAttack(player, boss rpgCharacter) (rpgCharacter, rpgCharacter) {
	player.hp -= atLeastOne(boss.dmg - player.ac)
	return player, boss
}

func processEffects(c rpgCharacter) rpgCharacter {
	if c.shieldTurns > 0 {
		c.shieldTurns--
		if c.shieldTurns == 0 {
			c.ac -= shieldAc
		}
	}
	if c.poisonTurns > 0 {
		c.hp -= 3
		c.poisonTurns--
	}
	if c.rechargeTurns > 0 {
		c.mana += 101
		c.rechargeTurns--
	}
	return c
}
