package main

import (
	"fmt"
	"strings"
)

//闭包
/*1.AddUpper 是一个函数，返回的数据类型是 fun (int) int
2.闭包的说明：返回的是一个匿名函数, 但是这个匿名函数引用到函数外的 n ,因此这个匿名函数就和 n 形成一 个整体，构成闭包
3.家可以这样理解: 闭包是类, 函数是操作，n 是字段。函数和它使用到 n 构成闭包。
4.当我们反复的调用 f 函数时，因为 n 是初始化一次，因此每调用一次就进行累计。
5.我们要搞清楚闭包的关键，就是要分析出返回的函数它使用(引用)到哪些变量，因为函数和它引用到的变量共同构成闭包*/
func main() {
	upper := AddUpper()
	fmt.Println("f(1)=", upper(1))
	fmt.Println("f(2)=", upper(2))
	fmt.Println("f(3)=", upper(3))
	fmt.Println("f(4)=", upper(4))

	suffix := makeSuffix(".jpg")
	fmt.Println("fileName1=", suffix("winter.jpg"))
	fmt.Println("fileName2=", suffix("spring"))
	fmt.Println("fileName2=", suffix("spring2.avi"))

}

func AddUpper() func(int) int {
	var n int = 10
	var str string = "hello"

	return func(i int) int {
		n = n + i
		str += "_"
		fmt.Println("str=", str)
		return n
	}
}

/*请编写一个程序，具体要求如下。
1) 编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
2) 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如果已经有.jpg 后缀，则返回原文件名。
3) 要求使用闭包的方式完成
4) strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。*/

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
