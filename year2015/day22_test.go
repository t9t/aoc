package year2015

import (
	"reflect"
	"testing"
)

func Test_parseBossInputWithoutArmor(t *testing.T) {
	input := "Hit Points: 1337\nDamage: 5521"
	want := rpgCharacter{hp: 1337, dmg: 5521, ac: 0}
	got, err := parseBossInputWithoutArmor(input)
	if err != nil {
		t.Errorf("parseBossInputWithoutArmor() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseBossInputWithoutArmor() = %v, want %v", got, want)
	}
}

func Test_determineMinimumManaToWin(t *testing.T) {
	type args struct {
		player   rpgCharacter
		boss     rpgCharacter
		hardMode bool
	}
	player := rpgCharacter{hp: 10, mana: 250}
	boss := rpgCharacter{hp: 13, dmg: 8}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"normal mode mode", args{player: player, boss: boss, hardMode: false}, 226, false},
		{"hard mode", args{player: rpgCharacter{hp: 15, mana: 250}, boss: rpgCharacter{hp: 20, dmg: 5}, hardMode: true}, 621, false},
		{"loser", args{player: player, boss: boss, hardMode: true}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := determineMinimumManaToWin(tt.args.player, tt.args.boss, tt.args.hardMode)
			if (err != nil) != tt.wantErr {
				t.Errorf("determineMinimumManaToWin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("determineMinimumManaToWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
