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

const (
	manaCostMagicMissile = 53
)
const shieldAc = 7

var wizardSpells = []wizardSpell{
	&SpellMagicMissile{baseWizardSpell{manaCost: manaCostMagicMissile}},
	&SpellDrain{baseWizardSpell{manaCost: 73}},
	&SpellShield{baseWizardSpell{manaCost: 113}},
	&SpellPoison{baseWizardSpell{manaCost: 173}},
	&SpellRecharge{baseWizardSpell{manaCost: 229}},
}

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

		if !branchPlayer.canAfford(spell.ManaCost()) {
			continue
		}

		var wasCast bool
		branchPlayer, branchBoss, wasCast = spell.Cast(branchPlayer, branchBoss)
		if !wasCast {
			continue
		}

		if inPlayer.totalManaSpent >= minMana.mana {
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

type wizardSpell interface {
	ManaCost() int
	Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool)
}

type baseWizardSpell struct {
	manaCost int
}

func (s *baseWizardSpell) payMana(player rpgCharacter) rpgCharacter {
	if !player.canAfford(s.manaCost) {
		panic(fmt.Sprintf("Player %+v cannot afford spell %+v", player, s))
	}
	player.mana -= s.manaCost
	player.totalManaSpent += s.manaCost
	return player
}

func (s *baseWizardSpell) ManaCost() int {
	return s.manaCost
}

type SpellMagicMissile struct{ baseWizardSpell }

// Magic Missile costs 53 mana. It instantly does 4 damage.
func (s *SpellMagicMissile) Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	player = s.payMana(player)
	boss.hp -= 4
	return player, boss, true
}

type SpellDrain struct{ baseWizardSpell }

// Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
func (s *SpellDrain) Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	player = s.payMana(player)
	boss.hp -= 2
	player.hp += 2
	return player, boss, true
}

type SpellShield struct{ baseWizardSpell }

// Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
func (s *SpellShield) Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if player.shieldTurns > 0 {
		return player, boss, false
	}
	player = s.payMana(player)
	player.shieldTurns = 6
	player.ac += shieldAc
	return player, boss, true
}

type SpellPoison struct{ baseWizardSpell }

// Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
func (s *SpellPoison) Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if boss.poisonTurns > 0 {
		return player, boss, false
	}
	player = s.payMana(player)
	boss.poisonTurns = 6
	return player, boss, true
}

type SpellRecharge struct{ baseWizardSpell }

// Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.
func (s *SpellRecharge) Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if player.rechargeTurns > 0 {
		return player, boss, false
	}
	player = s.payMana(player)
	player.rechargeTurns = 5
	return player, boss, true
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
