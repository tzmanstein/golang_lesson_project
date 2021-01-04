package less_exec

import "fmt"

type MethodUtils struct {
	//item
}

func (mu *MethodUtils) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu *MethodUtils) Print_var(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu *MethodUtils) area(len float64, width float64) float64{
	return len * width
}

func (mu *MethodUtils) judgeNum(num int){
	if num % 2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}


func (mu *MethodUtils) Print_var3(m int, n int, key string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

/**
 * 实现加减乘除： 四个方法
 * 一个方法
 */

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (cal *Calculator) getSum() float64{
	return cal.Num1 + cal.Num2
}

func (cal *Calculator) getSub() float64{
	return  cal.Num1 - cal.Num2
}

func (cal *Calculator) getMul() float64{
	return cal.Num1 * cal.Num2

}

func (cal *Calculator) getChu() float64{
	if cal.Num2 == 0 {
		return 0
	}
	return cal.Num1 / cal.Num2
}

func (cal *Calculator) getRes(operator byte) float64 {
	res := 0.0

	switch operator {
	case '+':
		res = cal.getSum()
	case '-':
		res = cal.getSub()
	case '*':
		res = cal.getMul()
	case '/':
		res = cal.getChu()
	default:
		fmt.Println("Input operator eroor...")
	}

	return res
}

/**
 * Pratice 九九乘法表
 */
type MultiplicationTable struct {
	Num int
}

func (mult *MultiplicationTable)  arithmeticalUnit() {
	fmt.Println(mult.Num)
	for i := 1; i <= mult.Num; i++ {
		for j:= 1; j <= i; j++ {
			//fmt.Printf("#{j} * #{i} = #{j*i} \t")
			fmt.Printf("%d * %d = %d \t", j ,i , j*i)
		}
		fmt.Println()
	}
}

type TraversArray struct {
	Arr33  [3][3]int
	ReArr [3][3]int
}

func (ta * TraversArray) setArray() {
	val := 1
	for i := 0; i< len(ta.Arr33); i++ {
		for j := 0; j < len(ta.Arr33[i]); j++ {
			ta.Arr33[i][j] = val
			ta.ReArr[j][i] = val
			val++
		}
	}
}

func (ta *TraversArray) reverseArray() {
	for i := 0; i< len(ta.Arr33); i++ {
		for j := 0; j < len(ta.Arr33[i]); j++ {
			ta.ReArr[j][i] = ta.Arr33[i][j]
		}
	}

}



func Run_print() {
	var mu MethodUtils
	mu.Print()
	fmt.Println()
	mu.Print_var(2,20)

	mianji := mu.area(4.3, 2.3)
	fmt.Println("mianji=", mianji)

	mu.judgeNum(10)

	mu.Print_var3(4, 8, "#")

	var cal Calculator
	cal.Num1 = 1.2
	cal.Num2 = 2.2
	fmt.Println("sum=", cal.getSum())
	fmt.Printf("sum=%v\n", fmt.Sprintf("%.2f", cal.getSum()))
	fmt.Println("sub=", cal.getSub())
	fmt.Printf("sum=%v\n", fmt.Sprintf("%.2f", cal.getSum()))

	res := cal.getRes('+')
	fmt.Println("res = ", res)

	var multbale MultiplicationTable
	multbale.Num = 9
	multbale.arithmeticalUnit()

	var ta TraversArray
	ta.setArray()
	fmt.Println("Current array=",ta.Arr33)
	ta.reverseArray()
	fmt.Println("Current array=",ta.ReArr)

}

