package exec

import "strconv"

func Test03() {
	//字符串拼接使用字节缓冲
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}