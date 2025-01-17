package main

import "fmt"

func main() {
	// x := [4]int{5, 6, 7, 8}
	// y := x[:2]
	// z := x[2:]
	// x[0] = 10
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)

}
