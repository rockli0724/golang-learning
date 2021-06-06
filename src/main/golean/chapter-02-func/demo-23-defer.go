package main

import (
	"fmt"
)

//defer
//1.为什么需要 defer:
//	在函数中，程序员经常需要创建资源(比如:数据库连接、文件句柄、锁等) ，为了在函数执行完 毕后，及时的释放资源，Go 的设计者提供 defer (延时机制)。
//值传递、引用传递
//----》传递给函数的都是变量的副本，不同的是，
//值传递的是值的拷贝，引用传递的是地址的拷贝，一般来说，地址拷贝效率高，因为数据量小，
//而值拷贝决定拷贝的 数据大小，数据越大，效率越低。

func main() {
	i := sum(1, 2)
	fmt.Println("main.result=", i)

}

func sum(n1 int, n2 int) int {
	//执行defer时，暂不执行，会将defer后面的语句压入独立的defer栈中，执行先进后出
	//在 defer 将语句放入到栈时，也会将相关的值拷贝同时入栈
	defer fmt.Println("n1=", n1)
	defer fmt.Println("n2=", n2)

	n1++
	n2++
	result := n1 + n2
	fmt.Println("n1+n2=", result)
	return result
}
