package routineChannelSample03

import (
	"fmt"
	"time"
)

//初始化使用匿名函数实现
func putNum(intChan chan int) {
	for i:= 1; i <= 80000; i++{
		intChan<- i
	}
	//关闭channel
	close(intChan)
}

//channel取出数据
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//num := 0
	//flag := false

	//var num int
	var flag bool

	for {
		//time.Sleep(time.Millisecond*10)
		num, ok := <- intChan
		if !ok {
			break
		}

		flag = true //预设素数初期值为真，
		//判断素数
		for i := 2; i < num; i++ {
			if num % i == 0 {
				//非素数
				flag = false
				break
			}
		}

		if flag {
			primeChan<- num
		}
	}
	fmt.Println("一个primeNum协程数据完成，退出")
	exitChan<- true
}

func sample03() {
	//任务管道
	intChan := make(chan int, 1000)
	//结果管道
	primeChan := make(chan int, 20000)
	//退出管道，基本上与启动协程数量相同
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()
	//开启一goroutine，向intChan放入1-8000个数字
	go putNum(intChan)

	//开启四个协程，从intChan中取出数据，并判断是否为素数,如果真放入primeChan中
	for i := 0; i < 4; i++{
		go primeNum(intChan, primeChan, exitChan)
	}

	//主线程处理，关闭primeChan以及exitChan
	go func() {
		for i:= 0 ; i < 4; i++ {
			<-exitChan
		}
		//当从exitChan中取出4个结果，可以关闭primeChan

		end := time.Now().Unix()
		fmt.Println("time Cost=", end -start)

		close(primeChan)
	}()

	for {
		_,ok := <- primeChan
		if !ok {
			break
		}
		//fmt.Printf("primeNum=%v\n", res)

	}

	fmt.Println("主线程退出!!")

}


func singleTest() {
	startTime := time.Now().Unix()

	flag := true

	for num := 1; num <= 80000; num++ {
		for i := 2; i <= num; i++ {
			if num % i == 0{
				flag = false
				break
			}
		}

		if flag{

		}
	}

	endT := time.Now().Unix()

	fmt.Println("singel = ", endT - startTime)
}

func Run_sample03() {
	sample03()
	//singleTest()
}

