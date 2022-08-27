package main

import (
	"fmt"
	"go/types"
)

//global constant-1, 全局变量声明方式1
var lilin = 10000

//global constant-2 全局变量申明方式2一次性申明全局变量
var (
	lilin2 = 1000
	name   = "lilin"
)

//global constant-1,

func main() {
	//定义变量/声明变量



	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))//3
	fmt.Printf("The capacity of s4: %d\n", cap(s4))//3
	fmt.Printf("The value of s4: %d\n", s4)//4,5,6
	var i int
	//变量复值
	i = 10
	fmt.Println(i)

	//1.有默认值，初始化默认值
	var n int
	fmt.Println("n=", n)

	//2.类型推倒
	var num = 11.20
	println("num=", num)

	// =:申明方式 省略var :=左边的变量不应该为已经声明过的，否则会编译错误
	//下面code等价于 var name string   name="xiaoming"
	name := "xiaoming"
	println("name=", name)

	//一次性声明多个变量-1
	var name1, name2, name3 int
	println("name1=", name1, "name2=", name2, "name3=", name3)

	//一次性声明多个变量-2
	//var a1, a2, a3 = 100, "xiaoming", 34
	//println("a1=", a1, "a2=", a2, "a3=", a3)

	//一次性声明多个变量-3
	a1, a2, a3 := 100, "xiaoming~~", 34
	println("a1=", a1, "a2=", a2, "a3=", a3)
	println("li=", lilin)

}
