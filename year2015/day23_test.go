package year2015

import (
	"reflect"
	"testing"
)

func Test_parseInstruction(t *testing.T) {
	tests := []struct {
		input   string
		wantOut instruction
		wantErr bool
	}{
		{"hlf a", instruction{instruction: "hlf", regA: true, arg: 0}, false},
		{"hlf b", instruction{instruction: "hlf", regA: false, arg: 0}, false},
		{"tpl a", instruction{instruction: "tpl", regA: true, arg: 0}, false},
		{"tpl b", instruction{instruction: "tpl", regA: false, arg: 0}, false},
		{"inc a", instruction{instruction: "inc", regA: true, arg: 0}, false},
		{"inc b", instruction{instruction: "inc", regA: false, arg: 0}, false},
		{"jmp +5", instruction{instruction: "jmp", regA: false, arg: 5}, false},
		{"jmp -10", instruction{instruction: "jmp", regA: false, arg: -10}, false},
		{"jie a, +4", instruction{instruction: "jie", regA: true, arg: 4}, false},
		{"jie a, -11", instruction{instruction: "jie", regA: true, arg: -11}, false},
		{"jie b, +3", instruction{instruction: "jie", regA: false, arg: 3}, false},
		{"jie b, -12", instruction{instruction: "jie", regA: false, arg: -12}, false},
		{"jio a, +2", instruction{instruction: "jio", regA: true, arg: 2}, false},
		{"jio a, -13", instruction{instruction: "jio", regA: true, arg: -13}, false},
		{"jio b, +1", instruction{instruction: "jio", regA: false, arg: 1}, false},
		{"jio b, -14", instruction{instruction: "jio", regA: false, arg: -14}, false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotOut, err := parseInstruction(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("parseInstruction() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
