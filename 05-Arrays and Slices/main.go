package main

import (
	"fmt"
)

func main(){

	//ARRAY (fixed size)
	arr := [...]int{1,2,3,4,5} // [...] undefined length arrays (must assigning)
	var arr3 [5]string // decralation not assigning

	fmt.Println(arr)
	fmt.Println(arr3)
	fmt.Println(arr[2])

	arr11 := []int{7,7,7}
	arr12 := arr11 // copying arrays
	arr13 := &arr11 // pointing to arrays
	fmt.Println(arr12)
	fmt.Println(arr13)


	//SLICE (not fixed size) copy slices that pointing to that slices
	sli := []int{1,2,3,4,5}
	sli2 := sli[:2]
	sli2[1] = 0
	fmt.Println(sli2)
	fmt.Println(sli)
	fmt.Println(sli[:]) // similar like python
	fmt.Println(sli[:3])

	slic := make([]int, 3, 100) // (type, length, capacity) -> add big capacity cause more fast operation
	fmt.Println(slic)
	slic = append(slic, 1)
	fmt.Println(slic)
}