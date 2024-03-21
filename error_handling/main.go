package main

import (
	"fmt"
)

func main(){


	// var x int
	// var y float32
	// fmt.Println(x) //? değişken tanımlanmazsa 0 döner.
	// fmt.Println(y) //? Go dilinde ilkel veri tipleri için varsayılan değerler vardır. int için 0, float için 0.0, string için "" ve bool için false.

	// var pointer *int
	// fmt.Println(pointer) //? pointer için nil döner. Referans tipleri için varsayılan değer nil'dir.

	// if pointer == nil {
	// 	fmt.Println("Pointer herhangi bir yeri göstermiyor")
	// } else {
	// 	fmt.Println("Pointer is not nil")
	// }	

// 	//? Error handling
// 	fileData,err := os.ReadFile("sample.txt") //? Dosya bulunamadığı için hata döner. Bu durumda program durur ve aşağıdaki kodlar çalışmaz.
// if err != nil {
// 	fmt.Println(err)
// 	return
// }else{
// 	fmt.Println(string(fileData))} //* If else yapısı ile yapılan hata kontrolü sonucunda kod spagetthiye dönüşebilir.

	// result, err := divide(10,0)

	// if err != nil {
	// 	fmt.Println(err)
		
	// }else{
	// 	fmt.Println(result)
	// }

	product := Product{id: 1, name: "", price: -1000}
	productService := ProductService{}
	err := productService.SaveProduct(product)
	fmt.Println(err)
}

// func divide(x int, y int) (int, error){
// 	if y == 0 {
// 		return 0, errors.New("0'a bölme hatası")
// 	}
// 	return x / y, nil
// }

//* Product Service örneği
type Product struct {
	id int
	name string
	price int
}

type ProductService struct {

}
type ValidationError struct {
	code string
	message string
}

func (ValidationError ValidationError) Error() string {
	return fmt.Sprintf("Error code: %s, Error message: %s", ValidationError.code, ValidationError.message)
}

 //? Validasyon işlemleri için error
func (p *ProductService) SaveProduct(product Product) error {
	
	if len(product.name) == 0{
		return ValidationError{code: "400", message: "Ürün adı boş olamaz"}
	}

	if product.price < 0 {
		return ValidationError{code: "401", message: "Ürün fiyatı 0'dan küçük olamaz"}
	}
	return nil
}

