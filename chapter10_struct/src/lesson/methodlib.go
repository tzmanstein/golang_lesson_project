package lesson

import (
	"fmt"
)

type Person_method struct {
	Name string
}

func (p Person_method) Test() { //此p为形式参数
	p.Name = "jack" //
	fmt.Println("test() name=", p.Name)
}

func (person_m Person_method) Speak() {
	fmt.Println(person_m.Name, "is a good man")
}

func (person_m Person_method) Calculate() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(person_m.Name, "计算结果是=", res)
}

func (person_m Person_method) Calculate2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(person_m.Name, "计算结果是=", res)
}

func (person_m Person_method) Getsum(n1 int, n2 int) int {
	return n1 + n2
}