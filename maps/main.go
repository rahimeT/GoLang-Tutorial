package main

import "fmt"

func main(){
	// var names map[string]int

	// names = make(map[string]int, 0) //* herhangi bir eleman olmadığı için 0 yazdık

	// names["haso"]=1
	// names["raho"]=2
	// names["maho"]=3

	// fmt.Println(names["haso"], names["melike"])

	names := map[string]int{
"haso": 1,
"raho": 2,
"maho": 3,
	}
delete(names, "haso")
	fmt.Println(names)
}