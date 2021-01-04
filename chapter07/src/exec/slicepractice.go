package exec

import "fmt"

func test(slice []int) {
	slice[0] = 100
}

func Pratice01() {
	var slice = []int{1,2,3,4}
	fmt.Println(slice)
	test(slice)
	fmt.Println(slice)
}

// 斐波那契数列
func Fbn(n int) ([] uint64) {

	// 斐波那契数列特点0 = 1， i= 1 for循环特殊起点

	fbnSlice := make([]uint64, n)

	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i -1] + fbnSlice[i - 2]
	}

	return fbnSlice
}