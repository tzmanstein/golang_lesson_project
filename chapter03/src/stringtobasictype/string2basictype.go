package main

import (
	"fmt"
	"strconv"
)
// 演示golang中string转基本数据类型
func main() {

	var str string = "true"
	var b1 bool

	//b1, _ = strconv.ParseBool(str)
	//说明
	// 1. strconv.ParseBool(str)函数会返回俩个值(value bool, err error)
	// 2. 因为只想获得value bool，不关注err，因此使用_忽略掉err返回值
	b1, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T, b=%v\n" ,b1, b1)

	var str2 string = "1234590"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type %T n1 = %v\n", n1, n1)

	n2 = int(n1)
	fmt.Printf("n1 type %T n1 = %v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1 = %v\n", f1, f1)


	var str4 string = "hello"
	var n3 int64 = 11


	n3, _ = strconv.ParseInt(str4, 10, 64) //转换失败，也会对目标进行赋值
	fmt.Printf("n3 type %T, n3 = %v\n", n3, n3)




}
