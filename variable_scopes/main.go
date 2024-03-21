package main

import "fmt"

var global ="I am a global variable" //? global scope (accessible from anywhere in the program) but should not change the value of global variable inside the function

func main(){
fmt.Println(global) //* global is accessible here (inside the main function)

	//? Block Scope
	// var condition = true

	// if condition{
	// 	var x=10

	// 	fmt.Println(x)
	// }

	// fmt.Println(x) //* x is not accessible here (outside the if block)
		


	//? Function Scope
	var condition = true

	if condition{
		fmt.Println(condition)
	}

	fmt.Println(condition) //* condition is accessible here (outside the if block)




}