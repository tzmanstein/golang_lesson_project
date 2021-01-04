package utils

import (
	"fmt"
)


func init() {
	fmt.Println("utils package init()")
}
/**
public 函数的函数名首字母为大写
首字符大写函数在go语言中为该函数可导出
 */
func Calculate(n1 float64, n2 float64, operater byte) float64 {

	var res float64
	switch operater {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		if n2 == 0 {
			res = 0.0
			break
		} else {
			res = n1 / n2
		}
	default:
		fmt.Println("操作符号有误...")

	}
	return res

}

func Selfplus1(n11 int ) int {
	var res int = n11
	res += 1
	fmt.Println("Selfplus1 n1= ", res)
	return  res
}


func SumSub(n1 int , n2 int) (int, int) {

	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}
func Recurtest(n int) {

	if n > 2 {
		n--
		Recurtest(n)
	} else {
		fmt.Println("n =======", n)
	}
}

/**
	斐波那契数列
	1, 1, 2, 3, 5, 8, 13,
	递归最重要的是数学表达式
二分查找时使用递归
 */

func FibonacciArray(n int) int {
	if (n ==1 || n == 2 ){
		return 1
	} else {
		return FibonacciArray(n -1) + FibonacciArray(n -2)
	}
}

func TestFunc(n int) int  {
	if n == 1 {
		return 3
	} else {
		return 2 * TestFunc(n -1) + 1
	}
}

/**
	递归调用
	猴子吃桃子，每天吃一半并多吃一个，第10天剩一个还没吃
	问最初有多少个桃子
 */
func Peach(n int) int {
	if n > 10 || n < 1 {
		return 0
	}

	if n == 10 {
		return 1
	} else {
		return (Peach(n + 1) + 1) * 2
	}
}

/**
	引用传递
 */
func PointerPara(n1 *int) {

	*n1 = *n1 + 10
	fmt.Println("PointPara n1 =", *n1)

}

func PrintPyramidExec (totalLevel int) {

	// i 标示层数
	for i := 1; i <= totalLevel; i++ {
		//在打印*前打印空格
		for k := 1; k <= totalLevel - i; k++{
			fmt.Print(" ")
		}

		for j := 1; j <= 2 * i -1; j++ {
			if j == 1 || j == 2 * i - 1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func PrintMultiExec (num int) {

	//打印九九乘法表
	// i表示层数
	for i := 1; i <= num; i++{
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j * i)
		}
		fmt.Println()
	}
}