package abc367

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	h := b
	sleeping := []int{}
	for i := b; h != c; i++ {
		h = i % 24
		sleeping = append(sleeping, h)
	}

	for _, h := range sleeping {
		if h == a {
			fmt.Print("No")
			return
		}
	}

	fmt.Print("Yes")
}
