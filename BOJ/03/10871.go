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

	n := ReadInt()
	x := ReadInt()
	for i := 0; i < n; i++ {
		if num := ReadInt(); num < x {
			fmt.Fprintf(out, "%d ", num)
		}
	}
	out.Flush()
}
