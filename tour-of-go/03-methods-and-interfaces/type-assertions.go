package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string) // assert that i holds string, and save it to s as string
	fmt.Println(s)

	s, ok := i.(string) // check if assertion is ok
	fmt.Println(s, ok)

	f, ok := i.(float64) // assertion is not ok so ok holds false, but no panic
	fmt.Println(f, ok)

	f = i.(float64) // assertion is not ok. panic
	fmt.Println(f)
}
