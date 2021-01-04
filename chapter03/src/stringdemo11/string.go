package main

import "fmt"

func main() {
	// string的基本使用
	var address string = "bejing changcheng dengdeng!"
	fmt.Println("str=", address)

	str := "hello"
	//str[0] = 'a'
	str = "haha"
	fmt.Println(str)

	str2 := "abc\nabc"

	fmt.Println(str2)

	//使用反引号来输出包含有特殊字符的文本
	str3 := `package main
			import (
				"fmt"
				"unsafe"
			)
			
			func main() {
			
				var b = false
				fmt.Println("b=", b)
				fmt.Println("b space is = ", unsafe.Sizeof(b))
			
			}`
	fmt.Println(str3)

	//字符串太长是使用多行字符串
	str4 := "hello " + "kawasaki "
	str4 += "asoboyo "+
		"nabussen " +
		"kenhentohokusen "+
		"tokaidousen "+
		"xiaotianjixian"
	fmt.Println(str4)

	var a int			// 0
	var b float32		// 0.000000
	var c float64		// 0.000000
	var isMarried bool	// false
	var name string		// ""

	fmt.Printf("a=%d, b=%f, c=%f, isMarried=%v, name=%v\n", a, b, c, isMarried, name)







}
