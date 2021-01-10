package service

import (
	"model"
)

//该CustomerService，完成对Customer的操作，包括增删改查。
type CustomerService struct {
	customers []model.Customer
	//声明宇哥字段，表示当前切片包含有多少个客户
	//该字段后面，还可以作为新客户的Id+1
	customerNum int
}

//编写一个方法，可以返回CustomService
func NewCustomerService() *CustomerService {
	//为了能够看到客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@shou.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加客户到Customers Slice中
func (this *CustomerService) Add(customer model.Customer) bool {

	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//添加客户到Customers Slice中
func (this *CustomerService) Add2(name string, gender string,
	age int, phone string, email string) bool {

	this.customerNum++
	customer := model.NewCustomer(this.customerNum, name, gender, age, phone, email)
	this.customers = append(this.customers, customer)

	return true
}

func (this *CustomerService) Update(id int) bool {

	index := this.FindById(id)
	if index == -1 {
		// 如果index==-1，相应客户不存在
		return false
	}
	//customer := this.customers[id-1]
	//
	//name := ""
	//fmt.Printf("姓名(%v):", customer.Name)
	//fmt.Scanln(&name)
	//if name != "" {
	//	customer.Name = name
	//}
	//
	//gender := ""
	//fmt.Printf("性别(%v):", customer.Gender)
	//fmt.Scanln(&gender)
	//if gender != "" {
	//	customer.Gender = gender
	//}
	//
	//age := 0
	//fmt.Printf("年龄(%v):", customer.Age)
	//fmt.Scanln(&age)
	//if age != 0 {
	//	customer.Age = age
	//}
	//
	//phone := ""
	//fmt.Printf("电话(%v):", customer.Phone)
	//fmt.Scanln(&phone)
	//if phone != "" {
	//	customer.Phone = phone
	//}
	//
	//email := ""
	//fmt.Printf("邮箱(%v):", customer.Email)
	//fmt.Scanln(&email)
	//if email != "" {
	//	customer.Email = email
	//}
	//this.customers[id-1] = customer

	//model.Update(&this.customers[id-1])
	this.customers[id-1].Update_inner()

	return true
}

// 删除一个客户，查找Id后才能删除
//根据id查找客户对应的slice下标，如果相应客户不存在返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1

	for ind, val := range this.customers {
		if val.Id == id {
			index = ind
			break
		}
	}
	return index
}

func (this *CustomerService) DeleteById(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		// 如果index==-1，相应客户不存在
		return false
	}
	//chapter7
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

