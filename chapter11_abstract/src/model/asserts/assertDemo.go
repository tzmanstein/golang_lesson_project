package asserts

import "fmt"

type Point struct {
	x int
	y int
}

func Run_assertDemo() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point

	b = a.(Point) // 类型断言,表示判断a是否指向Point类型变量，如果是就转成Point类型并赋值给b变量，否则报错。

	fmt.Println(b)

	////类型断言的其他案例
	//var x interface{}
	//var b2 float32 = 1.1
	//x = b2 //空接口可以接收任意数据类型
	//y := x.(float32)
	//fmt.Println(y)


	//带有类型检查的断言案例
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //空接口可以接收任意数据类型

	//断言案例(带检测的)
	if y, ok := x.(float32); ok {
		fmt.Println("convert successful")
		fmt.Printf("y的类型是%T,值是%v\n",y,y)
	} else {
		fmt.Println("convert fail")
	}


}