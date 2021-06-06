package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(bytes []byte) (int, error) {
	n, err := reader.r.Read(bytes)
	for i := 0; i < n; i++ {
		c := bytes[i]
		if 'A' <= c && c <= 'Z' {
			bytes[i] = 'A' + (bytes[i]-'A'+13)%26
		} else if 'a' <= c && c <= 'z' {
			bytes[i] = 'a' + (bytes[i]-'a'+13)%26
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
