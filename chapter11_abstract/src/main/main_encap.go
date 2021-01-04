package main

import (
	"fmt"
	"model/encapsulate"
)

func main() {

	p := encapsulate.NewPerson("smith")
	p.SetAge(28)
	p.SetSalary(5000.0)
	fmt.Println(p)

	fmt.Println(p.Name," age= ", p.GetAge(), "salary= ", p.GetSalary())



}
