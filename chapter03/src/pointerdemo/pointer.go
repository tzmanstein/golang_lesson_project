package main

import (
	"fmt"
	_ "reflect"
)
// 演示Golang中指针类型
func main() {

	// 基本数据类型在内存中布局
	 var i int = 10
	// i在内存中的地址
	 fmt.Println("i address ", &i)

	// 1. ptr是一个指针变量
	// 2. ptr类型为 *int
	// 3. ptr本身的值&i
	var ptr *int = &i
	fmt.Printf("ptr = %v\n", ptr)
	fmt.Printf("ptr address= %v\n", &ptr)
	fmt.Printf("ptr value= %v\n", *ptr)





}
