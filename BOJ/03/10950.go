package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for ; t > 0; t-- {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		fmt.Println(a + b)
	}
}
