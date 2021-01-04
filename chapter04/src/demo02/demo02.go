package main

import "fmt"

func test() bool {
	fmt.Println("test.....")
	return true
}

func Tickets(peopleInLine []int) string {

	result := "NO"
	counter25 := 0
	counter50 := 0
	counter100 := 0

	if peopleInLine[0] != 25 {
		return result
	}

	for i:=0; i < len(peopleInLine); i++ {

		switch peopleInLine[i] {
		case 25:
			counter25++
			break
		case 50:
			counter25--
			if counter25 >= 0 {
				counter50++
			} else {
				return result
			}
			break
		case 100:
			if counter25 >= 3 && counter50 == 0{
				counter25 = counter25 - 3
			} else {
				counter25--
				counter50--
				if counter25 >= 0 && counter50 >= 0 {
					counter100++
				} else {
					return  result
				}
			}
			break
		}

	}
	result = "YES"

	return result
}

func Tickets2(peopleInLine []int) string {
	t, f := 0, 0
	for _, bill := range peopleInLine {
		switch bill {
		case 25:
			t++
		case 50:
			f++
			t--
		case 100:
			if f > 0 {
				f--
				t--
			} else {
				t -= 3
			}
		}
		if t < 0 || f < 0 { return "NO" }
	}
	return "YES"
}

func SpinWords(str string) string {


	byteStr := []byte(str)
	//fmt.Println(byteStr[0])

	fmt.Printf("char = %c\n", byteStr[0])

	//先按照空格进行分割，
	// 判断单个单词长度
	// 进行倒叙




	return "pizza"
}// SpinWords




func main() {
	var i int = 8
	var a int

	//a = i++
	//a = i--
	i++
	//++i
	//--i

	a = i
	fmt.Println("a = ", a)

	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)

	fmt.Printf("%v 对应的摄氏温度=%v \n", huashi, sheshi)

	// demo04

	var age int = 40

	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}

	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}

	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	if !(age > 30) {
		fmt.Println("ok5")
	}

	if age > 30 && test() {
		fmt.Println("duanlu ce shi")
	}

	if age > 50 || test() {
		fmt.Println("duanlu ce shi2")
	}

	fmt.Println("result1= ", Tickets([]int{25,25,25,100}))
	//Tickets([]int{25,25,25,50})
	fmt.Println("result2= ", Tickets2([]int{25,25,25,25,25,100,50,50,25}))

	SpinWords("adfadfasf")


}
