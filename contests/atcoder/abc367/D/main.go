package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	in := make([]any, N, N)
	for i := 0; i < N; i++ {
		c := 0
		in[i] = &c
	}

	fmtStr := strings.TrimSpace(strings.Repeat("%d ", N))
	if _, err := fmt.Scanf(fmtStr, in...); err != nil {
		panic(err)
	}

	AS := make([]int, N, N)
	for i, v := range in {
		pv, _ := v.(*int)
		AS[i] = *pv
	}

	solve(N, M, AS)
}

func solve(N, M int, AS []int) {
	countPairs := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("checking pair (%d, %d)\n", i+1, j+1)
			steps := 0
			for k := i + 1; k <= j; k++ {
				fmt.Println("checking for k", k)
				steps += AS[k]
			}

			fmt.Println("steps:", steps)

			if steps%M == 0 {
				fmt.Println("Valid")
				countPairs++
			}
		}
	}

	fmt.Println(countPairs)
}
