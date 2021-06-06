package main

import (
	"fmt"
)

func GetNext(n int) int {
	r := n % 10
	ret := 0
	for n > 0 {
		ret += n % 10
		n /= 10
	}
	return 10*r + ret%10
}

func main() {
	var n int
	fmt.Scan(&n)

	cycle := 0
	next := n
	for {
		cycle++
		next = GetNext(next)
		if n == next {
			break
		}
	}
	fmt.Print(cycle)
}
