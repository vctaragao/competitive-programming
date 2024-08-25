package sumofallsubsetxortotals_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	sumofallsubsetxortotals "github.com/vctaragao/competitive-programming/topics/backtracking/leet-code/sum-of-all-sub-set-xor-totals"
)

type testCase struct {
	in  []int
	out int
}

func TestSolve(t *testing.T) {
	testCases := []testCase{
		{in: []int{1, 3}, out: 6},
		{in: []int{5, 1, 6}, out: 28},
		{in: []int{3, 4, 5, 6, 7, 8}, out: 480},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := sumofallsubsetxortotals.Solve(tc.in)
			assert.Equal(t, tc.out, result)
		})
	}
}

func TestSubSequenceXOR(t *testing.T) {
	testCases := []testCase{
		{in: []int{}, out: 0},
		{in: []int{5}, out: 5},
		{in: []int{1}, out: 1},
		{in: []int{6}, out: 6},
		{in: []int{5, 1}, out: 4},
		{in: []int{5, 6}, out: 3},
		{in: []int{1, 6}, out: 7},
		{in: []int{5, 1, 6}, out: 2},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := sumofallsubsetxortotals.SubSequenceXOR(tc.in)
			assert.Equal(t, tc.out, result)
		})
	}
}
