package main

import "fmt"

//+的使用
func main() {
	//数字类型为相加
	var i = 1
	var j = 2
	fmt.Println(i + j)

	//字符类型为拼接
	var str1 = "hello"
	var str2 = "world"
	fmt.Println(str1 + str2)
}
