package main

import (
	"fmt"
	"funcdemo01/exec"
	"time"
)

func main() {
	start := time.Now().Unix()
	exec.Test03()
	end := time.Now().Unix()
	fmt.Printf("执行test03锁耗时间为%v秒\n", end - start)

}
