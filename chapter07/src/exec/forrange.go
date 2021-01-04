package exec

import "fmt"

func TraversalHeros() {
	heros := [...]string{"IronMan", "SpiderMan", "BatMan", "SuperMan"}

	for _ ,value :=range heros {
		fmt.Println(value)
	}
}


//数组的应用
// 创建byte类型26个英文字母
// 使用for循环，利用字符可以使用'A'+1 -> B

func EnterChar() {
	var myChar [26]byte

	for i,_ := range myChar {
		myChar[i] = 'A' + byte(i)
	}

	for _,v := range myChar {
		fmt.Printf("%c ",v)
	}
	fmt.Println("")
}


//求数组最大值并给出相应下标
func GetMaxvalue() {
	var intArrys = [...]int{1, -1, 9, 90, 11}

	maxVal := intArrys[0]
	maxValueIndex := 0

	for i,v := range intArrys {
		if maxVal < v {
			maxVal = v
			maxValueIndex = i
		}
	}
	fmt.Printf("maxVal=%v, maxValIndex=%v\n", maxVal, maxValueIndex)

}

func SumandAverage (arr *[5]int) {
	////指针用%p打印时不需要取地址符
	//fmt.Printf("Test02 &arr=%p\n", arr)
	//(*arr)[0] = 88
	//fmt.Println("Test02",arr)
	sum := 0
	var ave float64

	for _, val := range arr {
		sum += val
	}

	ave = float64(sum) / float64(len(arr))

	fmt.Printf("sum=%v, average=%v\n", sum, ave)

}
