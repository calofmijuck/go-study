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

	for t := ReadInt(); t > 0; t-- {
		a := ReadInt()
		b := ReadInt()
		fmt.Fprintln(out, a+b)
	}
	out.Flush()
}
