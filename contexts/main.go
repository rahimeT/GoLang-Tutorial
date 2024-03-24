package main

import (
	"context"
	"fmt"
	"time"
)

//* İşlemin uzun sürme ihtimaline karşın
type Product struct {
	ID int
	Name string
}

var productChannel = make(chan Product)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	
	go getProductDetailsFromExternalService(1)
	select {
	case productDetail := <- productChannel:
		fmt.Println("Product ID: ", productDetail.ID)
	case <-ctx.Done():
		fmt.Println("Timeout exceeded")
	}
	
	fmt.Println("End of the main function")
}


func getProductDetailsFromExternalService (productID int) {
	time.Sleep(5 * time.Second)

	productChannel <- Product{
		ID: productID,
		Name: "Product Name",
	}
}