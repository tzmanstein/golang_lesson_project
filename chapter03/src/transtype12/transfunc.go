package main
import (
	"fmt"
	_ "fmt"	// 如此包未被使用，按此表示式进行忽略
)

func main() {

	var i32 int32 = 100
	// covert int to float
	var f32 float32 = float32(i32)

	var i8 int8 = int8(i32)


	fmt.Println(i32, f32)
	fmt.Printf("i32=%v, f32=%v, i8=%v\n", i32, f32, i8)

	// 课题低精度到高精度时，缺位问题。
	//被转换的是变量存储数据拷贝，变量本身保持不变。
	fmt.Printf("i32 type is %T\n", i32)

	// 大范围数据类型转换小范围类型的溢出处理
	var num1 int64 = 999999
	var num2 int8 = int8(num1)

	fmt.Println(num2)

	var n1 int32 = 12
	var n2 int64
	var n3 int8

	n2 = int64(n1) + 20
	n3 = int8(n1) + 20

	fmt.Println("=======this is a practice as below=======")
	fmt.Println("n2=",n2)
	fmt.Println("n3=",n3)

	var n33 int8
	//var n44 int8
	var n44 int32

	n33 = int8(n1) + 127	//  编译可通过，执行时按溢出处理
	//n44 = int8(n1) + 128
	n44 = n1 + 128

	fmt.Println("n33=", n33)
	fmt.Println("n44=", n44)




}
