package extends

import "fmt"

type A struct {
	Name string
	Age int
}

func (a *A) SayOk() {
	fmt.Println("banded A struct Say Ok", a.Name)
}

func (a *A) hello() {
	fmt.Println("HELLO A", a.Name)
}

func (a *B) hello() {
	fmt.Println("HELLO B", a.Name)
}

type B struct {
	A
	Name string

}

type C struct {
	Name string
	age int
}

type D struct {
	Name string
	age int
}

type E struct {
	C
	D
}

type F struct {
	a A
}

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type Television struct {
	Goods
	Brand
}

type Television2 struct {
	*Goods
	*Brand
}


// 演示以下匿名字段是基本数据类型
type Monster struct {
	Name string
	Age int
}

type Basetype struct {
	Monster
	int //嵌入匿名字段是基本数据类型, 不能重复
	n int //如果需要多个int字段需要指定名字
}


func Run_extendsdetails() {
	var b B

	b.Name = "tom"
	b.A.Name = "jerry scott"
	b.Age = 22
	b.SayOk()
	b.hello()

	//如果E中没有Name，而继承的C，D有字段Name那么必须明确指定匿名结构体的名字来区分
	//以上规则对方法同样适用。
	var e E
	e.C.Name = "jessy"
	fmt.Println(e)

	// 如果结构体中存在一个有名结构体时，使用时，字段名必须明确指定
	var f F
	f.a.Name = "harry"
	fmt.Println(f)

	//直接指定匿名结构体的值
	tv1 := Television{Goods{"dianshiji0002", 3000.99}, Brand{"sharp", "qingdao"}}

	tv2 := Television{
		Goods{
			Name:  "电视机0002",
			Price: 5000.99,
		},
		Brand{
			Name: "Hair",
			Address: "qingdao",
		},
	}


	tv3 := Television2{&Goods{"aaaa", 1999.99},&Brand{"bbbb", "ccccc"} }
	tv4 := Television2{
		&Goods{
			Name: "Tv",
			Price: 4000.99,
		},
		&Brand{
			Name: "Kangjia",
			Address: "Shichuan",
		},
	}

	fmt.Println(tv1)
	fmt.Println(tv2)

	fmt.Println("tv3",*tv3.Goods,*tv3.Brand)
	fmt.Println("tv4",*tv4.Goods,*tv4.Brand)

	// 演示以下匿名字段是基本数据类型
	var bt Basetype
	bt.Name = "Dragon"
	bt.Age = 1000
	bt.int = 20
	bt.n = 40

	fmt.Println("basetype", bt)
}
