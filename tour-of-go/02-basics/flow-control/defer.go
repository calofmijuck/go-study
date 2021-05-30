package main

import "fmt"

func hello() {
	// defers the execution of a function until the surrounding function returns
	defer fmt.Println("world")

	fmt.Println("hello")
}

func foo() {
	// evaluated immediately, but not executed until surrounding function returns
	i := 1

	defer fmt.Println(i) // 1

	i++
	fmt.Println(i) // 2
}

func deferStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// defer statements are pushed to stack
		// when function returns, statements are popped and executed
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	hello()
	foo()
	deferStack()
}
