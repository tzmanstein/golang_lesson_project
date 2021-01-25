package goruntime

import (
	"fmt"
	"runtime"
)

func Run_cpudemo() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置使用多少个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
