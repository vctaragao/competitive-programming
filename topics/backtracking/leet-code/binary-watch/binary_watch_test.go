package binarywatch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	binarywatch "github.com/vctaragao/competitive-programming/topics/backtracking/leet-code/binary-watch"
)

type testCase struct {
	in  int
	out []string
}

func TestSolve(t *testing.T) {
	testCases := []testCase{
		{
			in:  1,
			out: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{in: 9, out: []string{}},
		{in: 0, out: []string{"0:00"}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := binarywatch.Solve(tc.in)
			assert.Equal(t, tc.out, result)
		})
	}
}
