package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)

	k := 3 // short assignment can only be used inside functions
	fmt.Println(k)
}
