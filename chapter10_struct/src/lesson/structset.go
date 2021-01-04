package lesson


// 定义结构体
type Cat struct {

	//字段属性名首字母大写，公开内容
	Name string
	Age int
	Color string
	Hobby string
	Score [3]int

	// 可以在结构体中绑定方法

}

type Student struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}

type Monster struct {
	Name string
	Age int
}

type Humman struct {
	Name string
	Age int
}

type Point struct {
	x int
	y int //int默认为int64 8个byte
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

type A struct {
	Num int
}

type B struct {
	Num int
}

type Hummm Humman

/**
 * 结构体标签tag的使用
 */
type Monster_n struct {
	Name string `json:"name"` //全局字段名小写化，方便json访问。
	Age int `json:"age"`
	Skill string `json:"skill"`
}

//