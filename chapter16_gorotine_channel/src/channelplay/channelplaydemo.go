package channelplay

import "fmt"

type Cat struct {
	Name string
	Age int
}

func play01() {
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{
		Name: "Kitty",
		Age:  3,}
	allChan <- cat

	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)
	confirmCat := newCat.(Cat)//类型断言
	fmt.Printf("newCat name= %v\n", confirmCat.Name)
}

//play2
func play2_initnumChan(chanSize int, numChan chan int) {
	if chanSize < 1 {
		return
	}

	for i := 0; i < chanSize; i++ {
		numChan<- i
		fmt.Printf("numChan Entry = %v\n", i)
	}

	close(numChan)
}

func play2_readDataFrmnumChan(numChan chan int, resChan chan int, exitChan chan bool) {

	for {
		val, ok := <-numChan

		if !ok{
			fmt.Println("fail to read")
			close(resChan)
			break
		}
		resChan<- cal(val)
		fmt.Printf("resChan Entry = %v\n", val)
	}

	//加入8个协程，协程嵌套使用 vedio:274

	for v := range resChan {
		fmt.Printf("reschan value = %v\n", v)
	}

	exitChan <- true
	close(exitChan)
}

func cal(calNum int) int {
	res := 0
	for i := 1; i <= calNum; i++{
		res += i
	}
	//fmt.Printf("cal res=%v\n", res)
	return res
}

func play2() {
	chanSize := 2000
	numChan := make(chan int, chanSize)
	resChan := make(chan int, chanSize)
	exitChan := make(chan bool, 1)

	go play2_initnumChan(chanSize,numChan)
	go play2_readDataFrmnumChan(numChan, resChan, exitChan)
	//go play2_readDataFrmnumChan(numChan, resChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}



func Run_chanelPlay() {
	play2()

}