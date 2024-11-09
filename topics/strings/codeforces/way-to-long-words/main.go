package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d\n", &N)

	for i := 0; i < N; i++ {
		var w string
		fmt.Scanf("%s\n", &w)

		if len(w) > 10 {
			fmt.Printf("%c%d%c\n", w[0], len(w)-2, w[len(w)-1])
		} else {
			fmt.Println(w)
		}
	}
}
