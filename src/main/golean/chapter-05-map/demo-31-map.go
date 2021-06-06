package main

import "fmt"

func main() {
	//map的声明1
	var amap map[string]string
	//在使用map前 需要make(分配数据空间)
	amap = make(map[string]string, 10)

	amap["no2"] = "宋江"
	amap["no-1"] = "宋江"
	amap["no1"] = "宋s江"
	amap["no3"] = "宋江2"
	fmt.Println("map=", amap)

	//2.第二种方式
	map2 := make(map[string]string)
	map2["no1"] = "rockLi"
	fmt.Println("map2=", map2)

	//2.第二种方式
	map3 := map[string]string{
		"no2": "lilin",
		"no1": "lijia",
	}
	fmt.Println("map3=", map3)

	mapFunc()

}

//1.演示一个 key-value 的 value 是 map 的案例
//比如:我们要存放 3 个学生信息, 每个学生有 name 和 sex 信息
//思路: map[string]map[string]string

func mapFunc() {
	studentMap := make(map[string]map[string]string)
	studentMap["01"] = make(map[string]string, 2)
	studentMap["01"]["name"] = "小明"
	studentMap["01"]["sex"] = "男"
	studentMap["01"]["age"] = "22"

	studentMap["02"] = make(map[string]string, 3)
	studentMap["02"]["name"] = "小女"
	studentMap["02"]["sex"] = "女"
	studentMap["02"]["age"] = "24"

	studentMap["03"] = make(map[string]string, 3)
	//新增
	studentMap["03"]["name"] = "小明3"
	studentMap["03"]["sex"] = "男"
	studentMap["03"]["age"] = "22"

	//修改map
	//delete(studentMap["01"], "name")
	delete(studentMap["02"], "nam2e")

	//查询key1
	value, findResult := studentMap["03"]["age"]
	if findResult {
		fmt.Println("m=", value)
	}
	//查询key2
	findResult2 := studentMap["03"]["age"]
	fmt.Println("findResult2=", findResult2)

	fmt.Println("stu=", studentMap)

	//map删除，遍历-逐一删除；2-直接make新的空间
	//2.重新分配空间
	//studentMap["01"] = make(map[string]string)
	//fmt.Println("stu[\"01\"]=", studentMap["01"])

	//map遍历 用for range

	for key, value := range studentMap {
		fmt.Println("key=", key)
		for v_key, v_value := range value {
			fmt.Printf("key=%v, value=%v \n", v_key, v_value)
		}
		fmt.Println()
	}

	fmt.Println("map.len=", len(studentMap))

}
