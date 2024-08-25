package sumofallsubsetxortotals

func Solve(nums []int) int {
	return solve(0, 0, []int{}, nums)
}

func solve(cI int, acc int, subSequence []int, sequence []int) int {
	acc += SubSequenceXOR(subSequence)

	for i := cI; i < len(sequence); i++ {
		subSequence = append(subSequence, sequence[i])
		acc = solve(i+1, acc, subSequence, sequence)
		subSequence = subSequence[:len(subSequence)-1]
	}

	return acc
}

func SubSequenceXOR(subSequence []int) int {
	if len(subSequence) == 0 {
		return 0
	}

	xorResult := 0
	for i := 0; i < len(subSequence); i++ {
		xorResult ^= subSequence[i]
	}

	return xorResult
}
