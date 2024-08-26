package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	in := make([]any, N, N)
	for i := 0; i < N; i++ {
		c := 0
		in[i] = &c
	}

	fmtStr := strings.TrimSpace(strings.Repeat("%d ", N))
	if _, err := fmt.Scanf(fmtStr, in...); err != nil {
		panic(err)
	}

	RS := make([]int, N, N)
	for i, v := range in {
		pv, _ := v.(*int)
		RS[i] = *pv
	}

	generate(0, N, K, RS, []int{})
}

func generate(cP, N, K int, RS []int, sequence []int) {
	if cP == N {
		sum := 0
		for _, v := range sequence {
			sum += v
		}

		if sum%K == 0 {
			for i, v := range sequence {
				if i > 0 {
					fmt.Printf(" %d", v)
				} else {
					fmt.Print(v)
				}
			}
			fmt.Printf("\n")
		}

		return
	}

	for r := 1; r <= RS[cP]; r++ {
		sequence = append(sequence, r)
		generate(cP+1, N, K, RS, sequence)
		sequence = sequence[:len(sequence)-1]
	}
}
