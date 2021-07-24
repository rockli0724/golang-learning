package main

import "fmt"

// goto 和
//return:表示跳出所在的方法或函数。如果是在main函数中，表示终止整个程序

func main() {

	var num int = 19
	fmt.Println("ok1")
	if num > 10 {
		goto label1
	}
	println("ok2")
	println("ok3")
	println("ok4")
label1:
	println("ok5")
	println("ok6")
	println("ok7")

}
