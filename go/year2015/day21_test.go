package year2015

import (
	"reflect"
	"testing"
)

func Test_playRound(t *testing.T) {
	// For example, suppose you have 8 hit points, 5 damage, and 5 armor, and that the boss has 12 hit points, 7 damage, and 2 armor:
	player := rpgCharacter{hp: 8, dmg: 5, ac: 5}
	boss := rpgCharacter{hp: 12, dmg: 7, ac: 2}

	weakSauce := rpgCharacter{hp: 10, dmg: 0, ac: 10}
	megaArmor := rpgCharacter{hp: 10, dmg: 10, ac: 1000}

	withHp := func(c rpgCharacter, hp int) rpgCharacter {
		c.hp = hp
		return c
	}
	tests := []struct {
		name      string
		playerIn  rpgCharacter
		bossIn    rpgCharacter
		playerOut rpgCharacter
		bossOut   rpgCharacter
	}{
		// The player deals 5-2 = 3 damage; the boss goes down to 9 hit points.
		// The boss deals 7-5 = 2 damage; the player goes down to 6 hit points.
		{"round 1", player, boss, withHp(player, 6), withHp(boss, 9)},
		// The player deals 5-2 = 3 damage; the boss goes down to 6 hit points.
		// The boss deals 7-5 = 2 damage; the player goes down to 4 hit points.
		{"round 2", withHp(player, 6), withHp(boss, 9), withHp(player, 4), withHp(boss, 6)},
		// The player deals 5-2 = 3 damage; the boss goes down to 3 hit points.
		// The boss deals 7-5 = 2 damage; the player goes down to 2 hit points.
		{"round 3", withHp(player, 4), withHp(boss, 6), withHp(player, 2), withHp(boss, 3)},
		// The player deals 5-2 = 3 damage; the boss goes down to 0 hit points.
		{"round 4", withHp(player, 2), withHp(boss, 3), withHp(player, 2), withHp(boss, 0)},

		// Each attack reduces the opponent's hit points by at least 1
		{"no dmg still 1 damage", weakSauce, weakSauce, withHp(weakSauce, 9), withHp(weakSauce, 9)},
		{"mega armor still 1 damage", megaArmor, megaArmor, withHp(megaArmor, 9), withHp(megaArmor, 9)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			playerOut, bossOut := playRound(tt.playerIn, tt.bossIn)
			if !reflect.DeepEqual(playerOut, tt.playerOut) {
				t.Errorf("playRound() got = %+v, want %+v", playerOut, tt.playerOut)
			}
			if !reflect.DeepEqual(bossOut, tt.bossOut) {
				t.Errorf("playRound() got1 = %+v, want %+v", bossOut, tt.bossOut)
			}
		})
	}
}

func Test_playGame_playerWin(t *testing.T) {
	// For example, suppose you have 8 hit points, 5 damage, and 5 armor, and that the boss has 12 hit points, 7 damage, and 2 armor:
	player := rpgCharacter{hp: 8, dmg: 5, ac: 5}
	boss := rpgCharacter{hp: 12, dmg: 7, ac: 2}

	playerWant := rpgCharacter{hp: 2, dmg: 5, ac: 5}
	bossWant := rpgCharacter{hp: 0, dmg: 7, ac: 2}

	wantedTurns := 4

	playerOut, bossOut, turns := playGame(player, boss)
	if !reflect.DeepEqual(playerOut, playerWant) {
		t.Errorf("playGame() playerOut = %+v, want %+v", playerOut, playerWant)
	}
	if !reflect.DeepEqual(bossOut, bossWant) {
		t.Errorf("playGame() bossOut = %+v, want %+v", bossOut, bossWant)
	}
	if turns != wantedTurns {
		t.Errorf("playGame() turns = %+v, want %+v", turns, wantedTurns)
	}
}

func Test_playGame_bossWin(t *testing.T) {
	player := rpgCharacter{hp: 8, dmg: 1, ac: 1}
	boss := rpgCharacter{hp: 12, dmg: 5, ac: 2}

	playerWant := rpgCharacter{hp: 0, dmg: 1, ac: 1}
	bossWant := rpgCharacter{hp: 10, dmg: 5, ac: 2}

	wantedTurns := 2

	playerOut, bossOut, turns := playGame(player, boss)
	if !reflect.DeepEqual(playerOut, playerWant) {
		t.Errorf("playGame() playerOut = %+v, want %+v", playerOut, playerWant)
	}
	if !reflect.DeepEqual(bossOut, bossWant) {
		t.Errorf("playGame() bossOut = %+v, want %+v", bossOut, bossWant)
	}
	if turns != wantedTurns {
		t.Errorf("playGame() turns = %+v, want %+v", turns, wantedTurns)
	}
}

func Test_parseBossInput(t *testing.T) {
	input := "Hit Points: 123\nDamage: 7\nArmor: 3\n"
	want := rpgCharacter{hp: 123, dmg: 7, ac: 3}
	gotBoss, err := parseBossInput(input)
	if err != nil {
		t.Errorf("parseBossInput() error = %v", err)
		return
	}
	if !reflect.DeepEqual(gotBoss, want) {
		t.Errorf("parseBossInput() = %v, want %v", gotBoss, want)
	}
}
