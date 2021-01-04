package oop_exec

import "fmt"

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}

func (stu *Student) say() string {
	infoStr := fmt.Sprintf("学生的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		stu.name, stu.gender, stu.age, stu.id, stu.score)
	return infoStr
}

//structdemo
func Run_oopexec() {

	var stu = Student{
		name: "tom",
		gender: "male",
		age: 18,
		id: 1000,
		score: 99.8,
	}
	fmt.Println(stu.say())

}