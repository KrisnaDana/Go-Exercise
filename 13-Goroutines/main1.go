package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main(){
	
	msg := "Hello"

	wg.Add(1) // add 1  goroutine that below to weight group
	go func(msg string){ // add go routine, go routine like paralism processing in go, it very fast
		fmt.Println(msg)
		wg.Done() // delete goroutine by 1 (decrement)
	}(msg)
	msg = "Goodbye"
	wg.Wait() // wait before program ended until weight group being processing
}