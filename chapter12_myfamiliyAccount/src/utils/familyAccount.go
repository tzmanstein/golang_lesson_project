package utils

import "fmt"

type FamilyAccount struct {
	//接收用户收入
	key string
	//处理标示
	loop bool
	//账户余额
	balance float64
	//每次收支金额
	money float64
	//每次收支说明
	note string
	//收支详情
	//details := "收支\t账户金额\t收支金额\t说    明"
	details string

	initFlag bool
}

func NewFamilyAccount() *FamilyAccount {

	return &FamilyAccount{
		key: "",
		loop: true,
		balance: 10000.0,
		money: 0.0,
		note: "",
		initFlag: false,
		details: "收支\t账户金额\t收支金额\t说    明",
	}
}

func (this *FamilyAccount) DisplaymainMenu() {

	for this.loop {
		fmt.Println("\n-------------------家庭收支记账软件-------------------")
		fmt.Println("                    1 收支明细")
		fmt.Println("                    2 登记收入")
		fmt.Println("                    3 登记支出")
		fmt.Println("                    4 退出")
		fmt.Print("请选择(1-4): ")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()

		case "2":
			this.income()

		case "3":
			this.payment()
		case "4":
			this.exit()
		default:
			fmt.Println("请选择正确选项...")
		}
	}
	fmt.Println("退出家庭记账软件使用")


}

func (this *FamilyAccount) showDetails() {
	if this.initFlag {
		fmt.Println("-------------------当前收支明细记录-------------------")
		fmt.Println(this.details)
	} else {
		fmt.Println("未进行任何收支...")
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money //修改账户余额
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.initFlag = true

	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
}

func (this *FamilyAccount) payment() {
	fmt.Println("登记支出...")
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("钱不够..")
		//return
	}
	this.balance -= this.money //修改账户余额
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.initFlag = true

	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance,this.money,this.note)
}

func (this *FamilyAccount) exit() {

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
	this.loop = false
}
}