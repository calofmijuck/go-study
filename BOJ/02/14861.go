package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanf("%d\n%d", &x, &y)

	var quadrant int
	if x > 0 {
		if y > 0 {
			quadrant = 1
		} else {
			quadrant = 4
		}
	} else {
		if y > 0 {
			quadrant = 2
		} else {
			quadrant = 3
		}
	}
	fmt.Print(quadrant)
}
