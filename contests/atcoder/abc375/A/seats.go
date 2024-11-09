package main

import "fmt"

func main() {
	var l int
	fmt.Scanf("%d", &l)

	var s string
	fmt.Scanf("%s", &s)

	fmt.Printf("%d\n", seats(l, s))
}

func seats(l int, seats string) int {
	valid := 0

	for i := 0; i < l-2; i++ {
		if seats[i] == '#' && seats[i+1] == '.' && seats[i+2] == '#' {
			valid++
		}
	}

	return valid
}
