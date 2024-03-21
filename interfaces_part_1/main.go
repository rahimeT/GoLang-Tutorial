package main

import "fmt"

//* Interface yapısı, birbiri ile ilişkili farklı tipteki verileri bir arada tutmak için kullanılır.

// //* Cluster, managed service ilişkisi örneği
// type Cluster struct {
// 	clusterName string
// 	ClusterCount int
// }

// type ManagedService struct {
// 	ClusterCount int
// }

// type ICoster interface {
// 	ICost() int
// }

// func(cluster *Cluster) ICost() int{
// 	return cluster.ClusterCount*2
// }

// func (ms *ManagedService) ICost() int {
// 	return ms.ClusterCount*3
// }

// func main() {

// 	var product ICoster

// 	product = &ManagedService{ClusterCount: 5}

// 	product = &Cluster{clusterName:"test_cluster",ClusterCount: 10}
// 	fmt.Println("cluster cost",product.ICost())

// 	clusters := []Cluster{Cluster{clusterName:"test_cluster",ClusterCount: 10}, Cluster{clusterName:"dev_cluster", ClusterCount: 20}}
// 	fmt.Println("cluster count cost",calculateTotalICost(clusters))

// 	mss := []ManagedService{ManagedService{ClusterCount: 5}}
// 	fmt.Println("managed service cost",calculateTotalMsICost(mss))
// 	// ms := ManagedService{ClusterCount: 5}
// 	// fmt.Println(ms.MsCountCost())
// }

// func calculateTotalICost(cluster []Cluster) int {
// 	total := 0

// 	for _, c := range cluster {
// 		total += c.ICost()
// 	}

// 	return total
// }

// func calculateTotalMsICost(Ms []ManagedService) int {
// 	total := 0

// 	for _, c := range Ms {
// 		total += c.ICost()
// 	}

// 	return total
// }

//* Interface kullanmak Strategy Design Pattern

type XEncoder struct {

}

type IEncoder interface {

	Encode(value string)
	Decode(value string)
}


func (xEncoder *XEncoder) Encode(value string)(){
	fmt.Println("XEncoder ile encode edildi")
	
}
func (xEncoder *XEncoder) Decode(value string)(){
	fmt.Println("XEncoder ile decode edildi")
	
}



func main(){
 
	var encoder IEncoder
	encoder = &XEncoder{}
	encoder.Encode("test")
	encoder.Decode("test")

}