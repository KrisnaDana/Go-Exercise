package main

import (
	"fmt"
)

func main(){
	
	a := 5
	b := 7
	c := 3

	if a > b {
		if a > c {
			fmt.Println("a is biggest")
		}else {
			fmt.Println("c is biggest")
		}
	}else if b > a {
		if b > c {
			fmt.Println("b is biggest")
		}else {
			fmt.Println("c is biggest")
		}
	}else {
		fmt.Println("c is biggest")
	}

	m := 20
	switch {
		case m < 20:
			fmt.Println("less than 20")
			// dont need break (implicit break)
		case m == 20:
			fmt.Println("equal to 20")
		case m>= 20:
			fmt.Println("greater than 20")
		default:
			fmt.Println("error")
	}
}