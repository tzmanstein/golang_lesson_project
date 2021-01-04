package interface_extends

import "fmt"

type Monkey struct {
	Name string
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}

type LittleMonkey struct {
	Monkey
}

type BirdCapbility interface {
	Flying()
}

type FishCapbility interface {
	Swimming()
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习，会飞！")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习，会游泳！")
}

func Run_interface_extends() {

	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()

}