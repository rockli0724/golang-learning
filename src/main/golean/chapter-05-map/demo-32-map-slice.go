package main

import "fmt"

//map切片
//切片的数据类型如果是 map，则我们称为 slice of map，map 切片，这样使用则 map 个数就可以动 态变化了
func main() {
	//演示map切片的使用
	//1声明一个map切片
	var monsters []map[string]string
	//声明切片容量
	monsters = make([]map[string]string, 2) //只能添加两个。多了会报错
	//2.增加第一个妖怪的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "牛魔王2"
		monsters[1]["age"] = "500-2"
	}
	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string, 2)
	//	monsters[2]["name"] = "牛魔王-3"
	//	monsters[2]["age"] = "500-3"
	//}
	//动态添加，可用append函数
	newMonster := map[string]string{
		"name": "火云邪神",
		"age":  "222"}
	monsters = append(monsters, newMonster)
	fmt.Println("monter=", monsters)
}
