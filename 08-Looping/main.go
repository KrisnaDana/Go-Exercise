package main

import (
	"fmt"
)

func main(){

	
	for i:=0; i<5;i++{
		fmt.Println(i)
	}

	fmt.Println()

	j := 0
	for j<1 {
		fmt.Println(j)
		j++
	}
	
	fmt.Println()

	for a, b := 0, 0; a < 5; a, b = a+1, b+1{
		fmt.Println(a, b)
	}

	fmt.Println()
	// foreach
	// k = key, v = value, key of map = string
	s := []int{1,2,3,4,5}
	for k, v := range s{
		fmt.Println(k, v)
	}

	// use underscore (_) to avoid error compile that variable declared but not used
	for _, v := range s{
		fmt.Println(v)
	}
}