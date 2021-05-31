package main

import "fmt"

func main() {
	range1()
	range2()
}

func range1() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// similar to python enumerate
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func range2() {
	pow := make([]int, 10)
	for i := range pow { // index only
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow { // only use value
		fmt.Printf("%d\n", value)
	}
}
