// different array functions

package main

import ("fmt")

func main(){

	// access elements of an array
	prices := [3]int{10,20,30}

  	fmt.Println(prices[0])
  	fmt.Println(prices[2])

	// find length of an array
	fmt.Println(len(prices))

	// change elements of an array
	prices[2] = 40
	fmt.Println(prices)

	// initializing only specific elements of an array

	// 1:10 means: assign 10 to array index 1 (second element)
	// 2:40 means: assign 40 to array index 2 (third element)

	num_1 := [5]int{1:10,2:40}
	fmt.Println(num_1)

}