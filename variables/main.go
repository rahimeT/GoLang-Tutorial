package main

import (
	"fmt"
)

func main() {
 //* Variable declaration and assignment
	// var productName string
	// var quantity int
	// var discount float32
	// var isInStock bool

	// productName = "Golang Dersleri"
	// quantity = 10
	// isInStock = true
	// discount = 10.5

    // fmt.Println(productName, reflect.TypeOf(productName))
	// fmt.Println(quantity, reflect.TypeOf(quantity))
	// fmt.Println(discount, reflect.TypeOf(discount))
	// fmt.Println(isInStock, reflect.TypeOf(isInStock))

//* type inference
	// var productName string = "Golang Dersleri"
	// fmt.Println(productName, reflect.TypeOf(productName))

	// var productType = "Golang"
	// fmt.Println(productType, reflect.TypeOf(productType))

	// product := "Golang Dersleri"
	// fmt.Println(product, reflect.TypeOf(product))

	// quantity := 10
	// fmt.Println(quantity, reflect.TypeOf(quantity))

//*specifiers, String formatting
	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang Dersleri"
	quantity = 10
	isInStock = true
	discount = 10.5

	fmt.Println("Product Name:", productName, "quantity:", quantity, "isInStock:", isInStock, "discount:", discount)
	fmt.Printf("Product Name: %v, quantity: %v, isInStock: %v, discount: %v\n", productName, quantity, isInStock, discount) 
    main2();
}

func main2(){
fmt.Println("Hello, playground")
}