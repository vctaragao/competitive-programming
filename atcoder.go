package main

import "fmt"

func main() {
	var N, Q int

	fmt.Scanf("%d %d", &N, &Q)

	var s string

	for i := 0; i < N; i++ {
		s += string('A' + 1)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N-1; j++ {
			fmt.Printf("? %c %c\n", s[j], s[j+1])

			var ans byte
			fmt.Scanf("%c", &ans)
			if ans == '>' {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
