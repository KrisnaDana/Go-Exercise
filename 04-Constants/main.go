package main

import (
	"fmt"
)

const (
	i = iota
	ii = iota
)
const iii = iota

const (
	_ = iota + 2 // _ ignore first value by assigning to blank identifier
	k 
	l
)

func main(){
	const a = 10 // without type
	const b int64 = 20



	fmt.Println(i)
	fmt.Println(ii)
	fmt.Println(iii)

	fmt.Println(l)
}