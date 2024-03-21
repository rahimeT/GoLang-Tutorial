package main

import (
	"fmt"
	"reflect"
)

func main(){
	add := func (x int, y int) int{

		return x+y
	}(3,5) //? Bu şekilde fonksiyonu tanımlayıp, içine parametre alıyorsa veririz ve hemen çalıştırabiliriz.

	add1 := func (x int, y int) int{

		return x+y
	} //? Bu şekilde fonksiyonu tanımlayıp, içine parametre alıyorsa veririz ve hemen çalıştırabiliriz.

	fmt.Println(reflect.TypeOf(add))
	fmt.Println(add1(3,5))
}

