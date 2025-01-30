package main

import (
	"fmt"
)

func main() {
	// totalWins := map[string]int{}
	// teams := map[string]string {
	// 	"Orcas": []string{"Fred", "Ralph", "Bijou"},
	// 	"Lions": []string{"Sarah", "Peter", "Billie"},
	// 	"Kittens": []string{"Waldo", "Raul", "Ze"},
	// }

	// totalWins["Orcas"] = 1
	// totalWins["Lions"] = 2
	// fmt.Println(totalWins["Orcas"])
	// fmt.Println(totalWins["Kittens"])
	// totalWins["Kittens"]++
	// fmt.Println(totalWins["Kittens"])
	// totalWins["Lions"] = 3
	// fmt.Println(totalWins["Lions"])

	m := map[string]int{
		"hello":5,
		"world": 10,
	}
	delete(m, "hello")
	
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)
}
