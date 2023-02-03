package main

import (
	"fmt"
)

func main(){
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i<10;i++{
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface { // interface name is +er from method or function name
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}