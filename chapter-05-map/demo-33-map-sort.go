package main

import (
	"fmt"
	"sort"
)

//map排序
/*map 使用细节
1) map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的 map
2) map 的容量达到后，再想 map 增加元素，会自动扩容，并不会发生 panic，也就是说 map 能动 态的增长 键值对(key-value)
3) map 的 value 也经常使用 struct 类型，更适合管理复杂的数据(比前面 value 是一个 map 更好)，*/
func main() {

	//golang中map的排序，是先将key进行排序，然后根据key值遍历输出即可
	map1 := make(map[int]int, 10)
	map1[10] = 200
	map1[1] = 2300
	map1[3] = 3
	map1[5] = 3
	map1[2] = 2

	fmt.Println("map=", map1)

	//golang中map的排序，是先将key进行排序，然后根据key值遍历输出即可
	//1.将map的key放入切片中
	//2.对切片排序
	//3.遍历切片，让后安札key输出value
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println("key=", keys)

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}
