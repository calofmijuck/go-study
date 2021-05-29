package main

import "fmt"

// https://blog.golang.org/declaration-syntax

func add(x int, y int) int {
	return x + y
}

// x, y are both int
func add2(x, y int) int {
	return x + y
}

// can return multiple values
func swap(x string, y string) (string, string) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 'naked' return
	// Naked return statements should be used only in short functions.
	// They can harm readability in longer functions.
}

func main() {
	fmt.Println(add(42, 13))
}
