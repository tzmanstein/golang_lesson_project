package channelclose

import "fmt"

func Run_channelclose() {
	scanchannel()
}

func channelbase() {
	intChan := make(chan int, 3)

	intChan<- 100
	intChan<- 200
	close(intChan)
	//channel close后不能继续写入channel
	//intChan<- 300

	n1 := <-intChan
	fmt.Println("asdfafa", n1)
}

func scanchannel() {
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2<- i
	}
	close(intChan2)
	//遍历，管道本身在取出时，为动态变化，不能使用普通的for循环。
	//遍历管道使用for each

	for v := range intChan2 {
		fmt.Printf("value=%v\n", v)
	}
}
