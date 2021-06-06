package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// the `Abs` method has a receiver of type `Vertex` named `v`.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// written as a regular function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods with pointer receivers can modify the value to which the receiver points
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// declare a type and declare a method on non-struct type
// the receiver type should be in the same package as the method
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
