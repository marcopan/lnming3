package main

import "fmt"

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

const (
	a = 1 << iota
	b = 1 << iota
	c = 1 << iota
)

func main() {
	fmt.Println("hello")
	fmt.Println(a, b, c)

	var a = [3]int{1, 2, 3}
	var b = &a
	b[1]++
	fmt.Println(a, *b) //same
}
