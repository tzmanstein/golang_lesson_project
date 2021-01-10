package view

import (
	"fmt"
	"service"
)

type customerView struct {
	//定义必要字段
	key string // 接收用户输入...
	loop bool  //表示是否循环显示
	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示所有客户信息
func (this *customerView) list() {
	//获取当前所有客户信息（切片中）
	customers := this.customerService.List()
	//fmt.Println(customers)
	//显示
	fmt.Println("----------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _,val := range customers {
		info := val.GetInfo()
		fmt.Println(info)
	}
	fmt.Println("\n--------------------------客户列表完成--------------------------\n")

}


//添加客户信息
func (this *customerView) add() {

	//显示
	fmt.Println("----------------------------添加客户---------------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮件：")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//注意：id号，没有用户输入，id为唯一
	//customer := model.AddCustomer(name, gender, age, phone, email)
	//
	// if this.customerService.Add(customer) {
	//	 fmt.Println("----------------------------添加完成---------------------------")
	// } else {
	//	 fmt.Println("----------------------------添加失败---------------------------")
	// }

	if this.customerService.Add2(name, gender, age, phone, email) {
		fmt.Println("----------------------------添加完成---------------------------")
	} else {
		fmt.Println("----------------------------添加失败---------------------------")
	}
}

func (this *customerView) update() {
	fmt.Println("----------------------------修改客户----------------------------")
	fmt.Println("请选择待修改客户编号(-1退出)： ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否修改(y/n): ")
	choice := ""
	udtLoop := true

	for udtLoop {
		fmt.Scanln(&choice)
		switch choice {
		case "y", "Y":
			udtLoop = false
			if this.customerService.Update(id) {
				fmt.Println("----------------------------修改完成----------------------------")
			} else {
				fmt.Println("------------------------修改失败，该id不存在----------------------")
			}
		case "n", "N":
			udtLoop = false
		default:
			fmt.Println("----------------------------请输入Y(y)或N(n)----------------------------")

		}
	}
}

//得到用户输入id，删除id对应的客户
func (this *customerView) delete() {
	fmt.Println("----------------------------删除客户----------------------------")
	fmt.Println("请选择待删除客户编号(-1退出)： ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(y/n): ")
	choice := ""
	delLoop := true

	for delLoop {
		fmt.Scanln(&choice)
		//if choice == "y" || choice == "Y" {
		//	//调用Service的Delete方法进行删除
		//	delLoop = false
		//	if this.customerService.DeleteById(id) {
		//		fmt.Println("----------------------------删除完成----------------------------")
		//	} else {
		//		fmt.Println("------------------------删除失败，该id不存在----------------------")
		//	}
		//} else if choice == "n" || choice == "N" {
		//	delLoop = false
		//} else {
		//	fmt.Println("----------------------------请输入Y(y)或N(n)----------------------------")
		//}

		switch choice {
		case "y", "Y":
			delLoop = false
			if this.customerService.DeleteById(id) {
				fmt.Println("----------------------------删除完成----------------------------")
			} else {
				fmt.Println("------------------------删除失败，该id不存在----------------------")
			}
		case "n", "N":
			delLoop = false
		default:
			fmt.Println("----------------------------请输入Y(y)或N(n)----------------------------")

		}
	}
}

//系统退出函数
func (this *customerView) exit() {
	fmt.Println("确认是否退出输入(Y/N): ")

	extLoop := true
	for extLoop {
		fmt.Scanln(&this.key)
		switch this.key {
		case "y", "Y":
			extLoop = false
			this.loop = false
		case "n", "N":
			extLoop = false
		default:
			fmt.Println("输入有误，请重新输入")

		}
	}
}


//显示主菜点
func (this *customerView) mainMenu() {

	for this.loop{
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1.添 加 客 户")
		fmt.Println("                 2.修 改 客 户")
		fmt.Println("                 3.删 除 客 户")
		fmt.Println("                 4.客 户 列 表")
		fmt.Println("                 5.退      出")
		fmt.Println("                                              ")
		fmt.Println("                 请选择（1-5）：")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			fmt.Println("添加客户")
			this.add()
		case "2":
			fmt.Println("修改客户")
			this.update()
		case "3":
			fmt.Println("删除客户")
			this.delete()
		case "4":
			fmt.Println("客户列表")
			this.list()
		case "5":
			//this.loop = false
			this.exit()
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
	}

	fmt.Println("你退出了客户关系管理系统...")
}

func Run_customerView() {
	fmt.Println("customer view")
	//创建customerView实例，并进入主菜单

	customerview := customerView{
		key : "",
		loop: true,
	}
	//完成对customerView结构体的customerService字段初始化
	customerview.customerService = service.NewCustomerService()

	customerview.mainMenu()

}
