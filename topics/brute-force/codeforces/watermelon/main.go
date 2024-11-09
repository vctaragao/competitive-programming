package main

import "fmt"

func main() {
	var W int
	fmt.Scanf("%d\n", &W)

	res := "No"
	if W > 2 && W%2 == 0 {
		res = "Yes"
	}

	fmt.Println(res)
}
