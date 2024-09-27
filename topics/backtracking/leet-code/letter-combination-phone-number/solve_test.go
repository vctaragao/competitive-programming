package solve_test

import (
	"fmt"
	"slices"
	"testing"

	solve "github.com/vctaragao/competitive-programming/topics/backtracking/leet-code/letter-combination-phone-number"
)

type TestCase struct {
	in  string
	out []string
}

func TestSove(t *testing.T) {
	testCases := []TestCase{
		{
			in:  "23",
			out: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			in:  "",
			out: []string{},
		},
		{
			in:  "2",
			out: []string{"a", "b", "c"},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := solve.Run(tc.in)

			for _, r := range result {
				if !slices.Contains(tc.out, r) {
					t.Fatalf(fmt.Sprintf("result: %s not found in expected out: %v", r, tc.out))
				}
			}

			for _, o := range tc.out {
				if !slices.Contains(result, o) {
					t.Fatalf(fmt.Sprintf("expect out: %s not found in result: %v", o, result))
				}
			}
		})
	}

}
