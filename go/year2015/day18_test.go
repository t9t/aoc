package year2015

import (
	"reflect"
	"testing"
)

func Test_processLightsAnimationStep_workingLights(t *testing.T) {
	initial := [][]bool{
		{false, true, false, true, false, true},
		{false, false, false, true, true, false},
		{true, false, false, false, false, true},
		{false, false, true, false, false, false},
		{true, false, true, false, false, true},
		{true, true, true, true, false, false},
	}
	after1 := [][]bool{
		{false, false, true, true, false, false},
		{false, false, true, true, false, true},
		{false, false, false, true, true, false},
		{false, false, false, false, false, false},
		{true, false, false, false, false, false},
		{true, false, true, true, false, false},
	}
	after2 := [][]bool{
		{false, false, true, true, true, false},
		{false, false, false, false, false, false},
		{false, false, true, true, true, false},
		{false, false, false, false, false, false},
		{false, true, false, false, false, false},
		{false, true, false, false, false, false},
	}
	after3 := [][]bool{
		{false, false, false, true, false, false},
		{false, false, false, false, false, false},
		{false, false, false, true, false, false},
		{false, false, true, true, false, false},
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
	}
	after4 := [][]bool{
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
		{false, false, true, true, false, false},
		{false, false, true, true, false, false},
		{false, false, false, false, false, false},
		{false, false, false, false, false, false},
	}

	tests := []struct {
		name  string
		input [][]bool
		want  [][]bool
	}{
		{"1 step", initial, after1},
		{"2 steps", after1, after2},
		{"3 steps", after2, after3},
		{"4 steps", after3, after4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processLightsAnimationStep(tt.input, false); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processLightsAnimationStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processLightsAnimationStep_stuckLights(t *testing.T) {
	initial := [][]bool{
		{true, true, false, true, false, true},
		{false, false, false, true, true, false},
		{true, false, false, false, false, true},
		{false, false, true, false, false, false},
		{true, false, true, false, false, true},
		{true, true, true, true, false, true},
	}
	after1 := [][]bool{
		{true, false, true, true, false, true},
		{true, true, true, true, false, true},
		{false, false, false, true, true, false},
		{false, false, false, false, false, false},
		{true, false, false, false, true, false},
		{true, false, true, true, true, true},
	}
	after2 := [][]bool{
		{true, false, false, true, false, true},
		{true, false, false, false, false, true},
		{false, true, false, true, true, false},
		{false, false, false, true, true, false},
		{false, true, false, false, true, true},
		{true, true, false, true, true, true},
	}
	after3 := [][]bool{
		{true, false, false, false, true, true},
		{true, true, true, true, false, true},
		{false, false, true, true, false, true},
		{false, false, false, false, false, false},
		{true, true, false, false, false, false},
		{true, true, true, true, false, true},
	}
	after4 := [][]bool{
		{true, false, true, true, true, true},
		{true, false, false, false, false, true},
		{false, false, false, true, false, false},
		{false, true, true, false, false, false},
		{true, false, false, false, false, false},
		{true, false, true, false, false, true},
	}

	tests := []struct {
		name  string
		input [][]bool
		want  [][]bool
	}{
		{"1 step", initial, after1},
		{"2 steps", after1, after2},
		{"3 steps", after2, after3},
		{"4 steps", after3, after4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processLightsAnimationStep(tt.input, true); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processLightsAnimationStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseLightsGrid(t *testing.T) {
	small := `.#.
#.#
##.`
	bigger := `.#.#.#
...##.
#....#
..#...
#.#..#
####..`
	tests := []struct {
		name    string
		input   string
		want    [][]bool
		wantErr bool
	}{
		{"3x3", small, [][]bool{{false, true, false}, {true, false, true}, {true, true, false}}, false},
		{"6x6 1", bigger, [][]bool{
			{false, true, false, true, false, true},
			{false, false, false, true, true, false},
			{true, false, false, false, false, true},
			{false, false, true, false, false, false},
			{true, false, true, false, false, true},
			{true, true, true, true, false, false},
		}, false},
		{"not enough characters in line", "...\n..\n...", nil, true},
		{"invalid character", "...\n.z.\n...", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseLightsGrid(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLightsGrid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLightsGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countOnNeighborLights(t *testing.T) {
	lights := [][]bool{
		{false, true, false},
		{true, false, true},
		{true, true, false},
	}
	tests := []struct {
		name string
		y, x int
		want int
	}{
		{"top-left", 0, 0, 2},
		{"top-right", 0, 2, 2},
		{"bottom-left", 2, 0, 2},
		{"bottom-right", 2, 2, 2},
		{"top-middle", 0, 1, 2},
		{"middle-left", 1, 0, 3},
		{"middle-right", 1, 2, 2},
		{"center", 1, 1, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOnNeighborLights(lights, tt.y, tt.x); got != tt.want {
				t.Errorf("countOnNeighborLights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countLitLights(t *testing.T) {
	tests := []struct {
		name   string
		lights [][]bool
		want   int
	}{
		{"3x3", [][]bool{{false, true, false}, {true, false, true}, {true, true, false}}, 5},
		{"6x6", [][]bool{
			{false, true, false, true, false, true},
			{false, false, false, true, true, false},
			{true, false, false, false, false, true},
			{false, false, true, false, false, false},
			{true, false, true, false, false, true},
			{true, true, true, true, false, false},
		}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLitLights(tt.lights); got != tt.want {
				t.Errorf("countLitLights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isCornerLight(t *testing.T) {
	type args struct {
		y     int
		x     int
		total int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"3x3 top-left", args{0, 0, 3}, true},
		{"9x9 top-left", args{0, 0, 9}, true},
		{"3x3 top-right", args{0, 2, 3}, true},
		{"9x9 top-right", args{0, 8, 9}, true},
		{"3x3 bottom-left", args{2, 0, 3}, true},
		{"9x9 bottom-left", args{8, 0, 9}, true},
		{"3x3 bottom-right", args{2, 2, 3}, true},
		{"9x9 bottom-right", args{8, 8, 9}, true},
		{"3x3 middle", args{1, 1, 3}, false},
		{"3x3 side1", args{1, 2, 3}, false},
		{"3x3 side2", args{2, 1, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCornerLight(tt.args.y, tt.args.x, tt.args.total); got != tt.want {
				t.Errorf("isCornerLight() = %v, want %v", got, tt.want)
			}
		})
	}
}
