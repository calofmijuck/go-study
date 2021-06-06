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

	for i := n; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	out.Flush()
}
