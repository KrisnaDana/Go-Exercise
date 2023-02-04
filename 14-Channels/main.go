package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup	{}

func main(){
	ch := make(chan int, 50) // initialize channel with 50 stored value (buffered)
	wg.Add(2)
	go func(ch <- chan int){ // goroutine as receiver channel
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()	
	}(ch)
	go func(ch chan <- int){ // go routine as a sender channel
		ch <- 32
		ch <- 17
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}