package main

import (
	"fmt"
)

func main() {
	// Boolean
	var boolVar bool = true
	fmt.Println(boolVar)

	// Integer
	a := 10
	b := 20

	fmt.Println("Add = ", a+b)
	fmt.Println("Substract = ", a-b)
	fmt.Println("Multiply = ", a*b)
	fmt.Println("Divide = ", a/b)
	fmt.Println("Modulus = ", a%b)

	// Float
	var floatVar float64 = 9999
	fmt.Println()
	fmt.Print("Divide = ", float64(a)/floatVar)

	// String
	stringVar := "Hello, world!"
	fmt.Println()
	fmt.Println(stringVar)

	// Rune
	runeVar := 'h'
	fmt.Println("\n", runeVar)

}