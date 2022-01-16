// Go has three functions to output text:
// Print()
// Println()
// Printf()

// Print() - function prints its arguments with their default format.

// Println() - function is similar to Print() with the difference that a 
// whitespace is added between the arguments, and a newline is added at the end

// Printf() - function first formats its argument 
// based on the given formatting verb and then prints them.
// %v - is used to print the value of the arguments
// %T - is used to print the type of the arguments

package main

import ("fmt")

func main(){
	var i string = "Hello"
	var j int = 15

  	fmt.Printf("i has value: %v and type: %T\n", i, i)
  	fmt.Printf("j has value: %v and type: %T", j, j)
}

