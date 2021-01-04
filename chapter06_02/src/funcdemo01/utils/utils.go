package utils

import (
"fmt"
)

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
