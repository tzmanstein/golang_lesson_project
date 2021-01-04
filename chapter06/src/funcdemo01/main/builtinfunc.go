package main

import (
	"errors"
	"fmt"
)

// 626错误处理
func testError() {
	// 使用defer和recover来捕获和异常处理。使用匿名函数
	defer func() {
		//err := recover()
		//if err != nil {
		if err := recover(); err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

// 自定义错误
// 函数去读取一个配置文件init.conf的信息，
// 如果文件出入不正确，返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return  nil
	} else {
		return  errors.New("读取文件错误..")
	}
}

func test02() {
	err := readConf("configg.ini")
	if err != nil{
		panic(err)
	}
	fmt.Println("test02()继续执行")
}


func main() {
	num1 := 100
	fmt.Printf("num1类型=%T, num1值=%v, num1地址=%v\n", num1, num1,&num1)

	// num2是一个指针，使用new为作成一个指针
	num2 := new(int) //*int
	*num2 = 10
	// num2的类型%T => *int
	// num2的值的 = 地址
	// num2的指针的 = 地址
	fmt.Printf("num2类型=%T, num2值=%v, num2地址=%v, 指向值=%v\n", num2, num2, &num2, *num2)

	// 626错误处理
	testError()



	// 自定义错误
	test02()

	fmt.Println("go ahead to exec code in this line")

}
