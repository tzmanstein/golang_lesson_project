package main

import "fmt"

var n1 = 100
var n2 = 200
var name = "jack"

var (
	n3 = 300
	n4 = 900
	name2 = "mary"
)

func main() {
	fmt.Println("tom\tjack")
	fmt.Println("hello\nworld")
	fmt.Println("\\Users\\chongzhao\\Project/carrierUp/golang/chapter02/src/escapterchar")
	fmt.Println("tianlongbaquz\nhangfei")

	var num = 2 + 4 * 5
	fmt.Println(num)

	fmt.Println("helloworldhelloworldhelloworld",
		"helloworldhelloworldhelloworldhelloworld",
		"helloworldhelloworldhelloworldhelloworld")

	var i int
	fmt.Println(i)


	var num02 = 10.11
	fmt.Println(num02)

	//下面代码等价于 var name string
	//name = "tom"
	name := "tom"
	num03 := 11.11
	fmt.Println("name= ", name)
	fmt.Println(num03)

	//var n1, n2 ,n3 int
	//fmt.Println("n1 = ", n1, "n2= ", n2, "n3=", n3)
	//
	//var n11, name2, n33 = 100, "thomas", 888
	//fmt.Println("n11 = ", n11, "name ", name2, "n3=", n33)

	//n11, name2, n33 := 100, "thomas", 888
	//fmt.Println("n11 = ", n11, "name ", name2, "n3=", n33)

	//fmt.Println("n1 = ", n1, "n2= ", n2, "n3=", n3)
	fmt.Println("name = ", name, "name2= ", name2, "n4=", 4)


}
