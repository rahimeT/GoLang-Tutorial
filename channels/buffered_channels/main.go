package main

import (
	"fmt"
	"time"
)

func main(){
	// myChannel := make(chan int, 2)
	
	// myChannel <- 42
	// myChannel <- 100
	
	
	// fmt.Println(<-myChannel)
	// myChannel <- 200
	
	// fmt.Println(<-myChannel)
	// myChannel <- 300
	// fmt.Println(<-myChannel)
	
	
	//  message := make(chan string, 2)
	
// go func(){
// 	data1 := <-message
// 	fmt.Println("First", data1)
// 	data2 := <-message
// 	fmt.Println("Second", data2)

// }()

// message <- "Hello"
// message <- "World"
// message <- "World2"
// time.Sleep(time.Second * 1)
// fmt.Println("Done")


bufferedChannel := make(chan int, 4)


go func(){
	for i := 0; i < 10; i++{
		bufferedChannel <- i
		fmt.Println("sent", i)
	}

	close(bufferedChannel)
}()

time.Sleep(time.Second * 3)
for data := range bufferedChannel{
	fmt.Println("received", data)
	time.Sleep(time.Second * 1)
}

}