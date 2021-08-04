package year2015

import (
	"reflect"
	"testing"
)

func Test_resolveSignals(t *testing.T) {
	input := `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`

	want := map[string]uint16{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}

	got, err := resolveSignals(input)
	if err != nil {
		t.Errorf("resolveSignals() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("resolveSignals() = %v, want %v", got, want)
	}
}

func Test_toWireMap(t *testing.T) {
	input := `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`

	want := map[string]string{
		"x": "123",
		"y": "456",
		"d": "x AND y",
		"e": "x OR y",
		"f": "x LSHIFT 2",
		"g": "y RSHIFT 2",
		"h": "NOT x",
		"i": "NOT y",
	}

	got, err := toWireMap(input)
	if err != nil {
		t.Errorf("toWireMap() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("toWireMap() = %v, want %v", got, want)
	}
}

func Test_resolveWire(t *testing.T) {
	type args struct {
		def     string
		wireMap map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{"simple value", args{"123", nil}, 123, false},
		{"NOT", args{"NOT 7312", nil}, 58223, false},
		{"AND", args{"3214 AND 7541", nil}, 3076, false},
		{"OR", args{"4312 OR 8421", nil}, 12541, false},
		{"LSHIFT", args{"621 LSHIFT 2", nil}, 2484, false},
		{"RSHIFT", args{"9381 RSHIFT 3", nil}, 1172, false},
		{"simple reference", args{"x", map[string]string{"x": "456"}}, 456, false},
		{"NOT reference", args{"NOT x", map[string]string{"x": "7312"}}, 58223, false},
		{"NOT reference x3", args{"NOT x", map[string]string{"x": "NOT y", "y": "NOT z", "z": "7312"}}, 58223, false},
		{"AND reference", args{"x AND y", map[string]string{"x": "3214", "y": "7541"}}, 3076, false},
		{"OR reference", args{"x OR y", map[string]string{"x": "4312", "y": "8421"}}, 12541, false},
		{"LSHIFT reference", args{"x LSHIFT y", map[string]string{"x": "621", "y": "2"}}, 2484, false},
		{"RSHIFT reference", args{"x RSHIFT y", map[string]string{"x": "9381", "y": "3"}}, 1172, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolveWire(tt.args.def, tt.args.wireMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("resolveWire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("resolveWire() = %v, want %v", got, tt.want)
			}
		})
	}
}
