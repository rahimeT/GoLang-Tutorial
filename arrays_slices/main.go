package main

import "fmt"

func main(){
	// var names [3]string

	// names[0] = "Haso"
	// names[1] = "Meho"
	// names[2] = "Raho"
	// fmt.Println(names)


// var names =[3]string{"Haso", "Meho", "Raho"}
// names[1] = "hasoo"
// fmt.Println(names)
// fmt.Println(names[0:2])


var names =[]string{"Haso", "Meho", "Raho"}

names = append(names, "gokgoz")
names = append(names, "gokgoz2")
fmt.Println(names)
}