package main

import (
	"fmt"
	"sync"
	"time"
)

//* Concurency, bir programın aynı anda birden fazla işi yapabilmesine olanak tanıyan bir programlama modelidir.
//* Parallelism
//* Goroutine, go dilindeki thread benzeri yapıdır. Bir fonksiyonu go anahtar kelimesi ile başlatarak goroutine oluşturulur.

func main(){
// go f1()
// go f2()
// fmt.Println("main")
//time.Sleep(1 * time.Second) //* main fonksiyonu bitmeden goroutinelerin bitmesini beklemek için kullanılır.


// wg := sync.WaitGroup{}
// wg.Add(1)

// go func (){
// 	defer wg.Done()
// 	fmt.Println("f1")
// }()

// go func (){
// 	defer wg.Done()
// 	fmt.Println("f2")
// }()

// wg.Wait()

// fmt.Println("main")



startTime := time.Now()

wg := sync.WaitGroup{}

wg.Add(3)
go func() {
	defer wg.Done()
	fmt.Println("1. fonksiyon çalıştı")
	time.Sleep(
		3 * time.Second)
}()
go func() {
	defer wg.Done()
	fmt.Println("2. fonksiyon çalıştı")
	time.Sleep(
		3 * time.Second)
}()
go func() {
	defer wg.Done()
	fmt.Println("3. fonksiyon çalıştı")
	time.Sleep(
		3 * time.Second)
}()
wg.Wait()
fmt.Println("Passed time", time.Now().Sub(startTime))

}
