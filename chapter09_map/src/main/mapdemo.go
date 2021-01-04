package main

import (
	"exec"
	"fmt"
	"util"
)

func mapsConstruct() {
	//第一种map构成方式
	var demoMap map[string]string
	demoMap = make(map[string]string, 10)
	demoMap["no01"] = "IronMan"
	demoMap["no02"] = "ThunderThor"
	demoMap["no03"] = "ThunderThor"

	fmt.Println(demoMap)

	//第二种类map构成方式 推荐
	cities := make(map[string]string)

	cities["01"]  = "Beijing"
	cities["02"]  = "Shanghai"
	cities["03"]  = "Shenzhen"
	fmt.Println(cities)

	//第三种map构成方式
	//var heros map[string]string = map[string]string {
	var heros  = map[string]string {
		"hero01" : "Ironman",
		"hero02" : "Superman",
		"hero03" : "Batman",
	}
	fmt.Println(heros)

}


func main() {

	util.MapsConstruct()
	//mapsConstruct()

	//fmt.Println()
	//exec.Exec01()

	util.Maptraverse()

	util.MapTravseComlicate()

	util.Mapslice()

	util.Mapsort()

	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	fmt.Println(map1)
	util.Modifymap(map1)
	fmt.Println(map1)

	//map配合结构体使用
	util.MapandStruct()

	exec.Exec02()


}