package abstract

import "fmt"

//此抽象不是抽象类，而是一种将实际业务抽象成代码的一种思维。
type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

//存款
func (account *Account) SaveMoney(depositAmount float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("the passowrd you entered, is not correct")
		return
	}

	//检查输入存款余额是否正确
	if depositAmount <= 0 {
		fmt.Println("余额输入金额不正确")
		return
	}

	account.Balance += depositAmount
	fmt.Println("存款成功")
}

//取款
func (account *Account) WithDraw(withdrawAmount float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("the passowrd you entered, is not correct")
		return
	}

	//检查输入存款余额是否正确
	if withdrawAmount <= 0 || withdrawAmount > account.Balance{
		fmt.Println("余额输入金额不正确")
		return
	}

	account.Balance -= withdrawAmount
	fmt.Println("取款成功")
}

//查询余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("the passowrd you entered, is not correct")
		return
	}

	fmt.Printf("你的账号%v,余额=%v \n", account.AccountNo, account.Balance)
}