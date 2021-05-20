package univaltree

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

var scenarios = []struct {
	input    *Node
	expected int
}{
	{
		input:    New(0, New(1, nil, nil), New(0, New(1, New(1, nil, nil), New(1, nil, nil)), New(0, nil, nil))),
		expected: 5,
	},
	{
		input:    nil,
		expected: 0,
	},
	{
		input:    New(0, nil, nil),
		expected: 1,
	},

	{
		input:    New(0, New(0, nil, nil), New(0, nil, nil)),
		expected: 3,
	},
	{
		input:    New(0, New(1, New(1, nil, nil), nil), nil),
		expected: 2,
	},
}

func TestCountNumberOfUnivalSubtrees(t *testing.T) {
	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("test on scenario %d", i+1), func(t *testing.T) {

			actual := CountNumberOfUnivalSubtrees(scenario.input)
			assert.Equal(t, actual, scenario.expected)
		})
	}
}
