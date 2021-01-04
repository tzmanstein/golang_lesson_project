package main

import "fmt"

func main() {

	var c1 byte = 'a'
	var c2 byte = '0'

	//直接输出byte值时，输出为对应字符的码值
	fmt.Println("c1=",c1)
	fmt.Println("c2=", c2)

	//如果希望输出对饮的字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c \n", c1, c2)

	//var c3 byte = '南'	//类型长度不够
	//var c3 uint = '南'		//汉字需要一个更长类型
	c3 := '南'

	fmt.Printf("c3= %c\n", c3)

	var c4 uint = 22269	//->国汉字的编码
	fmt.Printf("c4=%c\n", c4)

	//字符类型可以进行运算，相当于一个整数，因为它对应有Unicode码
	var n1 = 10 + 'a'	// 97 + 10 = 107
	fmt.Println("n1=", n1)
	fmt.Printf("n1=%c\n", n1)







}
