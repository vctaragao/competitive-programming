package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d\n", &N)

	x := 0
	for i := 0; i < N; i++ {
		var s string
		fmt.Scanf("%s\n", &s)

		j := 0
		if s[j] == 'X' {
			j++
		}

		if s[j] == '+' {
			x++
		} else {
			x--
		}
	}

	fmt.Println(x)
}
