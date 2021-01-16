package json

import (
	"encoding/json"
	"fmt"
)

func Run_unseralization() {
	structUnseralization()	//stuct反序列化
	mapUnseralization()		//map反序列化
	sliceUnseralization()	//slice反序列化
}

func structUnseralization() {
	//get string from some response such as, file, http, etc.
	var str string = "{\"name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"NiuMOquan\"}"

	//反序列化用struct必须根据传输内容记性预定义
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster) //接收结构体必须是引用（&）传值

	if err != nil {
		fmt.Printf("err= %v\n", err)
	}
	fmt.Printf("STRUCT 反序列化结果=%v\n, Monster.name=%v\n", monster, monster.Name)

}

func mapUnseralization() {
	//var strMap = "{\"address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}"

	var strMap = mapSeralization()

	//反序列化用map必须根据传输内容记性预定义
	var mapRes map[string]interface{}

	//反序列化时，map不需要make。map操作直接封装到Unmarshal中会对map进行make。
	err := json.Unmarshal([]byte(strMap), &mapRes) //必须是地址

	if err != nil {
		fmt.Printf("err= %v\n", err)
	}
	fmt.Printf("MAP 反序列化结果=%v\n", mapRes)
}

func sliceUnseralization() {
	strSlice := " [{\"address\":\"Beijing\",\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"Mexico\",\"Hawayi\"],\"age\":\"20\",\"name\":\"tom\"}]"

	//反序列化用slice必须根据传输内容记性预定义
	var sliceRes []map[string]interface{}

	err := json.Unmarshal([]byte(strSlice), &sliceRes) //必须是地址

	if err != nil {
		fmt.Printf("err= %v\n", err)
	}
	fmt.Printf("SLICE 反序列化结果=%v\n", sliceRes)
}