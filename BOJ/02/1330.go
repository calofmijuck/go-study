package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	s := ""
	if a > b {
		s = ">"
	} else if a < b {
		s = "<"
	} else {
		s = "=="
	}
	fmt.Print(s)
}
