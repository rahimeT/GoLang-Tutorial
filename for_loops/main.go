package main

func main(){
// index :=1
// for index<=10{
// 	println(index)
// 	index++
// }


	// total:=0
	// for index:=1; index<=10; index++{
	// total+=index
	// }	
	// fmt.Println(total)

	// index:=0
	// var names = [3]string{"John", "Doe", "Smith"}
	// 	for index<len(names){
	// 		fmt.Println(names[index])
	// 		index++
	// 	}


	for index :=0; index<10; index++{
		if index==5{
			break
		} else if index==3 || index==7{
			continue
		}
		println(index)
	}
}
