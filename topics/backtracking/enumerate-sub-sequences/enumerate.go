package enumeratesubsequences

import (
	"slices"
)

func Run(sequence []int) [][]int {
	return solve(0, sequence, []int{}, [][]int{})
}

func solve(cI int, sequence []int, subSequence []int, result [][]int) [][]int {
	result = append(result, subSequence)

	for i := cI; i < len(sequence); i++ {
		subSequence = append(subSequence, sequence[i])
		result = solve(i+1, sequence, slices.Clone(subSequence), result)
		subSequence = subSequence[:len(subSequence)-1]
	}

	return result
}
