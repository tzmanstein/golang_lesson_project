package main

import (
	"fmt"
	"lesson"
)

func main() {
	var pmd lesson.Person_method
	pmd.Name = "thomas"
	pmd.Test()
	fmt.Println("main() p.Name=", pmd.Name)

	pmd.Speak()

	pmd.Calculate()
	pmd.Calculate2(10)

	n1 := 10
	n2 := 20

	resutl := pmd.Getsum(n1, n2)
	fmt.Println(resutl)


}
