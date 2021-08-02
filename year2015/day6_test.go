package year2015

import (
	"reflect"
	"testing"
)

func TestDay6Part1(t *testing.T) {
	tests := []struct {
		input   string
		want    int
		wantErr bool
	}{
		{"turn on 0,0 through 999,999", 1_000_000, false},
		{"toggle 0,0 through 999,0", 1000, false},
		{"turn on 0,0 through 999,999\nturn off 499,499 through 500,500", 999_996, false},
		{"bla", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Day6Part1(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Day6Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day6Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseLightInstruction(t *testing.T) {
	tests := []struct {
		input   string
		wantR   lightInstruction
		wantErr bool
	}{
		{"turn on 0,0 through 999,999", lightInstruction{op: turnLightOn, startX: 0, startY: 0, endX: 999, endY: 999}, false},
		{"toggle 0,0 through 999,0", lightInstruction{op: toggleLight, startX: 0, startY: 0, endX: 999, endY: 0}, false},
		{"turn off 499,499 through 500,500", lightInstruction{op: turnLightOff, startX: 499, startY: 499, endX: 500, endY: 500}, false},
		{"toggle 461,550 through 564,900", lightInstruction{op: toggleLight, startX: 461, startY: 550, endX: 564, endY: 900}, false},
		{"turn off 370,39 through 425,839", lightInstruction{op: turnLightOff, startX: 370, startY: 39, endX: 425, endY: 839}, false},
		{"turn on 599,989 through 806,993", lightInstruction{op: turnLightOn, startX: 599, startY: 989, endX: 806, endY: 993}, false},
		{"bla", lightInstruction{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotR, err := parseLightInstruction(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLightInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("parseLightInstruction() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func Test_parseCoords(t *testing.T) {
	tests := []struct {
		input   string
		wantX   int
		wantY   int
		wantErr bool
	}{
		{"0,0", 0, 0, false},
		{"123,456", 123, 456, false},
		{"999,999", 999, 999, false},
		{"1000,1000", 0, 0, true},
		{"34872,4356890", 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotX, gotY, err := parseCoords(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCoords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotX != tt.wantX {
				t.Errorf("parseCoords() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("parseCoords() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
