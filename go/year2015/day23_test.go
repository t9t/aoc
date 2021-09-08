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

func Test_runProgram(t *testing.T) {
	type args struct {
		instructions []instruction
		a            int
		b            int
	}

	single := func(s string, regA bool) []instruction {
		return []instruction{{instruction: s, regA: regA}}
	}

	incA := instruction{instruction: "inc", regA: true}

	tests := []struct {
		name  string
		args  args
		wantA int
		wantB int
	}{
		{"hlf a", args{single("hlf", true), 10, 30}, 5, 30},
		{"hlf b", args{single("hlf", false), 10, 30}, 10, 15},
		{"hlf uneven", args{single("hlf", true), 7, 10}, 3, 10},
		{"inc a", args{single("inc", true), 7, 13}, 8, 13},
		{"inc b", args{single("inc", false), 7, 13}, 7, 14},
		{"tpl a", args{single("tpl", true), 11, 15}, 33, 15},
		{"tpl b", args{single("tpl", false), 11, 15}, 11, 45},

		{"jmp +3", args{[]instruction{
			incA,                         // +1 = 6
			{instruction: "jmp", arg: 3}, // skip this + next two
			incA,                         // skipped
			incA,                         // skipped
			incA,                         // +1 = 7
		}, 5, 0}, 7, 0},

		{"jie +3/jmp -1", args{[]instruction{
			incA,                                     // first +1 = 5; second +1 = 6
			{instruction: "jie", regA: true, arg: 2}, // first +1, then +2
			{instruction: "jmp", arg: -2},            // jump to first inc a
			incA,                                     // +1 = 7
		}, 4, 0}, 7, 0},

		{"jie -1", args{[]instruction{
			incA, // first +1 = 8; second +1 = 9
			{instruction: "jie", regA: true, arg: -1}, // first jump back; second skipped
			incA, // +1 = 10
		}, 7, 0}, 10, 0},

		{"jio", args{[]instruction{
			incA, // first +1 = 1; second +1 = 2
			{instruction: "jio", regA: true, arg: -1}, // first jump 1 back; second ignored
			{instruction: "jio", regA: true, arg: 10}, // ignored (a=2)
			{instruction: "hlf", regA: true},          // 2/2=1
			{instruction: "jio", regA: true, arg: 2},  // skip this + next
			incA,
			incA, // +1 = 2
		}, 0, 9}, 2, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outA, outB := runProgram(tt.args.instructions, tt.args.a, tt.args.b)
			if outA != tt.wantA {
				t.Errorf("runProgram() outA = %v, want %v", outA, tt.wantA)
			}
			if outB != tt.wantB {
				t.Errorf("runProgram() outB = %v, want %v", outB, tt.wantB)
			}
		})
	}
}
