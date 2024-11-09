package main

import "fmt"

func main() {
	var X, Y int
	fmt.Scanf("%d %d\n", &X, &Y)

	s := X * Y
	if (s)%2 == 0 {
		fmt.Println(s / 2)
	} else {
		fmt.Println((s - 1) / 2)
	}
}
