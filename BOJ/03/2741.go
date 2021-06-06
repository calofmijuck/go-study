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
		fmt.Fprintln(out, i)
	}
	out.Flush()
}
