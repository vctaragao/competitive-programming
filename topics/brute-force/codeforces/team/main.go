package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d\n", &N)

	sol := 0
	for i := 0; i < N; i++ {
		var a, b, c int
		fmt.Scanf("%d %d %d\n", &a, &b, &c)

		if a+b+c >= 2 {
			sol++
		}
	}

	fmt.Println(sol)
}
