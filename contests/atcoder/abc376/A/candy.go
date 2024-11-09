package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, C int
	fmt.Scanf("%d %d\n", &N, &C)

	times := make([]int, N)
	input := bufio.NewReader(os.Stdin)
	inputStr, err := input.ReadString('\n')
	if err != nil {
		panic(err)
	}

	inputTimes := strings.Split(inputStr[:len(inputStr)-1], " ")
	for i := 0; i < N; i++ {
		t, err := strconv.Atoi(inputTimes[i])
		if err != nil {
			panic(err)
		}

		times[i] = t
	}

	candies := 0
	pTime := 0

	for i, t := range times {
		if i == 0 || t-pTime >= C {
			candies++
			pTime = t
		}
	}

	fmt.Println(candies)
}
