package main

import (
	"fmt"
	"reflect"
)

func main(){
	add := func (x int, y int) int{

		return x+y
	} //? Bu şekilde fonksiyonu tanımlayıp, içine parametre alıyorsa veririz ve hemen çalıştırabiliriz.

	multiply := func (x int, y int) int{

		return x*y
	} //? Bu şekilde fonksiyonu tanımlayıp, içine parametre alıyorsa veririz ve hemen çalıştırabiliriz.

	fmt.Println(reflect.TypeOf(add))
	fmt.Println(multiply(3,5))

	a,b :=calculator(3,5,add,multiply)
	fmt.Println(a,b)
}

func calculator(x int, y int, operator func(int, int) int, operator2 func(int,int) int) (int,int){
	
	return operator(x,y), operator2(x,y)
}