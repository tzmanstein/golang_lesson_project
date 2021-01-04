package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20

	a = a + b
	b = a -b
	a = a - b
	fmt.Printf("a=%v, b=%v\n", a ,b)

	// Including exec03
	n1 := 10
	n2 := 40

	var max int

	if n1 > n2 {
		max = n1
	 } else {
	 	max = n2
	}

	var n3 = 45

	if n3 > max {
		max = n3
	}


}
