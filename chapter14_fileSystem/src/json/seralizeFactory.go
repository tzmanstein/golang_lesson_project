package json

import (
	"encoding/json"
	"fmt"
)

func Run_seraliztion() {
	structSeralizaton()

	mapSeralization()

	sliceSeralization()

	basicdataSeralization()
}

//定义一个结构体
//如果需要跨包使用struct时，struct的成员必须使用首字母大写
type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday string `json:"birthday"`
	Sal float64 `json:"salary"`
	Skill string `json:"skill"`
}

//结构体的序列化
func structSeralizaton() {
	var monster Monster = Monster{
		Name: "牛魔王",
		Age: 500,
		Birthday: "2011-11-11",
		Sal: 8000.0,
		Skill: "NiuMOquan",
	}
	jsonByte, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("Fail to seralization of struct=%v\n", err)
	}

	fmt.Printf("the Result of STRUCT seralization= %v\n", string(jsonByte))
}

// Map的序列化
func mapSeralization() string{
	//定义一个map
	var a map[string]interface{}

	//map使用前，需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿-鬼畜"
	a["age"] = 30
	a["address"] = "火云洞"

	jsonData, err := json.Marshal(a)
	if err != nil {

	}

	fmt.Printf("the Result of MAP seralization= %v\n", string(jsonData))
	return  string(jsonData)

}

//Slice序列化
func sliceSeralization() {
	var slice []map[string]interface{}

	slice01 := make(map[string]interface{})
	slice01["name"] = "jack"
	slice01["age"] = "7"
	slice01["address"] = "Beijing"
	slice = append(slice, slice01)

	slice02 := make(map[string]interface{})
	slice02["name"] = "tom"
	slice02["age"] = "20"
	//slice02["address"] = "Mexica"
	slice02["address"] = [2]string{"Mexico", "Hawayi"}//数组形式
	slice = append(slice, slice02)

	jsonData, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("Fail to seralization of SLICE=%v\n", err)
	}
	fmt.Printf("the Result of SLICE seralization= %v\n", string(jsonData))
}

//基本数据类型序列化，基本数据类型序列化无意义。
func basicdataSeralization() {

	num1 := 2345.67
	//var num2 float64 = 23698.23

	jsonData, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("Fail to seralization of BASICDATA=%v\n", err)
	}
	//序列化结果是转化成字符串
	fmt.Printf("the Result of BASICDATA seralization= %v\n", string(jsonData))
}