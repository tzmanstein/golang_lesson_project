package main

import "fmt"

func main() {

	//aveWeight := fmt.Sprintf("%.2f", 6.4544554)
	//fmt.Printf("aveWeight=%v \n", aveWeight)

	//1. 定义一个数组（值）
	var hens [6]float64
	//2. 给数组的每个元素赋值，初始化数组
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	var totalWeight float64

	for i := 0; i < len(hens); i++{
		totalWeight += hens[i]
	}

	aveWeight := fmt.Sprintf("%.2f", totalWeight / (float64(len(hens))))
	fmt.Printf("totalWeight = %v, aveWeight=%v \n", totalWeight, aveWeight)



}