package main

import "fmt"

func main(){
	add(100, 10)
    total, multiply, message := calculation(100, 10)
	fmt.Println("Total:", total, "Multiply:", multiply, "Message:", message)
}

func add(x int, y int) int {
	return x + y;
}

func calculation(x int, y int) (int, int, string) {
	return x / y, x * y, "Result"
}