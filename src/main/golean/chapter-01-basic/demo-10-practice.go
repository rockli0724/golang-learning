package main

import (
	_ "fmt" // 符号_标识忽略
)

func main() {
	//var n1 int32 = 12
	//var n2 int64
	//var n3 int8
	//
	////n2 = n1 + 20 //结果是int32,n2是int64 不能转换
	//n2 = int64(n1) + 20 //修改
	////n3 = n1 + 20 //结果是int32,n3是int8 不能转换
	//n3 = int8(n1) + 20 //修改

	//var n1 int32 = 12
	//var n3 int8
	//var n4 int8
	//n4 = int8(n1) + 127 //编译通过，但是结果不是 127+32 按溢出处理
	////n3 = int8(n1) + 128 //编译不通过
	//fmt.Println(n3)
	//fmt.Println(n4)
}
