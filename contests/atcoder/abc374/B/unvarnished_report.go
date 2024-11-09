package main

import (
	"fmt"
)

func main() {
	var S, T string
	fmt.Scanf("%s\n", &S)
	fmt.Scanf("%s\n", &T)

	k := 0
	for i := 0; i < len(S) || i < len(T); i++ {
		if i == len(S) || i == len(T) || S[i] != T[i] {
			k = i + 1
			break
		}
	}

	fmt.Println(k)
}
