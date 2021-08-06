package year2015

import (
	"reflect"
	"testing"
)

func Test_permutate(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{"one", []string{"a"}, [][]string{{"a"}}},
		{"two", []string{"a", "b"}, [][]string{{"a", "b"}, {"b", "a"}}},
		{"three", []string{"a", "b", "c"}, [][]string{{"a", "b", "c"}, {"a", "c", "b"}, {"b", "a", "c"}, {"b", "c", "a"}, {"c", "a", "b"}, {"c", "b", "a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutate(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copySkipping(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		skip  int
		want  []string
	}{
		{"three-skip-0", []string{"a", "b", "c"}, 0, []string{"b", "c"}},
		{"three-skip-1", []string{"a", "b", "c"}, 1, []string{"a", "c"}},
		{"three-skip-2", []string{"a", "b", "c"}, 2, []string{"a", "b"}},
		{"four-skip-2", []string{"a", "b", "c", "d"}, 2, []string{"a", "b", "d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copySkipping(tt.input, tt.skip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copySkipping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prepend(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		s     string
		want  []string
	}{
		{"empty slice", []string{}, "x", []string{"x"}},
		{"not empty slice", []string{"a", "b", "c"}, "x", []string{"x", "a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepend(tt.slice, tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseHappinessSpecs(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    happinessChangeMap
		wantErr bool
	}{
		{"invalid", "bla", nil, true},
		{"gain", "Alice would gain 54 happiness units by sitting next to Bob.", happinessChangeMap{"Alice": {"Bob": 54}}, false},
		{"lose", "Alice would lose 2 happiness units by sitting next to David.", happinessChangeMap{"Alice": {"David": -2}}, false},
		{"both", "Alice would gain 54 happiness units by sitting next to Bob.\nAlice would lose 2 happiness units by sitting next to David.",
			happinessChangeMap{"Alice": {"Bob": 54, "David": -2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseHappinessSpecs(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseHappinessSpecs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseHappinessSpecs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay13Part1(t *testing.T) {
	input := `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`

	want := 330

	got, err := Day13Part1(input)
	if err != nil {
		t.Errorf("Day13Part1() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("Day13Part1() = %v, want %v", got, want)
	}
}

func TestDay13Part2(t *testing.T) {
	input := `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.
Alice would gain 0 happiness units by sitting next to You.
Bob would gain 0 happiness units by sitting next to You.
Carol would gain 0 happiness units by sitting next to You.
David would gain 0 happiness units by sitting next to You.`

	want := 286

	got, err := Day13Part2(input)
	if err != nil {
		t.Errorf("Day13Part1() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("Day13Part1() = %v, want %v", got, want)
	}
}

func Test_calculateTotalHappinessChange(t *testing.T) {
	three := happinessChangeMap{
		"a": {"b": 10, "c": -5},
		"b": {"a": -3, "c": 8},
		"c": {"a": 7, "b": 1},
	}
	four := happinessChangeMap{
		"a": {"b": 10, "c": -5, "d": 7},
		"b": {"a": -3, "c": 8, "d": -4},
		"c": {"a": 7, "b": 1, "d": -2},
		"d": {"a": 7, "b": 1, "c": 9},
	}
	tests := []struct {
		name     string
		m        happinessChangeMap
		seatings []string
		want     int
	}{
		{"abc", three, []string{"a", "b", "c"}, 18},
		// d= +7 +9; c= -2 +1; b= +8 -3; a= +10 +7
		{"dcba", four, []string{"d", "c", "b", "a"}, 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTotalHappinessChange(tt.m, tt.seatings); got != tt.want {
				t.Errorf("calculateTotalHappinessChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
