package year2015

import (
	"reflect"
	"testing"
)

func Test_parseCookieProperties(t *testing.T) {
	input := `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`
	want := []cookieProperties{
		{name: "Butterscotch", capacity: -1, durability: -2, flavor: 6, texture: 3, calories: 8},
		{name: "Cinnamon", capacity: 2, durability: 3, flavor: -2, texture: -1, calories: 3},
	}
	got, err := parseCookieProperties(input)
	if err != nil {
		t.Errorf("parseCookieProperties() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseCookieProperties() = %v, want %v", got, want)
	}
}

func Test_determineTotalScore(t *testing.T) {
	tests := []struct {
		name    string
		amounts map[cookieProperties]int
		want    int
	}{
		{"positive", map[cookieProperties]int{
			{name: "Butterscotch", capacity: -1, durability: -2, flavor: 6, texture: 3, calories: 8}: 44,
			{name: "Cinnamon", capacity: 2, durability: 3, flavor: -2, texture: -1, calories: 3}:     56,
		}, 62842880},
		{"negative becomes zero", map[cookieProperties]int{
			// Durability: (-5*90)+(3*10)=-450+30=-420 blaze it
			{name: "x", capacity: -1, durability: -5, flavor: 6, texture: 3, calories: 8}: 90,
			{name: "y", capacity: 2, durability: 3, flavor: -2, texture: -1, calories: 3}: 10,
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineTotalScore(tt.amounts); got != tt.want {
				t.Errorf("determineTotalScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_determineHighestScoringCookieScore2(t *testing.T) {
	props := []cookieProperties{{name: "Butterscotch", capacity: -1, durability: -2, flavor: 6, texture: 3, calories: 8},
		{name: "Cinnamon", capacity: 2, durability: 3, flavor: -2, texture: -1, calories: 3}}
	want := 62842880
	if got := determineHighestScoringCookieScore2(props); got != want {
		t.Errorf("determineHighestScoringCookieScore() = %v, want %v", got, want)
	}
}
