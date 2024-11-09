package combinationsum

import (
	"slices"
)

// space of possibilities: candidates []int
// at each step I can use every candidate even if it repeats
// 1. Base case: sum == target
// 2. Rule 1 to put a value: value + sum < target

// TODO:
// Create a bench test between the loop implementation and the recursive two step one
// https://leetcode.com/problems/combination-sum/solutions/5426168/video-simply-check-all-combinations/?envType=problem-list-v2&envId=backtracking

func combinationSum(candidates []int, target int) [][]int {
	return bt(0, candidates, target, []int{}, 0, [][]int{})
}

func bt(idx int, c []int, t int, values []int, sum int, results [][]int) [][]int {
	if sum == t {
		return append(results, slices.Clone(values))
	}

	for i := idx; i < len(c); i++ {
		if sum+c[i] > t {
			continue
		}

		values = append(values, c[i])
		results = bt(i, c, t, slices.Clone(values), sum+c[i], results)
		values = values[:len(values)-1]
	}

	return results
}
