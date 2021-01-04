package main

import (
	"exec"
	"fmt"
	"stud"
)

func main () {

	//定义一个数组
	arr := [5]int{24, 69, 80, 57, 13}
	//冒泡排序函数调用
	//stud.BubbleSort(&arr)

	//stud.SequenceFind()

	fmt.Println("Outter = ", arr)

	arr2 := [6]int{1, 8, 10, 89, 1000, 1234}
	stud.BinaryFind(&arr2, 0, len(arr2)-1, 9)

	//fmt.Println("BinaryFind=", arr2)

	stud.MatrixDemo()

	stud.TraverseArray()

	exec.Exec01_TraverseMatrix()


}
