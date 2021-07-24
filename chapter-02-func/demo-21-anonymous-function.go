package main

import "fmt"

var (
	funcGlobal = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

//匿名函数
func main() {

	//1..在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	result := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("result=", result)

	//2.将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数
	func2 := func(n1 int, n2 int) int {
		return n1 + n2
	}
	result2 := func2(1, 2)

	fmt.Println("result2=", result2)

	//3.全局匿名函数
	result3 := funcGlobal(3, 5)

	fmt.Println("result3=", result3)
}
