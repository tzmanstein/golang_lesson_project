package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello南" //go编码统一为utf-8 字符数字ASCII字符均为1个字节。汉字占用3个字节 len按字节返回

	fmt.Println("str len =",len(str))

	str2 := "hello辽阳"
	strrune := []rune(str2)
	for i := 0; i < len(strrune); i++{
		fmt.Printf("%c\n", strrune[i])
	}

	// string to int
	n, err := strconv.Atoi("123")  //可用于数据校验
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(n)
	}

	//4. int to string
	str = strconv.Itoa(12345)
	fmt.Printf("str = %v, str=%T\n", str, str)

	//5. 字符串转byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)
	fmt.Printf("bytes=%c\n", bytes)

	//6. []byte转字符串
	str = string([]byte{97,98,99})
	fmt.Println("str = ", str)

	//7. 10进制转2，8，16
	str = strconv.FormatInt(123,2)
	fmt.Println("str= ", str)

	//8. 查找子串是否在指定的字符串中
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b=%v\n", b)

	//9. 统计一个字符串有几个指定的子串
	numCount := strings.Count("abbbddddd", "b")
	fmt.Printf("numCount= %v\n", numCount)

	//10，字符串比较
	b = strings.EqualFold("abc", "ABC") //不区分大小写
	fmt.Println("b=",b)

	b = "abc" == "ABC"
	fmt.Println("b=",b) //区分大小写

	//11.返回字串第一次出现的index值
	index := strings.Index("NLT_abcabc", "abc")
	fmt.Printf("index=%v\n", index)

	//12.返回字符串子串最后一次出现的index.如果没有匹配返回-1,
	index = strings.LastIndex("go golang", "go") //3
	fmt.Printf("index=%v\n", index)

	//13. 将指定的子串替换成另外一个子串: strings.Replace("go go hello", "go", "go语言", n)
	//n可以指定希望替换的个数。如果n=-1表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "go语言", -1)
	fmt.Printf("str=%v\n", str)

	//14. 按照指定指定某个字符，为分隔符，将一个字符串拆分成字符串数组
	//strings.Split("hello,world,ok", ",")
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n",i ,strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	//15. 将字符串进行大小写转换. string.ToLower(), strings.ToUpper("Go")
	str = "golang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str ToLower=%v\n", str)
	str = strings.ToUpper(str)
	fmt.Printf("str ToUpper=%v\n", str)

	//16. 将字符串左右两边的空格去掉：
	str = strings.TrimSpace(" tn a lone gopher ntrn    ")
	fmt.Printf("str= %q\n", str)

	//17. 将字符串两边指定的字符串去掉
	str = strings.Trim("! he!llo! ", " !")
	fmt.Printf("str = %q\n", str)

	//strings.TimeLeft&TrimRigthht

	//20.判断字符串是否以指定字符串开头
	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b=%v\n", b)






}
