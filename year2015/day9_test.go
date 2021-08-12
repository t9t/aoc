package year2015

import (
	"reflect"
	"testing"
)

func TestDay9Part1(t *testing.T) {
	input := `Dublin to Belfast = 141
London to Dublin = 464
London to Belfast = 518
`
	want := 605

	got, err := Day9Part1(input)
	if err != nil {
		t.Errorf("Day9Part1() error = %v", err)
		return
	}

	if got != want {
		t.Errorf("Day9Part1() = %v, want %v", got, want)
	}
}

func TestDay9Part1_broken(t *testing.T) {
	// TODO: fix code and enable
	t.SkipNow()
	input := `
a to b = 1
a to c = 3
a to d = 1000
b to c = 2
b to d = 10
c to d = 1000
`
	// c -> a -> b -> d = 3 + 1 + 10 = 14
	want := 14

	got, err := Day9Part1(input)
	if err != nil {
		t.Errorf("Day9Part1() error = %v", err)
		return
	}

	if got != want {
		t.Errorf("Day9Part1() = %v, want %v", got, want)
	}
}

func TestDay9Part2(t *testing.T) {
	input := `Dublin to Belfast = 141
London to Dublin = 464
London to Belfast = 518
`
	want := 982

	got, err := Day9Part2(input)
	if err != nil {
		t.Errorf("Day9Part2() error = %v", err)
		return
	}

	if got != want {
		t.Errorf("Day9Part2() = %v, want %v", got, want)
	}
}

func Test_parseRoute(t *testing.T) {
	tests := []struct {
		input   string
		want    route
		wantErr bool
	}{
		{"London to Dublin = 464", route{from: "London", to: "Dublin", distance: 464}, false},
		{"bla", route{}, true},
		{"to=", route{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := parseRoute(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
