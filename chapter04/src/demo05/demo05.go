package main

import "fmt"

func main() {
	//var i int
	//i = 10

	a := 9
	b := 2

	fmt.Printf("the value tranfer is a = %v, b = %v\n", a, b)
	tmp := a
	a = b
	b = tmp
	fmt.Printf("after the value tranfer is a = %v, b = %v\n", a, b)

	a += 17
	fmt.Printf("Current a = %v\n", a)

	//2. 赋值运算符的左侧只能是变量，右侧是变量，表达式，常量值
	var d int
	d = a
	d = 8 + 2 * 8
	d +=90
	fmt.Printf("current d = %v\n", d)


}
