package main

import (
	"bufio"
	"fmt"
	"os"
)

var out = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Fprint(out, " ")
		}
		for j := 0; j < i; j++ {
			fmt.Fprint(out, "*")
		}
		fmt.Fprint(out, "\n")
	}
	out.Flush()
}
