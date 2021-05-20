package node

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bmizerany/assert"
)

func TestSerialize(t *testing.T) {
	var scenarios = []struct {
		node     *Node
		expected []string
	}{
		{
			node:     New("10", New("5", New("3", nil, nil), New("4", nil, nil)), New("8", New("6", nil, nil), New("7", nil, nil))),
			expected: []string{"3", "5", "4", "10", "6", "8", "7"},
		},
		{
			node:     nil,
			expected: []string{},
		},
		{
			node:     New("1", nil, nil),
			expected: []string{"1"},
		},
		{
			node:     New("7", New("3", nil, nil), nil),
			expected: []string{"3", "7"},
		},
	}

	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("test on scenario %d", i+1), func(t *testing.T) {
			actual := Serialize(scenario.node)
			assert.Equal(t, actual, scenario.expected)
		})
	}
}

func TestDeserialize(t *testing.T) {
	var scenarios = []struct {
		input    []string
		expected *Node
	}{
		{
			input:    []string{"3", "5", "4", "10", "6", "8", "7"},
			expected: New("10", New("5", New("3", nil, nil), New("4", nil, nil)), New("8", New("6", nil, nil), New("7", nil, nil))),
		},
		{
			input:    []string{},
			expected: nil,
		},
		{
			input:    []string{"1"},
			expected: New("1", nil, nil),
		},
		{
			input:    []string{"3", "7"},
			expected: New("7", New("3", nil, nil), nil),
		},
	}

	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("test on scenario %d", i+1), func(t *testing.T) {
			actual := Deserialize(scenario.input)
			status := reflect.DeepEqual(actual, scenario.expected)
			assert.Equal(t, status, true)
		})
	}
}
