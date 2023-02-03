package main

import (
	"fmt"
)

func main(){
	a := 10
	fungsi1("Hello, world!", a)

	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		//return	//stop the main function
	}
	fmt.Println(d)
	
	// Anonymous function
	func(a int){
		fmt.Println("Hello ", a)
	}(a)

	// decralating function on variable
	f := func(){ // same like this=> var f func() = func(){}
		fmt.Println("Hello f")
	}
	f()

	//Method
	g := greeter{
		greeting: "Hello",
		name: "Go",
	}
	g.greet()
	fmt.Println("The new name is ", g.name)

}

func fungsi1(str string, a int){
	fmt.Println(str, a)
}

func divide(a, b float64)(float64, error){
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a * b, nil
}

type greeter struct{
	greeting string
	name string
}


// method
func (g greeter) greet(){
	fmt.Println(g.greeting, g.name)
	g.name = ""
}