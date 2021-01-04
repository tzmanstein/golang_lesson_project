package stdy

import (
	"fmt"
	"math/rand"
	"time"
)

func ArrayPratice() {

	var arry01 [3] int
	for index, _ := range arry01 {
		arry01[index] = index + 10
	}


	fmt.Println(arry01)

}

func Test01 (arr [3]int) {
	arr[0] = 88
	fmt.Println(arr)
}

func Test02 (arr *[3]int) {
	//指针用%p打印时不需要取地址符
	fmt.Printf("Test02 &arr=%p\n", arr)
	(*arr)[0] = 88
	fmt.Println("Test02",arr)
}

func ReverseRandom () {
	// 1. 随机生成5个数，rand.Intn()函数
	// 2.当我们得到随机数后，就放到一个数组int数组
	// 3.反转数组打印,交换的次数为数组大小除以2.第一与倒数第一交换。
	// 数学归纳法
	var intArr3 [5]int
	//多次引用变量时，使用临时变量提高效率
	length := len(intArr3)
	rand.Seed(time.Now().UnixNano())

	for i,_ := range intArr3 {
		intArr3[i] = rand.Intn(100)
	}

	fmt.Println(intArr3)

	temp := 0
	for i:= 0; i < length / 2; i++ {
		temp = intArr3[length -1 -i]
		intArr3[length - 1 -i] = intArr3[i]
		intArr3[i] = temp
	}

	fmt.Println("Reverse Res =", intArr3)


}