package year2018

import (
	"aoc/registry"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testInput struct {
	input string
	want  string
}

func basicTest(t *testing.T, execution registry.Execution, input, want string) {
	basicMultiTest(t, execution, []testInput{{input: input, want: want}})
}

func basicMultiTest(t *testing.T, execution registry.Execution, tests []testInput) {
	for _, tt := range tests {
		input := strings.TrimSpace(tt.input)
		t.Run(input, func(t *testing.T) {
			got, err := execution(input)
			require.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
