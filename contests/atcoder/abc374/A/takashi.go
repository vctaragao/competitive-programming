package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	res := "No"
	if s[len(s)-3:] == "san" {
		res = "Yes"
	}

	fmt.Println(res)
}
