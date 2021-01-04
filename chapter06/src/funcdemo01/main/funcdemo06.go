package main

import (
	"fmt"
	"funcdemo01/utils"
	"strings"
)

/**
	匿名函数的意义为在函数中定义一个函数，好像和swift中的一种使用方式类似，想不起来了。
 */
var (
	Fun1 = func (n1 int, n2 int) int {
		return n1 * n2
	}
)

// 全局变量定义时不能包含赋值语句，即不能使用类型推导来定义全局变量
// Name := "tom"
var Name string = "tom"

// 闭包演示案例
func AddUpper() func (int) int {
	var n int = 10 //闭包变量可以保持，不使用全局变量亦可实现此效果
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += "a"
		fmt.Println("str=", str)
		return n
	}
}

func makeSuffix(suffix string) func (string) string {
	// suffix传递变量，作为通用变量使用，闭包中调用中始终有效。
	return func (inputStr string) string {
		if !strings.HasSuffix(inputStr, suffix) {
			return  inputStr + suffix
		}
		return inputStr
	}
}

func sumdefer(n1 int, n2 int) int {
	// 当执行到defer时，暂不执行，会将defer后的语句压入独立的defer栈中
	// 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈执行。
	// 进入defer栈时，将相应的值拷贝入栈，不会根据后续值变化而改变。
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok1 n2=", n2)

	n1++
	n2++

	res := n1 + n2
	fmt.Println("ok3 =", res)
	return res
}

func main() {
	// 在定义醚名函数时直接调用，这种匿名函数只能调用一次

	// 例子求和

	res1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("rest1=", res1)

	// 将匿名函数func (n1 int, n2 int) int 赋值给a变量
	// 则a的数据类型就是函数类型，此时，我们可以通过a来调用
	a := func (n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(30, 20)
	fmt.Println("res2=", res2)

	//全局匿名函数的使用
	res4 := Fun1(4,9)
	fmt.Println("res4=", res4)

	// 使用闭包代码
	f := AddUpper()
	fmt.Println("f(1)", f(1))
	fmt.Println("f(2)", f(2))
	fmt.Println("f(3)", f(3))

	//闭包练习
	//使用闭包时，通用型变量可以只传第一次。
	f1 := makeSuffix(".jpg")
	fmt.Println(f1("winter"))

	//函数defer用法
	res5 := sumdefer(10, 20)
	fmt.Println("res=", res5)

	utils.PrintPyramidExec(9)

	utils.PrintMultiExec(8)
}
