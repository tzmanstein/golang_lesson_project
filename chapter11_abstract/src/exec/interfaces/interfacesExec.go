package interfaces

import (
	"fmt"
	"math/rand"
	"sort"
)

type A_ExecIF interface {
	Test01()
	Test02()
}

type B_ExecIF interface {
	Test01()
	Test02()
}

type C_ExecIF interface {
	A_ExecIF
	B_ExecIF
}

type Student struct {
	Name string
	Age int
	Score float64
}

type StudentSlice []Student

func (ss StudentSlice) Len() int{
	return len(ss)
}

func (ss StudentSlice) Less(i, j int) bool {
	//return hs[i].Age < hs[j].Age
	return ss[i].Score < ss[j].Score
}

//元素交换
func (ss StudentSlice) Swap(i, j int) {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	// 上下处理功能相同
	ss[i], ss[j] = ss[j], ss[i]

}

// Less方法决定使用什么标准进行排序
// 按照年龄升序排序


//hero结构体切片排序
func (hs HeroSlice) Len() int{
	return len(hs)
}

// Less方法决定使用什么标准进行排序
// 按照年龄升序排序
func (hs HeroSlice) Less(i, j int) bool {
	//return hs[i].Age < hs[j].Age
	return hs[i].Name < hs[j].Name
}

//元素交换
func (hs HeroSlice) Swap(i, j int) {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	// 上下处理功能相同
	hs[i], hs[j] = hs[j], hs[i]

}

//1. 声明一个结构体
type Hero struct {
	Name string
	Age int
}

//2. 声明结构体切片类型
type HeroSlice []Hero

//3. 实现Interface接口

func Run_ExecIFs() {

	var intSlice = []int{0, -1, 10, 7, 90}

	//1.冒泡排序。。
	//2.使用系统Api排序
	sort.Ints(intSlice)
	fmt.Println(intSlice)


	//对一个结构体slice进行排序
	//sort.Sort()
	//测试是否可以对结构体切片进行排序
	var heros HeroSlice

	for i := 0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age : rand.Intn(100),
		}
		//加入
		heros = append(heros, hero)
	}
	for _,v := range heros{
		fmt.Println(v)
	}
	fmt.Println()

	sort.Sort(heros)

	for _, v := range heros {
		fmt.Println(v)
	}
}