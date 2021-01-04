package main

import (
	"exec"
	"fmt"
	"stdy"
)

func printArray() {
	var score [5] float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素的值", i+ 1)
		fmt.Scanln(&score[i])
	}

	//打印数组变量
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d] = %v\n", i, score[i])
	}
}

func main() {
	var intAttr[3] int
	// 去地址使用%p,, %2.f取出浮点位数
	fmt.Printf("intAttr[0] addr = %p, intAttr[1]=%p \n",&intAttr[0], &intAttr[1])

	//打印数组
	//printArray()

	//数组的初始化部分
	var numArr01 [3]int = [3]int{1,2,3}
	fmt.Println("numArr01",numArr01)

	var numArr02 = [3]int{5,6,7}
	fmt.Println("numArr-2=", numArr02)

	//使用类型推导第一不定长度数组
	var numArr03 = [...]int{8,9,10,11}
	fmt.Println("numArr03=", numArr03)

	var numArr04 = [...]int{1: 800, 0:900, 2:999}
	fmt.Println("numArr04=", numArr04)

	numArr05 := [...]string{1: "jackey", 0: "thomas", 2: "Edward"}
	fmt.Println("numArr05=", numArr05)

	//使用for range来遍历数组
	for _, v := range numArr04 {
		fmt.Println(v)
	}

	exec.TraversalHeros()

	stdy.ArrayPratice()

	stdy.Test01(numArr01)
	fmt.Println(numArr01)

	//数组值地址打印时必须明确使用出地址
	fmt.Printf("&numArr01=%p\n",&numArr01)
	stdy.Test02(&numArr01)
	fmt.Println("Main Test02", numArr01)

	exec.EnterChar()
	exec.GetMaxvalue()

	var intArrys2 = [...]int{1, -1, 9, 90, 11}

	exec.SumandAverage(&intArrys2)

	stdy.ReverseRandom()

}
