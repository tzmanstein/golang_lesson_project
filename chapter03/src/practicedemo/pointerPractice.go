package main

import "fmt"

func main() {

	var num1 int = 99
	fmt.Printf("num1 = %v\n", num1)

	var ptr *int = &num1
	*ptr = 100
	fmt.Printf("num1 again = %v\n", num1)
	fmt.Printf("ptr = %v\n", *ptr)


	var n int = 500000
	var i int = 1
	var count int = 0
	for i <= n {
		if n % i == 0 {
			count++
		}
		i++
	}

	fmt.Println("count = ", count)




}
