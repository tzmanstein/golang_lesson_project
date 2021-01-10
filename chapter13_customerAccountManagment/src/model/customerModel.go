package model

import "fmt"

//声明和定义一个Customer结构体，表示一个用户信息。
//以上用户信息，取决于需求分析。对客户信息描述
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//编写一个工厂模式，（即是构造函数）返回一个Customer实例
func NewCustomer(id int, name string, gender string,
					age int, phone string, email string) Customer {
	return Customer{
		Id: id,
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

//编写一个工厂模式，（即是构造函数）返回一个Customer实例
func AddCustomer(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		//Id: id,
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

//返回用户的信息，格式化字符串返回
func (this *Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age,
		this.Phone, this.Email)
	return info
}


func Update(customer *Customer) {

	cst := customer

	name := ""
	fmt.Printf("姓名(%v):", cst.Name)
	fmt.Scanln(&name)
	if name != "" {
		customer.Name = name
	}

	gender := ""
	fmt.Printf("性别(%v):", cst.Gender)
	fmt.Scanln(&gender)
	if gender != "" {
		customer.Gender = gender
	}

	age := 0
	fmt.Printf("年龄(%v):", cst.Age)
	fmt.Scanln(&age)
	if age != 0 {
		customer.Age = age
	}

	phone := ""
	fmt.Printf("电话(%v):", cst.Phone)
	fmt.Scanln(&phone)
	if phone != "" {
		customer.Phone = phone
	}

	email := ""
	fmt.Printf("邮箱(%v):", cst.Email)
	fmt.Scanln(&email)
	if email != "" {
		customer.Email = email
	}


}

func (this *Customer)Update_inner() {

	//cst := customer

	name := ""
	fmt.Printf("姓名(%v):", this.Name)
	fmt.Scanln(&name)
	if name != "" {
		this.Name = name
	}

	gender := ""
	fmt.Printf("性别(%v):", this.Gender)
	fmt.Scanln(&gender)
	if gender != "" {
		this.Gender = gender
	}

	age := 0
	fmt.Printf("年龄(%v):", this.Age)
	fmt.Scanln(&age)
	if age != 0 {
		this.Age = age
	}

	phone := ""
	fmt.Printf("电话(%v):", this.Phone)
	fmt.Scanln(&phone)
	if phone != "" {
		this.Phone = phone
	}

	email := ""
	fmt.Printf("邮箱(%v):", this.Email)
	fmt.Scanln(&email)
	if email != "" {
		this.Email = email
	}
}