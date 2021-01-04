package method_func

import "fmt"

type Person struct {
	Name string
}
// 对于函数如下：
// 普通函数，接受者为值（指针）类型时，不能将指针（值）类型的数据直接传递，反之亦然。

func test01(person Person ) {
	fmt.Println("test01", person.Name)
}

func test02(person *Person ) {
	fmt.Println("test02", person.Name)
}
//对于方法如下：
//接受者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以。
func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03", p.Name)
}

func (p *Person) test04() {
	p.Name = "jessy"
	fmt.Println("test04", p.Name)
}

func Run_methodfunc() {
	person := Person{"thomas"}
	test01(person)
	test02(&person)

}