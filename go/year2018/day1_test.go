package year2018

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Day1Part1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"+1\n-2\n+3\n+1", "3"},
		{"+1\n+1\n+1", "3"},
		{"+1\n+1\n-2", "0"},
		{"-1\n-2\n-3", "-6"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Day1Part1(tt.input)
			require.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_Day1Part2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"+1\n-1", "0"},
		{"+3\n+3\n+4\n-2\n-4", "10"},
		{"-6\n+3\n+8\n+5\n-6", "5"},
		{"+7\n+7\n-2\n-7\n-4", "14"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Day1Part2(tt.input)
			require.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
