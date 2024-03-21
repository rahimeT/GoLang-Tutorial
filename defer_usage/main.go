package main

import "fmt"

// //* Defer anahtar kelimesi, bir fonksiyonun çalışmasını erteler.
// //* Channel açıldıktan sonra kapatılması gerektiğinde kullanılır.
// //* Temizlenmesi gereken veritabanı bağlantıları için kullanılır.
// //* Defer yapısı stack mantığına göre çalışır. Yani en son eklenen ilk çalışır.

// func main(){

// 	defer fmt.Println("Defer1")
// 	defer fmt.Println("Defer2")
// 	defer fmt.Println("Defer3")

// 	fmt.Println("MAIN")
// }
// //? Output: MAIN - Defer3 - Defer2 - Defer1

// //* Defer Örnekleri

//* 1.ÖRNEK
// func main(){

// 	//* bu örnekte return ifadesi çalıştığı için defer çalışmaz. Return ifadesinden önce olan defer stack yapısına eklendiği için çalışır.
// 	var condition = true

// 	defer fmt.Println("Defer1")
// 	if condition{
// 		return
// 	}
// 	defer fmt.Println("Defer2")
// }

//* 2.ÖRNEK
// func main (){
// x :=10
// y := 20

// defer fmt.Println("x:", x)
// defer fmt.Print("y:", y)

// x = 100
// y= 200

// fmt.Println("x:", x)
// fmt.Println("y:", y)
// }

//* 3.ÖRNEK

func main(){
	var  condition = true

	defer cleanup()

	if condition{
		panic("Something went wrong")
	}
	
}

func cleanup(){
	fmt.Println("Cleaning worked ...")
}