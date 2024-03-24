package main

import (
	"fmt"
	"time"
)
func main() {
	// Select yapısı ile birden fazla channel üzerinden işlem yapabiliriz.


	// ? Bu yöntem mantıklı bir yöntem değil. Çünkü bir channel üzerinde işlem yapılırken diğer channel beklemeye alınır.
	// channel1 := make(chan string, 1)
	// channel2 := make(chan string)
	// var data string
	// var data2 string
	// go func ()  {
	// 	time.Sleep(10 * time.Second)
	// 	channel1 <- "Data 1"
	// }()
	// go func ()  {
	// 	time.Sleep(1 * time.Second)
	// 	channel2 <- "Data 2"
	// }()

	
	// data2 = <- channel2
	// data = <- channel1

	// fmt.Println("Data1: ", data)
	// fmt.Println("Data2: ", data2)


	// ? Bu yöntemde ise select yapısı ile birlikte channel üzerinde işlem yapılırken diğer channel beklemeye alınmaz.

	channel1 := make(chan string, 1)
	channel2 := make(chan string)

	go func ()  {
		time.Sleep(10*time.Second)
		channel1 <- "Data 1"
	}()

	go func ()  {
		time.Sleep(5*time.Second)
		channel2 <- "Data 2"
	}()
//* data1 ve data2 birer kez yazıldıktan sonra channel' lara veri gönderilmeyeceği için select yapısı içerisindeki default case çalışacaktır.
	for len(channel1) == 0 || len(channel2) == 0 {
	select{
	case data1 := <- channel1:
		fmt.Println("Data1: ", data1)
	case data2 := <- channel2:
		fmt.Println("Data2: ", data2)
    default:
		fmt.Println("Hiçbir veri gelmedi")
	}
	time.Sleep(1 * time.Second)
}
}