package year2015

import (
	"reflect"
	"testing"
)

func Test_parseReindeerDescriptions(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []reindeerDescription
	}{
		{"comet", `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.`, []reindeerDescription{{flySpeed: 14, flyTime: 10, restTime: 127}}},
		{"dancer", `Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`, []reindeerDescription{{flySpeed: 16, flyTime: 11, restTime: 162}}},
		{"both", "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.\nDancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
			[]reindeerDescription{{flySpeed: 14, flyTime: 10, restTime: 127}, {flySpeed: 16, flyTime: 11, restTime: 162}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseReindeerDescriptions(tt.input)
			if err != nil {
				t.Errorf("parseReindeerDescriptions() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseReindeerDescriptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reindeerDescription_distanceTraveledAfter(t *testing.T) {
	tests := []struct {
		name    string
		r       reindeerDescription
		seconds int
		want    int
	}{
		{"comet 1s", reindeerDescription{flySpeed: 14, flyTime: 10, restTime: 127}, 1, 14},
		{"comet 10s", reindeerDescription{flySpeed: 14, flyTime: 10, restTime: 127}, 10, 140},
		{"comet 50s", reindeerDescription{flySpeed: 14, flyTime: 10, restTime: 127}, 50, 140},
		{"comet 137s", reindeerDescription{flySpeed: 14, flyTime: 10, restTime: 127}, 137, 140},
		{"comet 138s", reindeerDescription{flySpeed: 14, flyTime: 10, restTime: 127}, 138, 154},
		{"comet 1000s", reindeerDescription{flySpeed: 14, flyTime: 10, restTime: 127}, 1000, 1120},

		{"dancer 1s", reindeerDescription{flySpeed: 16, flyTime: 11, restTime: 162}, 1, 16},
		{"dancer 10s", reindeerDescription{flySpeed: 16, flyTime: 11, restTime: 162}, 10, 160},
		{"dancer 10s", reindeerDescription{flySpeed: 16, flyTime: 11, restTime: 162}, 11, 176},
		{"dancer 50s", reindeerDescription{flySpeed: 16, flyTime: 11, restTime: 162}, 50, 176},
		{"dancer 137s", reindeerDescription{flySpeed: 16, flyTime: 11, restTime: 162}, 173, 176},
		{"dancer 138s", reindeerDescription{flySpeed: 16, flyTime: 11, restTime: 162}, 174, 192},
		{"dancer 1000s", reindeerDescription{flySpeed: 16, flyTime: 11, restTime: 162}, 1000, 1056},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.distanceTraveledAfter(tt.seconds); got != tt.want {
				t.Errorf("reindeerDescription.distanceTraveledAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxDistanceTraveled(t *testing.T) {
	input := `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`
	want := 1120
	got, err := findMaxDistanceTraveled(input, 1000)
	if err != nil {
		t.Errorf("findMaxDistanceTraveled() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("findMaxDistanceTraveled() = %v, want %v", got, want)
	}
}
