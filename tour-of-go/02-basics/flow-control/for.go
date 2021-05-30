package main

import "fmt"

func main() {
	// no parens around 'initialize; condition; update'
	// instead, {} are required
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// use for instead of while
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite looop
	for {
		// ...
	}
}
