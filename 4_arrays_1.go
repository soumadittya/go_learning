// arrays with fixed array length

package main

import ("fmt")

func main(){

	// method 1
	var arr1 = [3]int{1,2,3}

	// method 2
  	arr2 := [5]int{4,5,6,7,8}

  	fmt.Println(arr1)
  	fmt.Println(arr2)

	// different ways to initialize an array
	num_1 := [5]int{} // not initialized
  	num_2 := [5]int{1,2} // partially initialized
  	num_3 := [5]int{1,2,3,4,5} // fully initialized

	// the uninitialized are automatically initilized as 0
	fmt.Println("Array num_1 : ", num_1)
	fmt.Println("Array num_2 : ", num_2)
	fmt.Println("Array num_3 : ", num_3)
}