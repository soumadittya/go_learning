// arrays with length not mentioned

package main

import ("fmt")

func main(){

	// method 1
	var arr1 = [...]int{1,2,3}

	// method 2
  	arr2 := [...]int{4,5,6,7,8}

  	fmt.Println(arr1)
  	fmt.Println(arr2)

}