package main

import (
	"fmt"
)

func main() {

	//接收用户收入
	key := ""
	//处理标示
	loop := true
	//账户余额
	balance := 10000.0
	//每次收支金额
	money := 0.0
	//每次收支说明
	note := ""
	//收支详情
	details := "收支\t账户金额\t收支金额\t说    明"

	initFlag := false

	for loop {
		fmt.Println("\n-------------------家庭收支记账软件-------------------")
		fmt.Println("                    1 收支明细")
		fmt.Println("                    2 登记收入")
		fmt.Println("                    3 登记支出")
		fmt.Println("                    4 退出")
		fmt.Print("请选择(1-4): ")

		fmt.Scanln(&key)

		switch key {
			case "1":
				if initFlag {
					fmt.Println("-------------------当前收支明细记录-------------------")
					fmt.Println(details)
				} else {
					fmt.Println("未进行任何收支...")
				}

			case "2":
				fmt.Println("本次收入金额：")
				fmt.Scanln(&money)
				balance += money //修改账户余额
				fmt.Println("本次收入说明：")
				fmt.Scanln(&note)
				initFlag = true

				details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance,money,note)
			case "3":
				fmt.Println("登记支出...")
				fmt.Println("本次支出金额：")
				fmt.Scanln(&money)
				if money > balance {
					fmt.Println("钱不够..")
					break
				}
				balance -= money //修改账户余额
				fmt.Println("本次支出说明：")
				fmt.Scanln(&note)
				initFlag = true

				details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance,money,note)
			case "4":
				fmt.Println("你确定要退出吗？ y/n")
				choice := ""
				for {
					fmt.Scanln(&choice)
					if choice == "y" || choice == "n" {
						break
					}
					fmt.Println("你的输入有误，请从新输入 y/n")
				}
				if choice == "y" {
					loop = false
				}
			default:
				fmt.Println("请选择正确选项...")
		}
	}
	fmt.Println("退出家庭记账软件使用")
}
