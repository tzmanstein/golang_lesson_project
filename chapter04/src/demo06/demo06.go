package main

import "fmt"

func main() {
	a := 100
	fmt.Println("a address = ", &a)
	var ptr *int = &a
	fmt.Println("ptr value = ", *ptr)

	// demo07
	//要求：可以从控制台几首用户信息，姓名，年龄，薪水， 考试
	// 方式fmt.Scanln

	var name string
	var age byte
	var sal float32
	var isPass bool

	//fmt.Println("Enter name")
	//fmt.Scanln(&name)
	//
	//fmt.Println("Enter age")
	//fmt.Scanln(&age)
	//
	//fmt.Println("Enter Salary")
	//fmt.Scanln(&sal)
	//
	//fmt.Println("Enter Pass the test or not")
	//fmt.Scanln(&isPass)
	//
	//fmt.Printf("Name is %v, Age is %v, Salary is %v, Pass the test or not %v\n", name, age ,sal,isPass)

	// 方式2:fmt.Scanf,可以按指定格式输入
	fmt.Println("Enter Ur Name, Age ,Salary, Pass Test, and using space to split")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("Name is %v, Age is %v, Salary is %v, Pass the test or not %v\n", name, age ,sal,isPass)




}
