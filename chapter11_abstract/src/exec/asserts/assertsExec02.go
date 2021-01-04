package asserts

import (
	"fmt"
)

type Student struct {

}

func TypeJudge(itmes... interface{}) {

	for index, x := range itmes {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", index, x)
		case int, int32, int64 :
			fmt.Printf("第%v个参数是整数类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是字符串类型，值是%v\n", index, x)
		case Student :
			fmt.Printf("第%v个参数是Student类型，值是%v\n", index, x)
		case *Student :
			fmt.Printf("第%v个参数是**Student类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数是 类型不确定，值是%v\n", index, x)

		}
	}
}


func Run_assertsExec02() {
	//TypeJudge()
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "beijing"
	n4 := 300

	var std Student = Student{}
	var pstd *Student = &Student{}

	var std2 = Student{}
	var pstd2 = &Student{}

	std3 := Student{}
	pstd3 := &Student{}

	TypeJudge(n1, n2, n3, name, address, n4, std, pstd, std2, pstd2, std3, pstd3)
}
