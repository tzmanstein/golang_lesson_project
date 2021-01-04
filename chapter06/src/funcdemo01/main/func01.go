package main

import (
	"fmt"
	"funcdemo01/utils"
	//adafa "funcdemo01/utils" 给包取别名，使用时使用别名
)

var age int = test()

func test() int {
	fmt.Println("test")
	return 90
}

// 将函数定义一个别名方便最为参数进行传递
type myFuncType func(int, int ) int

func init() {
	fmt.Println("init()....")
}


func getSum(n1 int, n2 int) int {
	return  n1 + n2

}

// 函数最为一个参数进行传递
func myFun (funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func myFun2(funvar myFuncType, num1 int, num2 int) int {
	return  funvar(num1, num2)
}

func sum(n1 int, args... int) int{
	sum := n1
	//便利args
	for i:= 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

//直至对函数返回值命名
func getSumAndSub(n1 int , n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	var n1 float64 = 5.3
	var n2 float64 = 2.3
	var operator byte = '-'

	result := utils.Calculate(n1,n2,operator)
	fmt.Println("result =", result)


	var n11 int = 10

	result2 := utils.Selfplus1(n11)
	fmt.Println("resutl 2 = ", result2)


	var num1 int = 11
	var num2 int = 10
	res1, res2 := utils.SumSub(num1, num2)

	fmt.Printf("res1 = %v, res2 = %v\n", res1, res2)

	//var num4 int = 4
	utils.Recurtest(4)

	resFbn := utils.FibonacciArray(6)
	fmt.Println("resFbn = ", resFbn)

	resTestfun := utils.TestFunc(5)
	fmt.Println("resTestfun =", resTestfun)

	fmt.Println("Day 1 = ", utils.Peach(1))

	var num3 int = 20

	pp := utils.PointerPara
	//utils.PointerPara(&num3)
	pp(&num3)

	fmt.Println("main num3 = ", num3)
	fmt.Printf("num03 address = %v\n", &num3)

	a := getSum
	fmt.Printf(" a的类型%T, getSum的类型%T\n", a, getSum)

	res22 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res22)

	type myInt int
	var selfInt myInt
	var oriInt int

	selfInt = 10
	oriInt = int(selfInt)

	fmt.Println("selfInt = ", selfInt, "oriInt=", oriInt)



	res3 := myFun2(getSum, 60, 60)
	fmt.Println("res3 = ", res3)


	a1, b1 := getSumAndSub(60, 40)

	fmt.Printf("a1 = %v, b1 = %v\n", a1, b1)


	res4 := sum(10, 20, 30, 40)
	fmt.Println("res4=", res4)




}

