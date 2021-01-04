package util

import (
	"fmt"
	"sort"
)

// 定义一个结构体
type Stu struct {
	Name string
	Age int
	Address string
}

func MapsConstruct() {

	//第一种map构成方式
	var demoMap map[string]string
	demoMap = make(map[string]string, 10)
	demoMap["no01"] = "IronMan"
	demoMap["no02"] = "ThunderThor"
	demoMap["no03"] = "ThunderThor"

	fmt.Println(demoMap)

	//第二种类map构成方式
	cities := make(map[string]string)
	cities["01"]  = "Beijing"
	cities["02"]  = "Shanghai"
	cities["03"]  = "Shenzhen"
	fmt.Println(cities)
	fmt.Println("CCCCCCcity pairs= ", len(cities))

	//第三种map构成方式
	//var heros map[string]string = map[string]string {
	var heros  = map[string]string {
		"hero01" : "Ironman",
		"hero02" : "Superman",
		"hero03" : "Batman",
	}
	fmt.Println(heros)


	//delete(cities, "01")
	//使用make重新删除 推荐使用的全删除方法。
	//cities = make(map[string]string)
	//fmt.Println("deleted cities", cities)

	//演示map的查找
	val, ok := cities["01"]
	if ok {
		fmt.Printf("existed=%v\n", val)
	} else {
		fmt.Printf("not existed\n")
	}

}

// map 遍历只能使用for range
func Maptraverse() {

	//第二种类map构成方式
	cities := make(map[string]string)
	cities["01"]  = "Beijing T"
	cities["02"]  = "Shanghai T"
	cities["03"]  = "Shenzhen T"

	for i,v := range cities {
		fmt.Printf("No %vcity is %v\n", i ,v)
	}


}

func MapTravseComlicate() {
	studentMap := make(map[string]map[string]string)
	studentMap["stud01"] = make(map[string]string, 3)
	studentMap["stud01"]["name"] = "tom"
	studentMap["stud01"]["sex"] = "male"
	studentMap["stud01"]["address"] = "beijing"

	studentMap["stud02"] = make(map[string]string, 3)
	studentMap["stud02"]["name"] = "mary"
	studentMap["stud02"]["sex"] = "female"
	studentMap["stud02"]["address"] = "shanghai"

	//fmt.Println(studentMap)

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2,v2 := range v1 {
			fmt.Printf("\tk2=%v, v2=%v\n", k2, v2)
		}
	}

}

/**
 * 演示map切片
 */
func Mapslice() {
	//演示map切片
	// 1. 声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2) //slice本身需要make
	//2. 增加一个妖精信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string) //slice的map也需要make
		monsters[0]["name"] = "BUlL"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string) //slice的map也需要make
		monsters[1]["name"] = "Rabbit"
		monsters[1]["age"] = "400"
	}

	//panic: runtime error: index out of range [2] with length 2
	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string) //slice的map也需要make
	//	monsters[2]["name"] = "狐狸精"
	//	monsters[2]["age"] = "300"
	//}

	newmonter := map[string]string {
		"name" : "xinyaoguai",
		"age" : "200",
	}
	monsters = append(monsters, newmonter)


	fmt.Println(monsters)

}
/**
 * Map排序
 */
func Mapsort() {

	map1 := make(map[int]int)
	map1[10] = 100
	map1[1] = 123
	map1[4] = 56
	map1[8] = 190

	fmt.Println(map1)

	// 如果按照map的key顺序进行排序输出
	// 1. 先将map的key放入切片中
	// 2. 对切片进行排序
	// 3. 遍历切片，然后按照key来输出map值
	var keys []int
	for k,_ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _,k := range keys {
		fmt.Printf("map1[%v]=%v\n", k, map1[k])
	}

}


func Modifymap(map1 map[int]int) {
	map1[10] = 900

}

func MapandStruct() {
	students := make(map[string]Stu, 10)
	stu1 := Stu{"tom", 18, "Beijing"}
	stu2 := Stu{"marry", 28, "Shanghai"}
	students["no01"] = stu1
	students["no02"] = stu2
	fmt.Println("main", students)

	for k,v := range students {
		fmt.Printf("ID=%v ", k)
		fmt.Printf("Name=%v ", v.Name)
		fmt.Printf("Age=%v ", v.Age)
		fmt.Printf("Addr=%v", v.Address)
		fmt.Println()
	}
}