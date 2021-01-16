package argsPara

import (
	"flag"
	"fmt"
	"os"
)

func Run_args() {
	//args_demo1()
	flagpkg_demo1()
}


func args_demo1() {
	fmt.Println("命令行参数有", len(os.Args))
	for ind, val := range os.Args {
		fmt.Printf("args[%v]=%v\n", ind, val)
	}
}

func flagpkg_demo1() {

	//定义变量，接收命令行参数值。
	user := ""
	pwd := ""
	host := ""
	port := 0

	//&use，用来接收用户命令行中输入的-u后的参数值
	//"u",就是 -u 指定参数
	//"用户名，默认为空"说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	//重要：转换，必须调用
	flag.Parse()

	//输入结果
	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v\n", user, pwd, host, port)
}