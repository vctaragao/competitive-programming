package main

import "testing"

type (
	TestCase struct {
		length int
		seats  string
		out    int
	}
)

func TestSeats(t *testing.T) {
	testCases := []TestCase{
		{
			length: 6,
			seats:  "#.##.#",
			out:    2,
		},
		{
			length: 1,
			seats:  "#",
			out:    0,
		},
		{
			length: 9,
			seats:  "###.#.#.##",
			out:    3,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if result := seats(tc.length, tc.seats); result != tc.out {
				t.Fatalf("want %d, got %d\n", tc.out, result)
			}
		})
	}

}
