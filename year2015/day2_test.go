package year2015

import (
	"testing"
)

func Test_wrappingPaperNeeded(t *testing.T) {
	type args struct {
		l int
		w int
		h int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2x3x4", args{2, 3, 4}, 58},
		{"1x1x10", args{1, 1, 10}, 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrappingPaperNeeded(tt.args.l, tt.args.w, tt.args.h); got != tt.want {
				t.Errorf("wrappingPaperNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay2Part1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"valid", "2x3x4\n1x1x10", 58 + 43, false},
		{"too many dims", "2x3x4x5", 0, true},
		{"parse error", "z2x3x4", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Day2Part1(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Day2Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day2Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ribbonNeeded(t *testing.T) {
	type args struct {
		l int
		w int
		h int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2x3x4", args{2, 3, 4}, 34},
		{"1x1x10", args{1, 1, 10}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ribbonNeeded(tt.args.l, tt.args.w, tt.args.h); got != tt.want {
				t.Errorf("ribbonNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay2Part2(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"valid", "2x3x4\n1x1x10", 34 + 14, false},
		{"too many dims", "2x3x4x5", 0, true},
		{"parse error", "z2x3x4", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Day2Part2(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Day2Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day2Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
