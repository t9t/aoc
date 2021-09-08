package year2015

import (
	"reflect"
	"testing"
)

func Test_untangleMolecule(t *testing.T) {
	tests := []struct {
		molecule string
		want     []string
		wantErr  bool
	}{
		{"HOH", []string{"H", "O", "H"}, false},
		{"HOOH", []string{"H", "O", "O", "H"}, false},
		{"HOHO", []string{"H", "O", "H", "O"}, false},
		{"OHTieTiHO", []string{"O", "H", "Tie", "Ti", "H", "O"}, false},
		{"BuR", []string{"Bu", "R"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.molecule, func(t *testing.T) {
			got := untangleMolecule(tt.molecule)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("untangleMolecule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseMoleculeAndAtomReplacements(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		molecule     []string
		replacements []atomReplacement
		wantErr      bool
	}{
		{"valid", `H => HO
H => OH
O => HH

HOOHBu`, []string{"H", "O", "O", "H", "Bu"},
			[]atomReplacement{{from: "H", to: []string{"H", "O"}}, {from: "H", to: []string{"O", "H"}}, {from: "O", to: []string{"H", "H"}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMolecule, gotReplacements, err := parseMoleculeAndAtomReplacements(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMoleculeAndAtomReplacements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMolecule, tt.molecule) {
				t.Errorf("parseMoleculeAndAtomReplacements() gotMolecule = %v, want %v", gotMolecule, tt.molecule)
			}
			if !reflect.DeepEqual(gotReplacements, tt.replacements) {
				t.Errorf("parseMoleculeAndAtomReplacements() gotReplacements = %v, want %v", gotReplacements, tt.replacements)
			}
		})
	}
}

func Test_parseAtomReplacement(t *testing.T) {
	tests := []struct {
		input   string
		want    atomReplacement
		wantErr bool
	}{
		{"H => HO", atomReplacement{from: "H", to: []string{"H", "O"}}, false},
		{"Ca => CaCa", atomReplacement{from: "Ca", to: []string{"Ca", "Ca"}}, false},
		{"boo", atomReplacement{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := parseAtomReplacement(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseAtomReplacement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseAtomReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPossibleAtomReplacements(t *testing.T) {
	replacements := []atomReplacement{
		{from: "H", to: []string{"H", "O"}},
		{from: "H", to: []string{"O", "H"}},
		{from: "O", to: []string{"H", "H"}},
	}
	tests := []struct {
		name     string
		molecule []string
		want     int
	}{
		{"HOH", []string{"H", "O", "H"}, 4},
		{"HOHOHO", []string{"H", "O", "H", "O", "H", "O"}, 7},
		{"XHOHOHOZ", []string{"X", "H", "O", "H", "O", "H", "O", "Z"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPossibleAtomReplacements(tt.molecule, replacements); got != tt.want {
				t.Errorf("countPossibleAtomReplacements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readNextAtom(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"A", "A"},
		{"Ke", "Ke"},
		{"AKe", "A"},
		{"KeA", "Ke"},
		{"KeeeA", "Keee"},
		{"AaaaaKe", "Aaaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := readNextAtom(tt.input); got != tt.want {
				t.Errorf("readNextAtom() = %v, want %v", got, tt.want)
			}
		})
	}
}
