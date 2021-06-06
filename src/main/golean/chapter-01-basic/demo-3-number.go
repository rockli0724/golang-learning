package main

import (
	"fmt"
	"unsafe"
)

//整数类型的使用
func main() {

	var i int = 1
	fmt.Println("i=", i)
	//有符号
	// int8 range -128~127
	var j int8 = -128
	fmt.Println("j=", j)

	//无符号 uint8(0-255) uint16 uint32 uint64
	var n uint8 = 255
	fmt.Println("n=", n)

	// int uint byte range
	//有符号 类似于int32
	// byte 等价于uint8 >0 0~255

	var b byte = 255
	fmt.Println("b=", b)

	//整形的使用细节
	var n1 = 7000 //n1是什么类型
	//fmt.Printf() 用于格式化输出
	fmt.Printf("n1 的类型是 %T", n1)

	//如何查看程序某个变量的占用字节大小和数据类型（用的比较多）
	var n2 int64 = 10
	//unsafe.Sizeof(n2) 是unsafe包的一个函数，可以返回n2占用的字节数
	fmt.Printf("n2 的类型是 %T, 占用类型大小字节数%d", n2, unsafe.Sizeof(n2))

}
