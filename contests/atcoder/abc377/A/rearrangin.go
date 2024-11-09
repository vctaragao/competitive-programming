package main

import (
	"fmt"
	"strings"
)

var seen = map[rune]int{'A': 0, 'B': 0, 'C': 0}

func main() {
	var S string
	fmt.Scanf("%s\n", &S)

	for _, s := range S {
		if !strings.Contains("ABC", string(s)) || seen[s] > 0 {
			fmt.Println("No")
			return
		}

		seen[s]++
	}

	fmt.Println("Yes")
}
