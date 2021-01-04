package stdy

import "fmt"

func SliceDemo01() {

	// slice基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}


	//声明定义一个slice
	//1. slice切片名
	//2. intArr[1:3] 表示slice医用到intArr这个数组的第2个元素
	//3. 从数组下标1开始引用到下标3，但不包含下标3
	slice := intArr[:]

	fmt.Println("intArr=",intArr)
	fmt.Println("slice=", slice)
	fmt.Println("silce length = ", len(slice))
	fmt.Println("slice capacity = ", cap(slice))

	fmt.Printf("slice[0] Addr=%p\n", &slice[0])
	fmt.Printf("addr[1] Addr=%p\n", &intArr[1])
}

func CreateSlice() {
	//var slice []int = make([]int, 3, 10)
	// type, size, capbility
	//var slice = make([]int, 3, 10)

	// 对于slice必须make后才能使用
	var slice = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Println("slice len=", len(slice))
	fmt.Println("slice cap=", cap(slice))


	var slicestr = []string{"tom", "jack", "mary"}

	fmt.Println(slicestr)

	for i,v := range slicestr {
		fmt.Printf("i=%v, v=%v\n",i, v)
	}


}

func AppendSlice() {
	var slice3 = []int{100, 200, 300}
	fmt.Println("slice3", slice3)
	// 类型一致
	slice3 = append(slice3, 400, 500)
	fmt.Println(slice3)

	//slice中追加slice
	slice3 = append(slice3,slice3...)
	fmt.Println(slice3)

	var slice4 = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4", slice4)
	fmt.Println("slice5", slice5)

}

func StringToSlice() {
	//string底层是一个byte数组，一次string可以进行切片处理
	str := "hello@atguigu"

	//使用slice获得atguigu
	slice := str[6:]

	fmt.Println("slice=", slice)


	//string是不可变的，也就是说不能通过str[0] = 'z'的方式来修改字符串
	//str【0】= 'z' 编译错误，string不可变

	//如果需要修改字符串，可以先将string -> []byte / 或者 []rune -> 修改 -> 重写转换string
	// "hello@atguigu" => "zello@atguigu"
	//slice2 := str[:]
	//slice2[0] = 'z'
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)


	//兼容全角字符
	arr2 := []rune(str)
	arr2[0] = '东'
	str = string(arr2)
	fmt.Println("str=", str)




}