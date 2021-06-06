package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Printf("%d\n%d\n%d\n%d", (a+b)%c, (a+b)%c, (a*b)%c, (a*b)%c)
}
