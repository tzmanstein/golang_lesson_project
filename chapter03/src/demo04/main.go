package main

import (
	"fmt"
	"unsafe"

)
import "./model"

func main() {
	////var i int = 0
	//i := 0
	//i = 20
	//i = 30
	////i = 1.12
	//
	//fmt.Println(i)

	//i := 1
	//j := 2
	//r := i + j
	//fmt.Print(r)
	//

	var i int = 1
	fmt.Println("i=",i)

	var j int8 = 127
	fmt.Println("j=",j)

	var ui uint8 = 255
	fmt.Println("ui=", ui)

	var a int = 8080
	fmt.Println("a=", a)

	var b uint = 1
	var c byte = 255 //equal to uint8
	fmt.Println("b=", b, "c=", c)

	var n1 = 100 // int 类型取决于系统
	fmt.Printf("n1 type is %T\n", n1)
	//查看一个变量占用的字节大小和数据类型
	var n2 int64 = 10
	fmt.Printf("n2 type is %T, and number of byte is %d\n", n2, unsafe.Sizeof(n2))

	//Golang程序中整形变量在使用时，遵守保小不保大的原则，
//	即：在保证程序正确运行下，尽量使用占用空间小的数据类型
//	var age uint8 = 39
	fmt.Println(model.HeroName)




}
