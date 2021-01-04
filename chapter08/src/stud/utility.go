package stud

import (
	"fmt"
)

func BubbleSort(arr *[5]int) {

	fmt.Println("排序前arr=", (*arr))

	for i:= 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr) - 1 - i ; j++ {
			if (*arr)[j] > (*arr)[j + 1]{
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}

		}
	}

	fmt.Println("排序后arr=", (*arr))

}

/**
	一个数列； 白，金，紫，青
	从键盘数组一个名称，判断数列中是否包含此名称，顺序查找
	1. 定义一个数组，白，金，紫，青
	2. 控制台数组一个名词，依次比较，返回查找结果
 */
func SequenceFind() {

	heros := [4]string{"White", "Gold", "Purple", "Naviblue" }
	var heroName string = ""
	fmt.Println("Input a name for searching...")
	fmt.Scanln(&heroName)
	i := 0
	index := -1
	for {
		//if heros[i]==heroName{
		//	fmt.Println("Got a hero", heroName)
		//	index = i
		//	break
		//} else if i == (len(heros) - 1) {
		//	fmt.Println("Not Found")
		//	break
		//}

		if heros[i]==heroName{
			index = i
			break
		}

		if i == (len(heros) -1) {
			break
		}
		i++
	}

	if index != -1 {
		fmt.Printf("find name = %v, index=%v\n", heros[index], index)
	} else {
		fmt.Println("Not Found", heroName)
	}

}

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int)   {

	if leftIndex > rightIndex {
		fmt.Println("Not Find")
		return
	}

	middle := (leftIndex + rightIndex)/2

	if (*arr)[middle] > findVal {
		BinaryFind(arr,leftIndex, middle - 1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle + 1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为=%v\n", middle)
		//return middle
	}

}
/**
 * 8.10 二维数组初始化
 */
func MatrixDemo() {
	var matrix [4][6]int

	matrix[1][2] = 1
	matrix[2][1] = 2
	matrix[2][3] = 3

	for i:= 0; i < len(matrix); i++ {
		for j:= 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}

	fmt.Println()
	var arr2 [2][3]int
	arr2 [1][2] = 0

	fmt.Printf("ADDR a[0]=%p\n", &arr2[0])
	fmt.Printf("ADDR a[1]=%p\n", &arr2[1])

	fmt.Printf("ADDR a[0][0]=%p\n", &arr2[0][0])
	fmt.Printf("ADDR a[1][0]=%p\n", &arr2[1][0])

	//var arr3[2][3]int = [2][3]int{{1,2,3}, {4,5,6}}
	var arr3 = [2][3]int{{1,2,3}, {4,5,6}}
	fmt.Println("arr3", arr3)

}

/**
 * 8.1 二维数组遍历
 */

func  TraverseArray()  {
	var arr3 = [2][3]int{{1,2,3}, {4,5,6}}

	for i := 0; i < len(arr3); i++{
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("%v\t", arr3[i][j])
		}
	}

	fmt.Println()

	for i, v := range arr3{
		for j, v2 := range v {
			fmt.Printf("arr3[%v][%v]=%v\n", i, j, v2)
		}
	}
}

