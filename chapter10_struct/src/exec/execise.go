package exec

import (
	"fmt"
	"lesson"
	"math"
)

type Circle struct {
	radius float64
}

type integer int

func (i *integer) print() {
	fmt.Println("i=", *i)
}

func (i *integer) change()  {
	*i = *i  + 1
}

type Student struct {
	Name string
	Age int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name, stu.Age)
	return str
}

/**
 * 为了提高效率，通常我们方法和结构体的指针类型绑定。
 */
func (c *Circle) area() float64  {
	// 因为c是指针，因此我们标准的访问字段的方式是(*c).radius
	// 编译器底层优化后则，(*c).radius 等价于 c.radius
	//return math.Pi * (*c).radius * (*c).radius

	fmt.Printf("area c addr %p\n", c)
	return math.Pi*c.radius*c.radius

}

func Exec01 () {

	var p1 lesson.Humman = lesson.Humman{}
	p1.Name = "tom"
	p1.Age = 39

	fmt.Println(p1)

	var p2 *lesson.Humman = &p1
	p2.Name = "thomas"

	fmt.Println(p1)

	fmt.Printf("%p\n", &p1)
	fmt.Printf("p2地址%p, p2的值%p\n", &p2, p2)
}

func Exec_method() {
	var c Circle
	c.radius = 5.0

	fmt.Printf("main c addr = %p\n", &c)
	res := (&c).area()
	//编译器底层作了优化 (&c).area() 等价于c.area()
	fmt.Println("AREA = ", res)
}

func Exec_details() {
	var i integer  = 10

	i.change()
	i.print()

	stu := Student{
		Name: "tom",
		Age: 20,
	}
	fmt.Println(&stu)
}