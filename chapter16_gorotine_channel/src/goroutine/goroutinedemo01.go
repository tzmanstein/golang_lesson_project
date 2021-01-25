package goroutine

import (
	"fmt"
	"strconv"
	"time"
)

func Run_gorouDemo01() {

	go goroutine_01() //开启了一个goroutine

	for i := 1; i <= 10; i++ {
		fmt.Println("hello,golang!", strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}

func goroutine_01() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello,world!", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}