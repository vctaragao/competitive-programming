package binarytreepaths_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	binarytreepaths "github.com/vctaragao/competitive-programming/topics/backtracking/leet-code/binary-tree-paths"
)

type testCase struct {
	in  *binarytreepaths.TreeNode
	out []string
}

func TestSolve(t *testing.T) {
	rootNode := &binarytreepaths.TreeNode{
		Val:   1,
		Left:  &binarytreepaths.TreeNode{Val: 2, Right: &binarytreepaths.TreeNode{Val: 5}},
		Right: &binarytreepaths.TreeNode{Val: 3},
	}

	oneNode := &binarytreepaths.TreeNode{Val: 1}

	deeperTree := &binarytreepaths.TreeNode{
		Val: 6,
		Left: &binarytreepaths.TreeNode{
			Val: 1,
			Right: &binarytreepaths.TreeNode{
				Val:  3,
				Left: &binarytreepaths.TreeNode{Val: 2},
				Right: &binarytreepaths.TreeNode{
					Val:  5,
					Left: &binarytreepaths.TreeNode{Val: 4},
				},
			}},
	}

	testCases := []testCase{
		{in: rootNode, out: []string{"1->2->5", "1->3"}},
		{in: oneNode, out: []string{"1"}},
		{in: deeperTree, out: []string{"6->1->3->2", "6->1->3->5->4"}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := binarytreepaths.Solve(tc.in)
			assert.Equal(t, tc.out, result)
		})
	}
}
