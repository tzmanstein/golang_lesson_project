package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var bo bool = true
	var mychar byte = 'h'

	var str string 			//空str

	// 方式1: fmt.Sprintf方法转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type = %T, str=%q\n", num1, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type = %T, str=%q\n", num2, str)

	str = fmt.Sprintf("%t", bo)
	fmt.Printf("str type = %T, str=%q\n", bo, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type = %T, str=%q\n", mychar, str)

	//方式2: strconv
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", num3, str)

	str = strconv.Itoa(num3)		//只能接收Int类型，如果十一哦那个int32/int64需要使用强制转换
	fmt.Println("Itoa= ", str)

	// 格式说明： f: 格式 10:小数位保留10位， 64:表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", num4, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type = %T, str=%q\n", b2, str)


}
