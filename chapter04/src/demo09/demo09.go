package main

import "fmt"

func main() {
	var a int = 2
	var b int = 3

	c := a & b

	fmt.Println("result & =", c)
	fmt.Println("result | =", a | b)
	fmt.Println("result ^ =", a ^ b) //亦或，且军备补码运算
	fmt.Println("-2^2=", -2^2)
	//负数的反码为，符号为不变，其他为取反
	//负数的补码为，负数的补码为负数反码+1
	// -2 原码 1000 0010
	// -2 反码 1111 1101
	// -2 补码 1111 1110 (反码+1)

	d := 1 >> 1
	e := 2 << 2

	fmt.Println("d=", d, "e=",e)


}
