package main

import "fmt"

var name string = "tom"

//var age := "rwwer" //编译错误，等价于 var age int ; age = "wr",赋值语句不能在函数体外，所以报错

func main() {
	name = "lilin"
	fmt.Println("name=", name)
}
