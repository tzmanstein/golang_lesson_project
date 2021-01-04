package encapsulate

import "fmt"

type account struct {
	accountNo string
	pwd string
	balance float64
}

func NewAccount(accountNo string, pwd string, balance float64) *account {

	if len(accountNo) < 6 || len(accountNo) > 10{
		fmt.Println("账号长度错误")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码长度错误")
		return nil
	}

	if balance < 20.0 {
		fmt.Println("余额错误")
		return nil
	}

	var acc = account{
		accountNo: accountNo,
		pwd: pwd,
		balance: balance,
	}

	return &acc
}

/**
 * 需要加入账户校验
 */
func (account *account) GetAccountNo() string {
	return account.accountNo
}


func (account *account) SetAccountNo(accountNo string) {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度错误")
		return
	}
	account.accountNo = accountNo
}

func (account *account) GetPassowrd() string {
	return account.pwd
}


func (account *account) SetPassword(password string) {
	if len(password) != 6 {
		fmt.Println("密码长度错误")
		return
	}
	account.pwd = password
}

func (account *account) GetBalance() float64 {
	return account.balance
}


func (account *account) SetBalance(balance float64) {
	if balance < 20.0 {
		fmt.Println("余额错误")
		return
	}
	account.balance = balance
}

//存款
func (account *account) SaveMoney(depositAmount float64, pwd string) {
	if pwd != account.pwd {
		fmt.Println("the passowrd you entered, is not correct")
		return
	}

	//检查输入存款余额是否正确
	if depositAmount <= 0 {
		fmt.Println("余额输入金额不正确")
		return
	}

	account.balance += depositAmount
	fmt.Println("存款成功")
}

//取款
func (account *account) WithDraw(withdrawAmount float64, pwd string) {
	if pwd != account.pwd {
		fmt.Println("the passowrd you entered, is not correct")
		return
	}

	//检查输入存款余额是否正确
	if withdrawAmount <= 0 || withdrawAmount > account.balance{
		fmt.Println("余额输入金额不正确")
		return
	}

	account.balance -= withdrawAmount
	fmt.Println("取款成功")
}

//查询余额
func (account *account) Query(pwd string) {
	if pwd != account.pwd {
		fmt.Println("the passowrd you entered, is not correct")
		return
	}

	fmt.Printf("你的账号%v,余额=%v \n", account.accountNo, account.balance)
}