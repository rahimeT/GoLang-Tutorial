package main

import (
	"fmt"
)

//* Channel' lar, goroutine' ler arasında veri alışverişi yapmamızı sağlar.
func main() {
	// myChannel := make(chan string) //* size belirtilmezse default olarak 0 olur. 0 olunca buffer yoktur.

	// go func() {
	// 	myChannel <- "Hello"
	// }()
	// //* Channel' a veri gelene kadar main fonksiyonu bekler.
	// message, isClosed := <-myChannel

	// fmt.Println(message, isClosed)

myChannel := make(chan string)

go func(){
	message := <-myChannel
	fmt.Println(message)
}()

fmt.Println("end of main")
}