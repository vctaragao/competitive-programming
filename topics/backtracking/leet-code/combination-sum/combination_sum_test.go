package combinationsum

import (
	"fmt"
	"slices"
	"testing"
)

type TestCase struct {
	candidates []int
	target     int
	out        [][]int
}

func TestCombinationSum(t *testing.T) {

	testCases := []TestCase{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			out:        [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			out:        [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates: []int{2},
			target:     1,
			out:        [][]int{},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := combinationSum(tc.candidates, tc.target)

			fmt.Println(result)

			for i, o := range tc.out {
				r := result[i]

				if slices.Compare(o, r) != 0 {
					t.Fatalf("results dont match:\nout: %v\nresult:%v", o, r)
				}
			}
		})
	}

}
