package encapsulate

import "fmt"

type person struct {
	Name string
	age int // 其他包不能直接访问..
	sal float64
}

//写一个工厂模式的函数，相当于构造函数。
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age和salary，编写一对SetXxx和GetXxx的方法

func (per *person) SetAge(age int) {
	if !(age > 0 && age < 150) {
		fmt.Println("年龄异常，不能正常设定")
		return
	}
	per.age = age
}
func (per *person) GetAge() int {
	return per.age
}


func (per *person) SetSalary(sal float64) {
	if sal < 3000 || sal > 300000 {
		fmt.Println("员工薪资异常")
		return
	}
	per.sal = sal

}
func (per *person) GetSalary() float64 {
	return per.sal
}
