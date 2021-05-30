package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// no parens are required, but {} are required
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// can execute a short statement before checking condition
	// scope of v is limited to if/else
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// v doesn't exist anymore
	return lim
}

// exercise
func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	Sqrt(2)
}
