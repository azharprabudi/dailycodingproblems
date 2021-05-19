package productsummaries

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

var inputs = []struct {
	numbers  []int
	expected []int
}{
	{
		numbers:  []int{1, 2, 3, 4, 5},
		expected: []int{120, 60, 40, 30, 24},
	},
	{
		numbers:  []int{3, 2, 1},
		expected: []int{2, 3, 6},
	},
}

func TestProductSummaries(t *testing.T) {
	for seq, input := range inputs {
		t.Run(fmt.Sprintf("test on scenario %d", seq+1), func(t *testing.T) {
			actual := ProductSummaries(input.numbers)
			assert.Equal(t, actual, input.expected)
		})
	}
}
