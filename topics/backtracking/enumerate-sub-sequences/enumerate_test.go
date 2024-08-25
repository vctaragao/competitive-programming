package enumeratesubsequences_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	enumeratesubsequences "github.com/vctaragao/competitive-programming/topics/backtracking/enumerate-sub-sequences"
)

func TestEnumerate(t *testing.T) {
	sequence := []int{1, 2, 3}
	subSequences := [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}

	result := enumeratesubsequences.Run(sequence)

	assert.Equal(t, subSequences, result)
}
func TestEnumerate2(t *testing.T) {
	sequence := []int{1, 2, 3, 4}
	subSequences := [][]int{{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 4},
		{1, 3},
		{1, 3, 4},
		{1, 4},
		{2},
		{2, 3},
		{2, 3, 4},
		{2, 4},
		{3},
		{3, 4},
		{4},
	}

	result := enumeratesubsequences.Run(sequence)

	assert.Equal(t, subSequences, result)
}
