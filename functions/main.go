package main

import "fmt"

func main(){
	//add(100, 10)
	var numbers =[]int{1,2,3,100,5,101,7,8,9,10}

    //divison, multiply, message := calculation(100, 10)
	//fmt.Println("", message,"division:", divison, "Multiply:", multiply )
	//fmt.Println(sum(numbers))
	fmt.Println(arr(numbers))
	fmt.Println(factorial(10))
}


//* Basic matematiksel işlemler
func add(x int, y int) int {
	return x + y;
}

func calculation(x int, y int) (int, int, string) {
	return x / y, x * y, "Result"
}

func sum( numbers [] int) int{
	
sum:=0

for _, value := range numbers {
	sum += value
}
return sum
}

//* Dizideki en büyük sayıyı bulan fonksiyon
func arr(nums []int)int{
 temp:=nums[0]
	for i:=0; i<len(nums); i++{
		
		if nums[i]>temp{
			temp=nums[i]
			i++
		}
	}
	return temp
}

//* Faktoriyel hesaplayan fonksiyon
func factorial(x int) int{
	if x==0{
		return 1
	}
	return x*factorial(x-1)
}