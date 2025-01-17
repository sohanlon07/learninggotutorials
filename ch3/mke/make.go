package main

import "fmt"

func main() {
	x := make([]int, 5)
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))

	y := make([]int, 5, 10)
	fmt.Println(y, len(y), cap(y))

	z := make([]int, 0, 10)
	z = append(z, 5, 6, 7, 8)
	fmt.Println(z, len(z), cap(z))
}
