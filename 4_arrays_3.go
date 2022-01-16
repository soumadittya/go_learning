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

}