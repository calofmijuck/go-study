package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func ReadInt() int {
	in.Scan()
	value := 0
	for _, c := range in.Bytes() {
		value = 10*value + int(c-'0')
	}
	return value
}

func main() {
	in.Split(bufio.ScanWords)

	t := ReadInt()
	for i := 1; i <= t; i++ {
		a := ReadInt()
		b := ReadInt()
		fmt.Fprintf(out, "Case #%d: %d\n", i, a+b)
	}
	out.Flush()
}
