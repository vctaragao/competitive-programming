package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanf("%d %d\n", &N, &M)

	squares := N * N
	seen := map[string]struct{}{}

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scanf("%d %d\n", &a, &b)

		squares = checkSeen(seen, squares, a, b)

		if a+2 <= N && b+1 <= N {
			squares = checkSeen(seen, squares, a+2, b+1)
		}

		if a+1 <= N && b+2 <= N {
			squares = checkSeen(seen, squares, a+1, b+2)
		}

		if a-1 > 0 && b+2 <= N {
			squares = checkSeen(seen, squares, a-1, b+2)
		}

		if a-2 > 0 && b+1 <= N {
			squares = checkSeen(seen, squares, a-2, b+1)
		}

		if a-2 > 0 && b-1 > 0 {
			squares = checkSeen(seen, squares, a-2, b-1)
		}

		if a-1 > 0 && b-2 > 0 {
			squares = checkSeen(seen, squares, a-1, b-2)
		}

		if a+1 <= N && b-2 > 0 {
			squares = checkSeen(seen, squares, a+1, b-2)
		}

		if a+2 <= N && b-1 > 0 {
			squares = checkSeen(seen, squares, a+2, b-1)
		}

	}

	fmt.Println(squares)
}

func checkSeen(seen map[string]struct{}, squares, a int, b int) int {
	index := fmt.Sprintf("%d%d", a, b)
	if _, ok := seen[index]; !ok {
		squares--
		seen[index] = struct{}{}
	}

	return squares
}
