package interfaces

import "fmt"

type A_Interface interface {
	Say()
}

type Stu struct {
	Name string
}
func (stu *Stu) Say() {
	fmt.Println("Stu")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i = ", i)
}

type B_Interface interface {
	Hello()
}

type Monster struct {

}

func (m *Monster) Hello() {
	fmt.Println("Monster Hello()~~")
}

func (m *Monster) Say() {
	fmt.Println("Monster Say()~~")
}

//接口之间继承
type B_ExtendIF interface {
	test01()
}

type C_ExtendIF interface {
	test02()
}

type A_ExtendIF interface {
	B_ExtendIF
	C_ExtendIF
	test03()
}

type Military struct {

}

//需要实现A_ExtendIF下所有接口
func (mit Military) test01() {

}

func (mit Military) test02() {

}

func (mit Military) test03() {

}

//空接口声明定义
type T interface {

}

func Run_interfaceDetails() {
	var stu Stu
	var a A_Interface = &stu
	a.Say()

	var i integer = 10
	var b A_Interface = i
	b.Say()

	//Monster实现了A_Interface，B_Interface
	var monster Monster
	var ai A_Interface = &monster
	var bi B_Interface = &monster

	ai.Say()
	bi.Hello()

	//如果要实现A_ExtendIF，需要把B_ExtendIF,C_ExtendIF的方法都实现
	var m Military
	var aEx A_ExtendIF = m
	aEx.test01()

	//空接口使用
	var t T = m
	fmt.Println(t)
	var num1 float64 = 8.8
	var t2 interface{} = num1
	fmt.Println(t2)



}