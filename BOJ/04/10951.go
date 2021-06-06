package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func ReadInt() (int, bool) {
	ok := in.Scan()
	if !ok {
		return -1, false
	}

	value := 0
	for _, c := range in.Bytes() {
		value = 10*value + int(c-'0')
	}
	return value, true
}

func main() {
	in.Split(bufio.ScanWords)

	for {
		a, ok := ReadInt()
		if !ok {
			break
		}
		b, ok := ReadInt()
		if !ok {
			break
		}

		fmt.Fprintln(out, a+b)
	}
	out.Flush()
}
