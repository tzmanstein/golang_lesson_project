package exec

import "fmt"

func Exec01() {

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

func modifyUser(users map[string]map[string]string, name string, nickname string)  {

	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = nickname
	}

}

func Exec02() {
	users := make(map[string]map[string]string)
	users["smith"] = make(map[string]string)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "大老虎"

	modifyUser(users,"tom", "tt")
	modifyUser(users, "mary", "mm")
	fmt.Println(users)


}