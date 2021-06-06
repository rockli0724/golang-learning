package main

import "fmt"

//演示golang中基本数据类型的转换
//细节说明
//1.标识范围小-->表示范围大
//2.被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
//3.在转换中，比如将int64转成int8【-128~127】,编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样
func main() {
	var i int32 = 100
	//希望将i转成float类型
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i) //低精度-->高精度
	fmt.Printf("i=%v,类型=%T n1=%v 类型=%T n2=%v 类型=%T n3=%v 类型=%T", i, i, n1, n1, n2, n2, n3, n3)

	//在转换中，比如将int64转成int8【-128~127】,编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样
	var num1 int64 = 99999
	var num2 int8 = int8(num1)
	fmt.Println("\n num2=", num2)

}
