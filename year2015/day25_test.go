package year2015

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_generateNextManualCode(t *testing.T) {
	tests := []struct {
		previous int
		want     int
	}{
		{20151125, 31916031},
		{31916031, 18749137},
		{18749137, 16080970},
		{16080970, 21629792},
		{21629792, 17289845},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.previous), func(t *testing.T) {
			if got := generateNextManualCode(tt.previous); got != tt.want {
				t.Errorf("generateNextManualCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findManualCode(t *testing.T) {
	table := `
   |    1         2         3         4         5         6
---+---------+---------+---------+---------+---------+---------+
 1 | 20151125  18749137  17289845  30943339  10071777  33511524
 2 | 31916031  21629792  16929656   7726640  15514188   4041754
 3 | 16080970   8057251   1601130   7981243  11661866  16474243
 4 | 24592653  32451966  21345942   9380097  10600672  31527494
 5 |    77061  17552253  28094349   6899651   9250759  31663883
 6 | 33071741   6796745  25397450  24659492   1534922  27995004
	 `

	for i, row := range toManualCodeTestData(table) {
		for j, want := range row {
			rowNum, colNum := i+1, j+1
			t.Run(fmt.Sprintf("row %d col %d", rowNum, colNum), func(t *testing.T) {
				if got := findManualCode(rowNum, colNum); got != want {
					t.Errorf("findManualCode() = %v, want %v", got, want)
				}
			})
		}
	}
}

func toManualCodeTestData(table string) [][]int {
	out := make([][]int, 0)
	for _, line := range strings.Split(strings.TrimSpace(table), "\n")[2:] {
		row := make([]int, 0)
		for _, field := range strings.Fields(strings.Split(line, "|")[1]) {
			n, err := strconv.Atoi(strings.TrimSpace(field))
			if err != nil {
				panic(err)
			}
			row = append(row, n)
		}
		out = append(out, row)
	}
	return out
}

func Test_parseManualCodeInstruction(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantRow    int
		wantColumn int
		wantErr    bool
	}{
		{"1,1", "To continue, please consult the code grid in the manual.  Enter the code at row 1, column 1.", 1, 1, false},
		{"5521,1337", "To continue, please consult the code grid in the manual.  Enter the code at row 5521, column 1337.", 5521, 1337, false},
		{"invalid row", "To continue, please consult the code grid in the manual.  Enter the code at row 3l, column 70.", 0, 0, true},
		{"invalid column", "To continue, please consult the code grid in the manual.  Enter the code at row 31, column 7O", 0, 0, true},
		{"invalid", "bla", 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRow, gotColumn, err := parseManualCodeInstruction(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseManualCodeInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRow != tt.wantRow {
				t.Errorf("parseManualCodeInstruction() gotRow = %v, want %v", gotRow, tt.wantRow)
			}
			if gotColumn != tt.wantColumn {
				t.Errorf("parseManualCodeInstruction() gotColumn = %v, want %v", gotColumn, tt.wantColumn)
			}
		})
	}
}
