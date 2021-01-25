package channel

import "fmt"

func Run_chanelDemo() {
	//演示管道使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2.看看intChan是什么,确认结果引用类型
	fmt.Printf("LINE1 intChan的值=%v, intChan本身得治=%p\n", intChan, &intChan)

	//3. channel中写入数据

	intChan<- 10
	num := 211
	intChan<- num

	//注意： 当给管道写入数据时，不能超过其容量
	intChan<- 50
	<-intChan
	intChan<- 98


	//4.看看管道的长度和cap（容量），管道不能自动增长而map，slice可以
	fmt.Printf("LINE2 channel len=%v cap=%v \n", len(intChan), cap(intChan))

	//5. 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	//6.在没有使用goroutine的情况下，如果管道数据已经全部取出，再进行读取会报deadlock
	num3 := <-intChan
	num4 := <-intChan
	intChan<- 9527

	num5 := <-intChan 	//(报错)
	fmt.Println("num3=", num3)
	fmt.Println("num4=", num4)
	fmt.Println("num5=", num5)

}
