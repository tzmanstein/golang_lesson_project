package goroutine_channel

import (
	"fmt"
	"sync"
	"time"
)

/**
需求：计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中
	最后显示出来。要求使用goroutine完成

思路：
1.编写一个函数，来计算各个数的阶乘，并放入到map中
2.启动多个goroutine，统计将结果放入到map中
3. 需要一个全局map变量
 */

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁
	lock sync.Mutex
)

func subtask(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()
	//将结果放入map中
	myMap[n] = res
	lock.Unlock()
}

func routineByLock() {
	// 将这里开启多个goroutine完成这个任务
	for i := 1; i <= 20; i++ {
		go subtask(i)
	}

	time.Sleep(time.Second*3)

	lock.Lock()
	for idx, val := range myMap  {
		fmt.Printf("map[%d] = %d\n", idx, val)
	}
	lock.Unlock()
}

// goroutine&channel Sample2
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		//数据压入管道
		intChan<- i
		fmt.Printf("writeData=%v\n ", i)
		//time.Sleep(time.Second)
	}
	//关闭
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <- intChan
		if !ok {
			//fmt.Println("No data?")
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("readdata = %v\n", v)
	}
	//读取我那数据后，任务完成
	exitChan<- true
	close(exitChan)

}

func routineByChannel() {
	//创建2个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	//go writeData(intChan)
	//go readData(intChan,exitChan)
	//
	//time.Sleep(time.`Second*10)

	go writeData(intChan)
	go readData(intChan,exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

func Run_RoutineChannel() {
	//routineByLock()
	routineByChannel()
}



