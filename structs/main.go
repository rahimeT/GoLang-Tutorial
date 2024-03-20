package main

import "fmt"

func main(){
	var c1 Customer
	// var c2 Customer
	c1= Customer{id: 1, name: "Ahmet", age: 30, adress: Adress{city: "Istanbul", street: "Bagdat"}}
	// c2= Customer{id: 2, name: "Mehmet", age: 40, adress: Adress{city: "Istanbul", street: "Bagdat"}}

	// fmt.Println(c1)
	// fmt.Println(c2)
	c1.ToString()
	c1.ChangeName("Rahime Türkmen")
	c1.ToString() //* struct yapısı pass by value olduğu için değişiklik yapılmaz


	ChangeNameP(&c1)
	c1.ToString() //* struct yapısı pass by reference olduğu için değişiklik yapılır

}

type Customer struct{
	id int64
	name string
	age int
	adress Adress
}

type Adress struct{
	city string
	street string
}

func (customer *Customer) ToString(){ //* struct method
	fmt.Printf("ID: %d, Name: %s, Age: %d, City: %s, Street: %s \n", customer.id, customer.name, customer.age, customer.adress.city, customer.adress.street)
}

func (customer *Customer) ChangeName(newName string){ //* struct method
	customer.name = newName
}


func ChangeNameP(customer *Customer){
	customer.name = "New Name"
}
