package main

import (
	"encoding/json"
	"fmt"
)

//结构体标签
//struct 的每个字段上，可以写上一个 tag, 该 tag 可以通过反射机制获取，常见的使用场景就是序 列化和反序列化
func main() {

	monster := Monster{
		"牛魔王", 222, "",
	}
	//序列化,使用到了反射的原理
	json, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("error=", err)
	}
	fmt.Println("result=", string(json))
}

type Monster struct {
	Name  string //`json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Skill string `json:"skill,omitempty"`
}
