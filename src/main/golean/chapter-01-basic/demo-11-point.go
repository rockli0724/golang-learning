package main

import "fmt"

/*
指针类型
获取变量的地址：用&变量
1.基本数据类型，变量存的值，也叫值类型
2.获取变量的地址，用&，比如 var num int, 获取num的地址： &num
3.指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值，比如： var ptr *int = &num
4.获取指针类型所指向的值，只用：*，比如 var ptr *int,使用 *ptr获取p指向的值
5.举例说明
	var num int = 1
	var i = 999
	var ptr *int =&num*/
func main() {

	//基本数据类型在内存中布局
	var i int = 10
	//i的地址= 0xc000124008
	fmt.Println("i的地址=", &i)

	//下面的	var ptr *int = &i
	//1. ptr 是一个指针变量
	//2. ptr 的类型是 *int
	//3. ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v \n", &ptr)
	fmt.Printf("ptr指向的值=%v \n", *ptr)

	//var num int =1
	var i2 = 999
	var ptrs *int = &i2
	fmt.Printf("i2=%v\n", &i2)
	fmt.Printf("ptrs=%v\n", ptrs)
	fmt.Printf("ptrs的地址=%v \n", &ptrs)
	fmt.Printf("ptrs指向的值=%v \n", *ptrs)

}
