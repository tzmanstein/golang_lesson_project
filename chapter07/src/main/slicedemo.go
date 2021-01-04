package main

import (
	"exec"
	"fmt"
	"stdy"
)

func main()  {
	stdy.SliceDemo01()

	//演示slice的make方法使用
	stdy.CreateSlice()

	//append 方法
	stdy.AppendSlice()

	// Slice练习
	exec.Pratice01()

	//string & lice
	stdy.StringToSlice()

	// fbn练习
	fnbSlice := exec.Fbn(10)
	fmt.Println("fbnSlice=", fnbSlice)

}
