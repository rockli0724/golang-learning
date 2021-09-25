package main

import (
	"fmt"
	"golang-learning/chapter-02-func/pkg/util"
)

var age = test()

func /**/ test() int {
	fmt.Println("test() ......")
	return 79
}

//全局执行顺序，先执行全局变量，在执行init(),在执行main方法
func main() {
	operation := util.Operation(23, 11, '+')
	fmt.Println("main(),oper=", operation)
}

//初始化。相当于Java的构造函数
func init() {
	fmt.Println("init()")
}
