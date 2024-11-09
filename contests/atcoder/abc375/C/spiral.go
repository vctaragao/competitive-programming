package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	grid := make([][]rune, N)

	input := bufio.NewReader(os.Stdin)
	for i := 0; i < N; i++ {
		grid[i] = make([]rune, N)

		inputStr, _ := input.ReadString('\n')
		for j, r := range inputStr[:N] {
			grid[i][j] = r
		}
	}

	printSpiral(spiral(grid, N))
}

func spiral(grid [][]rune, N int) [][]rune {
	tmpGrid := make([][]rune, N)
	for i := 0; i < N; i++ {
		tmpGrid[i] = make([]rune, N)
	}

	for i := 0; i < N/2; i++ {
		// should be done simultanesly
		for x := i; x < N-i; x++ {
			for y := i; y < N-i; y++ {
				if i == 1 {
				}
				tmpGrid[y][N-1-x] = grid[x][y]
			}
		}

		for i, row := range tmpGrid {
			copy(grid[i], row)
		}
	}

	return grid
}

func printSpiral(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%s", string(cell))
		}

		fmt.Println()
	}
}
