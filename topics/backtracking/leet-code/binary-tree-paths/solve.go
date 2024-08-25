package binarytreepaths

import (
	"slices"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solve(root *TreeNode) []string {
	paths := solve(root, []string{}, [][]string{})

	formattedPaths := []string{}
	for _, path := range paths {
		formattedPaths = append(formattedPaths, strings.Join(path, "->"))
	}

	return formattedPaths
}

func solve(node *TreeNode, path []string, paths [][]string) [][]string {
	if node == nil {
		return paths
	}

	path = append(path, strconv.Itoa(node.Val))

	if node.Left == nil && node.Right == nil {
		paths = append(paths, slices.Clone(path))
	}

	paths = solve(node.Left, path, paths)
	paths = solve(node.Right, path, paths)

	return paths
}
