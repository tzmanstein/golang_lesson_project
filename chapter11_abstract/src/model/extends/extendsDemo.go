package extends

import (
	"fmt"
)

//小学生
type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) showInfo() {
	fmt.Printf("学生名=%v, 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

func (stu *Student) getSum(n1 int, n2 int) int {
	return n1 + n2
}


type Pupile struct {
	Student

}

type Graduate struct {
	Student
}

func (p *Pupile) testing() {
	fmt.Println("Pupil testing")
}

func (p *Graduate) testing() {
	fmt.Println("Graduate testing")
}

func Run_extendsDemo (){

	//使用方法1
	pupile := Pupile{}
	pupile.Student.Name = "tom"
	pupile.Student.Age = 8
	pupile.testing()
	pupile.Student.SetScore(90)
	pupile.Student.showInfo()

	pupile.getSum(1,2)

	//使用方法2

	var gradute = Graduate{Student{Name: "jerry", Age: 20}}
	gradute.testing()
	gradute.SetScore(80)
	gradute.showInfo()
}