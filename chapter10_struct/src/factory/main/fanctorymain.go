package main

import (
	"factory/model"
	"fmt"
)

func main() {
	//var stu = model.Student{
	//	Name: "tom",
	//	Score: 89.8,
	//}

	//定义student结构体是首字母小写，通过工厂模式来解决

	var stu = model.NewStudent("tome~", 88.8)

	fmt.Println(*stu)
	fmt.Println("Name=", stu.GetName(), "Score=", stu.GetScore())

}