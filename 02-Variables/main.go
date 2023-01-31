package main

import "fmt"

//package scope declaration
var l int = 0

var (
	m int = 1
	n float64 = 2.2
)

var o int = 9

func main(){
	//1st decralation
	var i int
	i = 1

	//2nd decralation
	var j float32 = 10

	//3nd decralation
	k := 100

	fmt.Printf("Value = %v, Type = %T \n", i, i,)
	fmt.Printf("Value = %v, Type = %T \n", j, j)
	fmt.Printf("Value = %v, Type = %T \n", k, k)

	var o int = 10
	fmt.Println("O Value = ", o)

	// naming visibility
	// lowercase scope only same package name example: variable1
	// Uppercase global scope example: Variable1
	// block scope inside function

	//convert
	var x int = 10
	y := float64(x)
	fmt.Printf("y Type = %T", y)
}