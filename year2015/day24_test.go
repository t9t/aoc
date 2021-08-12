package year2015

import (
	"math"
	"reflect"
	"testing"
)

func Test_parsePackages(t *testing.T) {
	input := "2\n3\n5\n8\n13\n21\n"
	want := []int{2, 3, 5, 8, 13, 21}
	got, err := parsePackages(input)
	if err != nil {
		t.Errorf("parsePackages() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parsePackages() = %v, want %v", got, want)
	}
}

func Test_sumInts(t *testing.T) {
	tests := []struct {
		name string
		ints []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one", []int{5521}, 5521},
		{"some more", []int{2, 3, 5, 8, 13, 21}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumInts(tt.ints); got != tt.want {
				t.Errorf("sumInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateQuantumEntanglement(t *testing.T) {
	tests := []struct {
		name string
		ints []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one number", []int{42}, 42},
		{"bunch of numbers", []int{2, 3, 5, 8, 13, 21}, 65520},
		{"11 4", []int{11, 4}, 44},
		{"9 4 2", []int{9, 4, 2}, 72},
		{"9 3 2 1", []int{9, 3, 2, 1}, 54},
		{"10 4 3 2 1", []int{10, 4, 3, 2, 1}, 240},
		{"overflow", []int{991234, 991337, 994408, 995521, 999999}, math.MaxInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateQuantumEntanglement(tt.ints); got != tt.want {
				t.Errorf("calculateQuantumEntanglement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFirstSubsetSummingTo(t *testing.T) {
	type args struct {
		target   int
		ints     []int
		maxDepth int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"depth 1", args{7, []int{2, 3, 4, 5, 7}, 1}, []int{7}},
		{"deeper", args{7, []int{2, 3, 4, 5}, 10}, []int{2, 5}},
		{"1-5/7-11 2", args{20, []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}, 2}, []int{9, 11}},
		{"1-5/7-11 3", args{20, []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}, 3}, []int{1, 8, 11}},
		{"unreachable", args{7, []int{9}, 10}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstSubsetSummingTo(tt.args.target, tt.args.ints, 1, tt.args.maxDepth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFirstSubsetSummingTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSmallestSubsetSummingTo(t *testing.T) {
	type args struct {
		target   int
		intsOrig []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"7", args{7, []int{2, 3, 4, 5, 7}}, []int{7}},
		{"9/11", args{20, []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}}, []int{9, 11}},
		{"unreachable", args{7, []int{9}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSmallestSubsetSummingTo(tt.args.target, tt.args.intsOrig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSmallestSubsetSummingTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
