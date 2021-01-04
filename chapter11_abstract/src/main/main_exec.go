package main

import (
	"exec/encapsulate"
	"fmt"
)

func main() {
	//创建一个银行账户
	account := encapsulate.NewAccount("jzh11111", "999999", 60)

	account.SetAccountNo("jzh11112")

	if account != nil {
		fmt.Println("账户创建成功",*account)
	} else {
		fmt.Println("账户创建失败")
	}

}
