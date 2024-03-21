package main

import (
	"fmt"
	"mymodule/helper"
	"mymodule/helper/rest"
	//* rest2 "mymodule/helper2/rest"  package yapıları aynı olduğu için rest2 olarak import edildi.
)
func main() {
	fmt.Println("Hello, World!")
	helper.Helper1()
	rest.Rest1()
	//! rest.rest1() //! rest1 fonksiyonu dışarıdan erişilemez olduğu için hata alırız.
}