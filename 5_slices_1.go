// Unlike arrays, the length of a slice can grow and shrink as you see fit.

package main

import ("fmt")

func main(){
	// declare a slice
	// this is a slice with 0 length and 0 capacity
	slice_1 := []int{}
	fmt.Println("slice_1 : ", slice_1)

	// initialize a slice

	slice_2 := []int{1, 2, 3, 4}
	fmt.Println("slice_2 : ", slice_2)

	
}

