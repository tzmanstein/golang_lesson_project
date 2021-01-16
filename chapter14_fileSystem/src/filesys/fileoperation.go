package filesys

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func fileaccess() {
	//打开文件
	//概念说明：file叫法
	//1. file 叫file对象
	//2. file 叫file指针
	//3. file 叫file文件句柄
	file, err := os.Open("/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/test")

	if  err != nil {
		fmt.Println("opne file err=",err)

	}
	fmt.Printf("file=%v\n",file)
	//Close File

	err = file.Close()
	if err != nil {
		fmt.Println("close file err:",err)
	}
}


func fileReadingWithBuf() {

	file, err := os.Open("/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/test")

	if  err != nil {
		fmt.Println("open file err=",err)
	}

	//函数退出时，要及时关闭file
	defer file.Close()

	//创建一个Reader
	/*
		const(
		defaultBufSize = 4096 //默认缓冲区4096
	)
	 */
	reader := bufio.NewReader(file)

	//循环读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {//代表文件结尾
			break
		}
		//文件内容输出
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}

//无缓存载入文件
func fileReadingNobuf() {
	file := "/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/test"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v",err)
	}
	//把读取到的内容显示到终端
	fmt.Printf("%v", string(content))
	//没有显式的Open文件，因此也不需要显式的Close文件
	//因为，文件Open和Close被封装到ReadFile函数内部。
}

//写文件
func fileWriter() {
	//创建一个新文件，写入内容 5句hello, Gardon
	// 1.打开文件 abc.txt
	filePath := "/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/writerTarget"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	str := "hello,Gardon\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入缓存的，所以需要调用Flush方法，将缓存的数据
	//写入到文件中，否则文件中会没有数据
	writer.Flush()
}


func fileWrite_exe02_overwrite() {
	//创建一个新文件，写入内容 5句hello, Gardon
	// 1.打开既存文件 test.txt
	filePath := "/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/test"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)

	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	str := "你好,硅谷!\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入缓存的，所以需要调用Flush方法，将缓存的数据
	//写入到文件中，否则文件中会没有数据
	writer.Flush()

}


func fileWrite_exe03_append() {
	//创建一个新文件，写入内容 5句hello, Gardon
	// 1.打开既存文件 test.txt
	filePath := "/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/test"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	str := "ABC!ENGLISH!\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入缓存的，所以需要调用Flush方法，将缓存的数据
	//写入到文件中，否则文件中会没有数据
	writer.Flush()
}


func fileWrite_exe04_readappend() {
	//创建一个新文件，写入内容 5句hello, Gardon
	// 1.打开既存文件 test.txt
	filePath := "/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/test"
	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	//先读取原先内容并显示在终端

	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//显示到终端
		fmt.Print(str)
	}

	str := "Hello,北京!\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入缓存的，所以需要调用Flush方法，将缓存的数据
	//写入到文件中，否则文件中会没有数据
	writer.Flush()
}


func fileWrite_exe05_tranferfile() {
	//将文件abc文件导入到kkk中
	// 1.首先将abc文件内容读取到内存
	// 2.将读取文件内容写入到kkk
	//filePath := "/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/test"
	//file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)

	file1Path := "/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/abc"
	file2Path := "/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/kkk"

	//读出的content为一个byte slice类型内容
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		//文件读取错误
		fmt.Println("read file err=%v",err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		//文件读取错误
		fmt.Println("write file err=%v",err)
		return
	}

}

//判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true,nil
	}

	if os.IsNotExist(err) {
		return false,nil
	}

	return false,err
}

//实施文件拷贝
func CopyFile (dstFilename string, srcFileName string) (written int64, err error) {

	//通过文件绝对路径，获得文件对象（句柄）
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	defer srcFile.Close()

	//通过srcfile，获得到Reader-通过读取文件句柄将文件内容载入缓存
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstfile, err := os.OpenFile(dstFilename, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	defer dstfile.Close()

	writter := bufio.NewWriter(dstfile)

	return io.Copy(writter,reader)

}
//binaryFile Copy Sample
func filecopy_exe06_binaryCopy() {

	////	/Users/chongzhao/Downloads/IMG_5130.jpg

	srcFile := "/Users/chongzhao/Downloads/IMG_5130.jpg"
	dstFile := "/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/peggi.jpg"

	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("File Copy is accomplished.")
	} else {
		fmt.Println("File Copy error occured=%v\n", err)
	}

}

//统计结果保存结构体
type CharCount struct {
	Chount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func filecopy_exe07() {

	//思路；打开一个文件，创建一个Reader
	//每读取一样，就去统计该行有多少个英文，数字，空格和其他字符
	//然后将结果保存到一个结构体
	fileName := "/Users/chongzhao/Project/carrierUp/golang/chapter14_fileSystem/res/exec07"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	//延时关闭文件对象（句柄）
	defer file.Close()

	//定义一个CharCount实例
	var count CharCount

	//创建一个缓存Reader
	reader := bufio.NewReader(file)

	//开始循环的读取fileName的内容
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		//为了兼容中文字符将str转换为[]rune
		str = string([]rune(str))

			for _,val := range str {

				switch {
					case val >= int32('a') && val <= 'z':
						fallthrough
					case val >= 'A' && val <= 'Z':
						count.Chount++
					case val == ' ' || val == '\t':
						count.SpaceCount++
					case val >= '0' || val <= '9':
						count.NumCount++
					default:
						count.OtherCount++
				}
			}
	}

	//输出统计的结果看看是否正确
	fmt.Printf("CharCount=%v NumCount=%v SpaceCount=%v otherCount=%v\n", count.Chount,
		count.NumCount, count.SpaceCount, count.OtherCount)

}

//文件操作运行
func Run_fileoper() {
	//fileReadingWithBuf()
	//
	//fileReadingNobuf()
	//
	//fileWriter()

	//fileWrite_exe03_append()

	//fileWrite_exe04_readappend()

	//fileWrite_exe05_tranferfile()

	//filecopy_exe06_binaryCopy()

	filecopy_exe07()
}