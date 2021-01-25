package channeldetails

import (
	"fmt"
	"time"
)

func Run_chdetails01() {

	//demo02()

	demo03()
}

func chDetail01() {
	//默认情况下，双向读写
	var chan1 chan int
	chan1 = make(chan int, 3)
	chan1<- 10

	//2.声明只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2<- 20
	//num := <-chan2 //错误
	fmt.Println("chan2=", chan2)

	//3.声明为只读
	var chan3 <-chan int
	num2 := <-chan3
	//chan3<- 30  //错误
	fmt.Println("chan3=", num2)

	//只读只写管道，通常作为管道参数标示使用。参数管道仍然是读写双管道,
	//只不过在特定场合限定管道使用范围

}

func demo02() {
	//使用select可以解决从管道去数据的阻塞问题

	//1. 定义一个10 int channel
	var intChan chan int
	intChan = make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan<- i
	}

	//2.定义一个5 string channel
	var strChan chan string
	strChan = make(chan string, 5)

	for i := 0; i < 5; i++ {
		strChan<- "hello" + fmt.Sprintf("%d", i)
	}

	//传统方法在遍历管道时，如果不关闭会阻塞而导致deadlock

	//实际开发中，实际开发中不能确切的确定关闭管道的时机，
	//可以十一哦那个select方式可以解决

	label:
	for {
		select {
		//注意：如果intChan一直没有关闭，不会一直阻塞而deadlock
		//，会自动到下一个case匹配
			case v := <-intChan :
				fmt.Printf("从intChan读取的数据%d\n", v)
			case v := <- strChan :
				fmt.Printf("从strChan中去读数据%v\n", v)
			default:
				fmt.Printf("无法取得数据，退出\n")
				break label
				//return
		}
	}
}


func demo03() {
	go sayHello()
	go doMap()

	for i := 0; i < 10; i++ {
		fmt.Println("main ok", i)
		time.Sleep(time.Second)
	}

}

func doMap() {

	//使用defer，recover通过匿名函数来捕获panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("get err=",err)
		}
	}()

	var myMap map[int]string
	//myMap = make(map[int]string, 3)
	myMap[0] = "hehehe"
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}
