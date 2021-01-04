package interfaces

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

//实现一个方法Working
// 此方法接收一个方法，接收一个Usb接口类型变量
// 只要实现Usb接口，所谓实现接口就是 实现Usb接口声明的所有方法
func (c *Computer) Working(usb Usb) {

	//通过Usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()

}

func Run_interfaces() {
	// 先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//调用
	computer.Working(phone)
	computer.Working(camera)

	//实现多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"xiaomi"}
	usbArr[1] = Phone{"vivo"}
	usbArr[2] = Camera{"canon"}

	fmt.Println(usbArr)


	}