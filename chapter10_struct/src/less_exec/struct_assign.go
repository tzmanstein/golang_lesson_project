package less_exec

import "fmt"

type Stu struct {
	Name string
	Age int
}



func Run_sturctassign() {

	//方式1
	//在创建结构体变量时，直接指定字段值
	var stu1 = Stu{"xiaoming", 19}
	stu2 := Stu{"xiaohong", 20}
	fmt.Println(stu1)
	fmt.Println(stu2)
	//创建结构体变量时，吧字段名和字段值写在一起，此种写法不依赖字段的定义顺序
	var stu3 = Stu{
		Name: "jack",
		Age: 22,
	}
	stu4 := Stu{Age:30, Name:"Marray",}
	fmt.Println(stu3)
	fmt.Println(stu4)

	//方式2 返回结构体的指针类型
	var stu5 = &Stu{"xiaowang", 29} //地址->结构体数据
	stu6 := &Stu{"xiaowang~", 39}

	var stu7 = &Stu{
		Name: "xiaoli",
		Age: 49,
	}
	stu8 := &Stu{
		Age:59,
		Name: "xiaoli~",
	}
	fmt.Println(*stu5, *stu6, stu7,stu8)

}


