package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // mutex: locking and unlocking read and write access

//BASICALLY THIS PROGRAM WILL WASTING FEATURES FROM GOROUTINE CAUSE USE LOCK AND UNLOCK ACCESS THAT BETTER USE NOT GO ROUTINE ON THIS CASE
//USE GOROUTINE ONLY NECESSARY CASE

func main(){
	runtime.GOMAXPROCS(100) // define number threads of os set to runtime (minimum core thread of hardware, ex: 16, max: will increase RAM usage and fast processing), set value to 1 will make it single thread
	for i:=0; i<10; i++{
		wg.Add(2)
		m.RLock() // read lock access
		go sayHello()
		m.Lock() // write lock access
		go increment()
	}
	wg.Wait()
	fmt.Println()
}

func sayHello(){
	fmt.Printf("Hello %v\n", counter)
	m.RUnlock() // read unlock access
	wg.Done()
}

func increment(){
	counter++
	m.Unlock() // write unlock access
	wg.Done()
}