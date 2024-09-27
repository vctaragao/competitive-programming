package generateparentheses

import (
	"slices"
	"testing"
)

type TestCase struct {
	in  int
	out []string
}

func TestGenerateParenthesis(t *testing.T) {
	testCases := []TestCase{
		{in: 3, out: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{in: 1, out: []string{"()"}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := generateParenthesis(tc.in)

			for _, r := range result {
				if !slices.Contains(tc.out, r) {
					t.Fatalf("Wrong result: missing %s in expected output", r)
				}
			}

			for _, o := range tc.out {
				if !slices.Contains(result, o) {
					t.Fatalf("Wrong result: missing %s in expected output", o)
				}
			}
		})

	}

}
