package asserts

import "fmt"

// 模拟现实Usb接口实现
type Usb interface {
	//声明2个未实现方法
	Start()
	Stop()
}

// Phone 手机结构体
type Phone struct {
	Name string

}
// Phone 结构体实现Usb接口方法
func (p Phone) Start() {
	fmt.Println("手机开始连接....")
}

func (p Phone) Stop() {
	fmt.Println("手机停止连接....")
}

func (p Phone) Call() {
	fmt.Println("手机在打电话...")
}

// Camera相机结构体
type Camera struct {
	Name string

}
// Camera 结构体实现Usb接口方法
func (c Camera) Start() {
	fmt.Println("相机开始连接....")
}

func (c Camera) Stop() {
	fmt.Println("相机停止连接....")
}

type Computer struct {

}

func (cp Computer) Working(usb Usb){
	usb.Start()

	if tmpPhone, ok := usb.(Phone); ok {
		tmpPhone.Call()
	} else {
		fmt.Println("不是电话")
	}


	usb.Stop()
}

func Run_AssertApply() {
	//实现多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"xiaomi"}
	usbArr[1] = Phone{"vivo"}
	usbArr[2] = Camera{"canon"}

	var computer Computer
	for _, value := range usbArr {
		computer.Working(value)



		fmt.Println()
	}




}
