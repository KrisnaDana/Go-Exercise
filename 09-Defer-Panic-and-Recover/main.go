package main

import (
	"fmt"
)

func main(){
	

	// Defer => delay
	// defer function => executed while scope function is done but before scope function calling returnss
	fmt.Println("start")
	defer fmt.Println("middle1")
	defer fmt.Println("middle2")
	fmt.Println("end")

	// Defer takes arguments first not takes at the end
	a := "string: start"
	defer fmt.Println(a)
	a = "string: end"

	// panic is not real error but developer determined that error
	// fmt.Println("start")
	// panic("something went wrong")
	// fmt.Println("end")

	// recover() to recover from panic
}