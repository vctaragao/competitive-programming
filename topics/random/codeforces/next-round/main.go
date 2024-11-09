package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, K int
	fmt.Scanf("%d %d\n", &N, &K)

	g := make([]any, N)
	for i := 0; i < N; i++ {
		c := 0
		g[i] = &c
	}

	fmtStr := strings.TrimSpace(strings.Repeat("%d ", N)) + "\n"
	fmt.Scanf(fmtStr, g...)

	res := 0
	kScore := g[K-1].(*int)

	for _, n := range g {
		i, ok := n.(*int)
		if !ok {
			panic("not int")
		}

		if *i > 0 && *i >= *kScore {
			res++
		}
	}

	fmt.Println(res)
}
