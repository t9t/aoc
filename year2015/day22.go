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

var wizardSpells = []wizardSpell{
	&spellMagicMissile{baseWizardSpell{manaCost: 53}},
	&spellDrain{baseWizardSpell{manaCost: 73}},
	&spellShield{baseWizardSpell{manaCost: 113}},
	&spellPoison{baseWizardSpell{manaCost: 173}},
	&spellRecharge{baseWizardSpell{manaCost: 229}},
}

var statusEffects = []statusEffect{
	&spellShield{},
	&spellPoison{},
	&spellRecharge{},
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

	inPlayer, inBoss = processAllEffects(inPlayer, inBoss)

	if inBoss.isDead() {
		// Boss died from a status effect (eg. poison)
		minMana.update(inPlayer)
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

type wizardSpell interface {
	ManaCost() int
	Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool)
}

type statusEffect interface {
	Process(rpgCharacter) rpgCharacter
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

type spellMagicMissile struct{ baseWizardSpell }

// Magic Missile costs 53 mana. It instantly does 4 damage.
func (s *spellMagicMissile) Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	player = s.payMana(player)
	boss.hp -= 4
	return player, boss, true
}

type spellDrain struct{ baseWizardSpell }

// Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
func (s *spellDrain) Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	player = s.payMana(player)
	boss.hp -= 2
	player.hp += 2
	return player, boss, true
}

type spellShield struct{ baseWizardSpell }

// Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
func (s *spellShield) Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if player.shieldTurns > 0 {
		return player, boss, false
	}
	player = s.payMana(player)
	player.shieldTurns = 6
	player.ac += 7
	return player, boss, true
}

func (s *spellShield) Process(c rpgCharacter) rpgCharacter {
	if c.shieldTurns > 0 {
		c.shieldTurns--
		if c.shieldTurns == 0 {
			c.ac -= 7
		}
	}
	return c
}

type spellPoison struct{ baseWizardSpell }

// Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
func (s *spellPoison) Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if boss.poisonTurns > 0 {
		return player, boss, false
	}
	player = s.payMana(player)
	boss.poisonTurns = 6
	return player, boss, true
}

func (s *spellPoison) Process(c rpgCharacter) rpgCharacter {
	if c.poisonTurns > 0 {
		c.poisonTurns--
		c.hp -= 3
	}
	return c
}

type spellRecharge struct{ baseWizardSpell }

// Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.
func (s *spellRecharge) Cast(player, boss rpgCharacter) (rpgCharacter, rpgCharacter, bool) {
	if player.rechargeTurns > 0 {
		return player, boss, false
	}
	player = s.payMana(player)
	player.rechargeTurns = 5
	return player, boss, true
}

func (s *spellRecharge) Process(c rpgCharacter) rpgCharacter {
	if c.rechargeTurns > 0 {
		c.rechargeTurns--
		c.mana += 101
	}
	return c
}

func bossAttack(player, boss rpgCharacter) (rpgCharacter, rpgCharacter) {
	player.hp -= atLeastOne(boss.dmg - player.ac)
	return player, boss
}

func processEffects(c rpgCharacter) rpgCharacter {
	for _, effect := range statusEffects {
		c = effect.Process(c)
	}
	return c
}

func processAllEffects(player, boss rpgCharacter) (rpgCharacter, rpgCharacter) {
	player, boss = processEffects(player), processEffects(boss)
	return player, boss
}
