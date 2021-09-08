package year2015

import (
	"reflect"
	"testing"
)

func Test_findNumberOfCombinationsOfContainersToFitEggnogz(t *testing.T) {
	containers := []int{20, 15, 10, 5, 5}
	eggnog := 25
	wantTotal := 4
	wantMin := 3

	total, min := findNumberOfCombinationsOfContainersToFitEggnog(containers, eggnog)
	if total != wantTotal {
		t.Errorf("findNumberOfCombinationsOfContainersToFitEggnog() got = %v, want %v", total, wantTotal)
	}
	if min != wantMin {
		t.Errorf("findNumberOfCombinationsOfContainersToFitEggnog() got1 = %v, want %v", min, wantMin)
	}
}

func Test_parseContainers(t *testing.T) {
	input := "20\n15\n10\n5\n5\n"
	want := []int{20, 15, 10, 5, 5}
	got, err := parseContainers(input)
	if err != nil {
		t.Errorf("parseContainers() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseContainers() = %v, want %v", got, want)
	}
}
