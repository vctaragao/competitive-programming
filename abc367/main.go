// https://atcoder.jp/contests/abc367/tasks/abc367_c
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	in := make([]any, n, n)
	for i := 0; i < n; i++ {
		c := 0
		in[i] = &c
	}

	fmtStr := strings.TrimSpace(strings.Repeat("%d ", n))
	if _, err := fmt.Scanf(fmtStr, in...); err != nil {
		panic(err)
	}

	rs := make([]int, n, n)
	for i, v := range in {
		pv, _ := v.(*int)
		rs[i] = *pv
	}

	generate(0, 1, rs)
}

func generate(i, k int, s []int) {
	if k == len(s) {
		return
	}

	subSequence := s[i:k]
	fmtStr := strings.TrimSpace(strings.Repeat("%d ", len(subSequence)))
	fmt.Printf(fmtStr, subSequence)

	generate(i, k+1, s)

	subSequence = s[i : k+1]
	fmtStr = strings.TrimSpace(strings.Repeat("%d ", len(subSequence)))
	fmt.Printf(fmtStr, subSequence)
}
