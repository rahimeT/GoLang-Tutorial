package rest

import "fmt"

//? Aynı paketteki dosyalar arasında fonksiyonlar birbirlerini çağırabilirler. Bu da package scope kavramını anlamamızı sağlar.
func rest2() {
	fmt.Println("Rest2")

	Rest1() 
	rest1()
}