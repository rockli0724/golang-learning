package main

import (
	"fmt"
	"unsafe"
)

//布尔类型的使用 只能是true/false
//bool类型占用1个字节
//bool类型适用于逻辑计算
func main() {
	var b = false
	fmt.Println("b=", b)
	//bool占用一个字节
	fmt.Println("占用字节=", unsafe.Sizeof(b))

}
