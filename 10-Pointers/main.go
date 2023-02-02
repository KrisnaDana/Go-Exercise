package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main(){

	var a int = 27
	var b *int = &a // *int mean b pointing to int data type, &a mean address of a variable
	fmt.Println(a, *b) // *b mean pointing to pointer value (this case is a variable)
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	var ms *myStruct
	ms = new(myStruct)
	fmt.Println(ms) // same like this: (*ms).foo
	ms.foo = 42 // same like this: (*ms).foo
	fmt.Println(ms.foo) // same like this: (*ms).foo
}