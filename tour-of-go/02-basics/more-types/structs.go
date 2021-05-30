package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func literals() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
		vy = Vertex{Y: 2}  // X:0 is implicit
	)

	fmt.Println(v1, p, v2, v3, vy)
}

func main() {
	v := Vertex{1, 2}

	fmt.Println(v)

	v.X = 4 // fields are accessed through dot (.)

	fmt.Println(v)

	p := &v
	p.X = 1e9 // don't have to write (*p).X

	literals()
}
