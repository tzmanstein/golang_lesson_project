package main

import (
	"fmt"
	"model/asserts"
	"model/interfaces"
)

func main() {
	interfaces.Run_interfaces()

	fmt.Println()

	interfaces.Run_interfaceDetails()

	fmt.Println()
	asserts.Run_assertDemo()



	fmt.Println("Current run main_interfaces")

}
