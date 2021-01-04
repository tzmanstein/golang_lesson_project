package lesson

import (
	"encoding/json"
	"fmt"
)



func Sample() {
	//结构体实例化，为值类型。定义即创建
	var cat1 Cat

	fmt.Printf("cat1的地址是%p\n", &cat1)

	cat1.Name = "Socker"
	cat1.Age = 3
	cat1.Color = "White"
	cat1.Hobby = "Eating fish"

	fmt.Println(cat1)

	fmt.Println("The cat info is as below:")
	fmt.Println("Name=", cat1.Name)
}

func Sample02_reftype() {

	var person Student

	// 使用引用类型要先make
	person.ptr = new(int)
	person.slice = make([]int, 10)
	person.slice[0] = 100
	person.map1 = make(map[string]string)
	person.map1["key1"] = "Tom~"

	fmt.Println(person)

	if person.ptr == nil {
		fmt.Println("*ptr is nil")
	}

	if person.slice == nil {
		fmt.Println("slice is nil")
	}

}

func Sample03_encapulate() {
	var monster1 Monster

	monster1.Name = "Bulls"
	monster1.Age = 500

	monster2 := monster1 //值拷贝，monster2中存储一个monster1的拷贝副本。独立存在。
	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)
	monster2.Name = "Tigger"
	fmt.Println("monster2 MODIFY = ", monster2)
}

func Sample04 () {

	//方式2
	p2 := Humman{}
	p2.Name = "tom"
	p2.Age = 18

	//方式2
	p3 := Humman{"mary",20}
	fmt.Println(p2)
	fmt.Println(p3)

	//方式3
	var p4 *Humman =  new(Humman)
	(*p4).Name = "smith"
	//(*p4).Age = 30
	p4.Age = 30 // 在底层编译器有优化处理写法 p4.Age 与 (*p4).Age相同

	fmt.Println(*p4)

	//方式4
	var human *Humman = &Humman{}
	//var human *Humman = &Humman{"marry", 23}

	human.Name = "scoot"
	human.Age = 33
	(*human).Name = "alibaba"
	(*human).Age = 40

	fmt.Println(*human)

}

func Structdetails() {
	r1 := Rect{Point{1,2}, Point{3,4}}
	//r1 中有4个int，在内存中是连续分布
	fmt.Printf("r1.lefUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	//r2有两个*Point类型，这两个*Point类型得到本身的地址也是连续的
	//但他们指向的地址不一定连续。
	r2 := Rect2{&Point{10,20}, &Point{30,40}}

	//本身的地址是连续的，但是指针指向的地址不一定为连续。
	fmt.Printf("r2.lefUp 本身地址=%p r2.rightDown 本身地址=%p\n",
		&r2.leftUp, &r2.rightDown)
	fmt.Printf("r2.lefUp 指向地址=%p r2.rightDown 指向地址=%p\n",
		r2.leftUp, r2.rightDown)



}

func Structdetails0203() {

	var a A
	var b B
	a = A(b) // 可以转换，但是转换结构体的字段名字，类型，个数必须完全相同
	fmt.Println(a, b)

	var human1 Humman
	var human2 Hummm
	human2 = Hummm(human1)

	fmt.Println(human1)
	fmt.Println(human2)

}

func Structdetails04() {

	//1. 创建一个Monster变量
	monster := Monster_n{"Bulls", 500, "Wind"}

	//2.结构体序列化 将mosnter变量序列化为json字串
	// Masshal中使用反射
	jsonStr, err := json.Marshal(monster)

	if err != nil {
		fmt.Println("json error",err)
	}
	fmt.Println("jsonStr", string(jsonStr))

}