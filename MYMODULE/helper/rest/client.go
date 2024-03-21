package rest

import "fmt"

func Rest1(){  //! Eğer fonksiyonun baş harfi büyükse dışarıdan erişilebilir demektir.
	fmt.Println("Rest1")
}

func rest1(){  //! Eğer fonksiyonun baş harfi küçükse dışarıdan erişilemez demektir.
	fmt.Println("rest1")
}