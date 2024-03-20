package main

import "fmt"

func main(){
//  var a int
//  var p *int

//  a = 10

//  p = &a
//  fmt.Println(a)
//  *p = 20
//  fmt.Println(p)
//  fmt.Println(*p)


// //* Adress of operator
// var a=10
// var b int
// var p *int

// b=a
// p=&a
// *p=20
// fmt.Println(a, b, *p)


// var a int = 10
// fmt.Println(a)
// fmt.Println(add12(a))
// fmt.Println(a)

var numbers = []int{1,2,3,4,5,6,7,8,9,10}
fmt.Println(numbers)
changeValue(numbers)
fmt.Println(numbers)
}

func changeValue(numbers []int){ //* slice ve arraylerde pass by reference otomatik olur. Referans tipleri olduğu için
	numbers[0] = 100
}

func add12(x int) int{ //* pass by value
	  x = x + 12
	  return x
}

func add12p(x *int) int{ //* pass by reference
	  *x = *x + 12
	  return *x
}


