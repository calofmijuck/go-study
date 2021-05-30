package main

import "fmt"

func main() {
	array()
	slice()
	sliceLiterals()
	sliceBounds()
	sliceLenCap()
	makeSlice()
	appendSlice()
}

func array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[1:4] // take elements 1, 2, 3
	fmt.Println(s)

	s[2] = -1           // slices store references!
	fmt.Println(primes) // element 2 is changed
}

func sliceLiterals() {
	// array literal without the length
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceBounds() {
	// just like python
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func sliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	/*
		length of slice is the number of elements in the slice

		capacity is the number of elements in the underlying array,
		counting from the first element in the slice
	*/
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeSlice() {
	// can use dynamically-sized arrays
	a := make([]int, 5)
	printSliceInfo("a", a)

	b := make([]int, 0, 5)
	printSliceInfo("b", b)

	c := b[:2]
	printSliceInfo("c", c)

	d := c[2:5]
	printSliceInfo("d", d)
}

func printSliceInfo(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func appendSlice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}
