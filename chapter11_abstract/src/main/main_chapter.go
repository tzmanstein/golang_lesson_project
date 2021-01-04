package main

import (
	"abstract"
)

func main() {

	account := abstract.Account{
		AccountNo: "gs1111111",
		Pwd: "666666",
		Balance: 100.0,
	}

	//var selectMode byte
	enterPwd := "666666"

	//for {
	//	//用户通过控制台输入
	//	fmt.Println("选择业务种类，1，查询，2，存款，3取款，5退出")
	//	fmt.Scanln(&selectMode)
	//	switch selectMode {
	//	case 1:
	//		account.Query(enterPwd)
	//	case 2:
	//		account.SaveMoney(300, enterPwd)
	//	case 3:
	//		account.WithDraw(100,enterPwd)
	//	case 4:
	//		break
	//	default:
	//		fmt.Println("输入错误，再次输入")
	//	}
	//}



	account.Query(enterPwd)

	account.SaveMoney(200,enterPwd)
	account.Query(enterPwd)

	account.WithDraw(150.0, enterPwd)
	account.Query(enterPwd)

}
