package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	ans := a * b
	for i := 0; i < 3; i++ {
		r := b % 10
		fmt.Println(a * r)
		b /= 10
	}
	fmt.Println(ans)
}
