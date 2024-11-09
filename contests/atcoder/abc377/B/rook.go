package main

import "fmt"

var Board = [8][8]rune{}

func main() {
	for i := 0; i < 8; i++ {
		var row string
		fmt.Scanf("%s\n", &row)

		for j, s := range row {
			if s == '#' {
				for k := 0; k < 8; k++ {
					Board[i][k] = 'X'
					Board[k][j] = 'X'
				}
			}
		}
	}

	count := 0
	for _, row := range Board {
		for _, s := range row {

			if s != 'X' {
				count++
			}
		}
	}

	fmt.Println(count)
}
