package couplesum

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

var inputs = []struct {
	arr      []int
	k        int
	expected [2]int
}{
	{
		arr:      []int{10, 15, 3, 7},
		k:        17,
		expected: [2]int{10, 7},
	},
	{
		arr:      []int{1, 2, 3, 4},
		k:        4,
		expected: [2]int{1, 3},
	},
	{
		arr:      []int{1, 2, 3, 4},
		k:        5,
		expected: [2]int{2, 3},
	},
}

func TestCoupleSum(t *testing.T) {
	for seq, input := range inputs {
		t.Run(fmt.Sprintf("test on scenario %d", seq+1), func(t *testing.T) {
			actuals := CoupleSum(input.arr, input.k)
			assert.Equal(t, actuals[0], input.expected[0])
			assert.Equal(t, actuals[1], input.expected[1])
		})
	}
}
