package main

import "fmt"

func main() {
	var s string = "Hello 😀"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	// var s2 string = s[4:7]
	// var s3 string = s[:5]
	// var s4 string = s[6:]
	// var b byte = s[6]
	fmt.Println(bs)
	fmt.Println(rs)

	// var a rune = 'x'
	// var s string = string(a)
	// var b byte = byte(a)
	// var s2 string = string(b)
	// fmt.Println(s, s2, b)
}
