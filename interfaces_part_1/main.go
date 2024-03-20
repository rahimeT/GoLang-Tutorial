package main

import "fmt"

type Cluster struct {
	clusterName string
	ClusterCount int
}

type ManagedService struct {
	ClusterCount int
}

type ICoster interface {
	ICost() int
}

func(cluster *Cluster) ICost() int{
	return cluster.ClusterCount*2 
}

func (ms *ManagedService) ICost() int {
	return ms.ClusterCount*3
}


func main() {	

	var product ICoster
	product = &Cluster{clusterName:"test_cluster",ClusterCount: 10}
	fmt.Println("cluster cost",product.ICost())

	clusters := []Cluster{Cluster{clusterName:"test_cluster",ClusterCount: 10}, Cluster{clusterName:"dev_cluster", ClusterCount: 20}}
	fmt.Println("cluster count cost",calculateTotalICost(clusters))


	mss := []ManagedService{ManagedService{ClusterCount: 5}}
	fmt.Println("managed service cost",calculateTotalMsICost(mss))
	// ms := ManagedService{ClusterCount: 5}
	// fmt.Println(ms.MsCountCost())
}

func calculateTotalICost(cluster []Cluster) int {
	total := 0

	for _, c := range cluster {
		total += c.ICost()
	}

	return total
}

func calculateTotalMsICost(Ms []ManagedService) int {
	total := 0

	for _, c := range Ms {
		total += c.ICost()
	}

	return total
}