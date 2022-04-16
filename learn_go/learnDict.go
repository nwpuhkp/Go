package main

import "fmt"

var opMap = func(name map[string]int) {
	for key, value := range name {
		fmt.Println(key, value)
	}
	name["Life"] = 100
	if value, ok := name["Go"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("no exists Go")
	}
	delete(name, "java")
}

func main() {
	nameMap := make(map[string]int) //map是引用类型，使用make初始化
	nameMap["java"] = 200
	nameMap["php"] = 100
	nameMap["python"] = 180
	nameMap["js"] = 220
	//nameMap["Go"] = 100
	opMap(nameMap)
}
