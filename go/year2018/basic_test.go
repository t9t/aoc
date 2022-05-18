package year2018

import (
	"aoc/registry"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testInput struct {
	input string
	want  string
}

func basicTest(t *testing.T, execution registry.Execution, tests []testInput) {
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := execution(tt.input)
			require.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
