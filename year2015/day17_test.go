package year2015

import (
	"reflect"
	"testing"
)

func Test_findNumberOfCombinationsOfContainersToFitEggnog(t *testing.T) {
	containers := []int{20, 15, 10, 5, 5}
	eggnog := 25
	want := 4
	if got := findNumberOfCombinationsOfContainersToFitEggnog(containers, eggnog); got != want {
		t.Errorf("findNumberOfCombinationsOfContainersToFitEggnog() = %v, want %v", got, want)
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
