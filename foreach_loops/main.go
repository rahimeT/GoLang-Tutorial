package main

import "fmt"

func main() {
	// var numbers = []int {1, 2, 3, 4, 5}

	// for _, value := range numbers { //* _ ---> index yerine kullanılır.
	// 	fmt.Println( "Value:", value)
	// }


	// var language ="golang"

	// for _, character := range language {
	// 	fmt.Printf("Character: %c\n",	 character)
	// }


	var names = map[string]int{
		"John": 10,
		"Jane": 20,
		"Jack": 30,
	}
//* Dizi ve map objelerinde gezinme yapacağımız zaman foreach döngüsü kullanmak efektiftir.
	for key, value := range names {
		fmt.Println("Key:", key, "Value:", value)
	}
}
