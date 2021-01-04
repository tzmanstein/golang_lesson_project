package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var b = false
	fmt.Println("b=", b)
	fmt.Println("b space is = ", unsafe.Sizeof(b))

}
