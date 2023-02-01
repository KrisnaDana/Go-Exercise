package main

import (
	"fmt"
)

type power struct{
	number int
	str string
	slices []string
}

func main(){

	// Map
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m["a"])

	m1 := make(map[string]int)
	m1["x"] = 2

	fmt.Println(m1)
	delete(m, "c")
	fmt.Println(len(m))


	// Struct
	p := power{
		number: 3,
		str: "a",
		slices: []string{
			"Hello",
			"World",
		},
	}

	fmt.Println(p.str)
}