package main

import "fmt"

// 打印键盘输入
//可以从控制台接收用户信息，【姓名，年龄，薪水, 是否通过考试 】
func main() {

	var name string
	var age byte
	var sal float64
	var isPass bool

	// 1.
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过")
	fmt.Scanln(&isPass)

	//2
	fmt.Println("请输入 姓名 年龄 薪水 是否通过考试，中间用空格隔开")

	fmt.Scanf("%d %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字%v \t 年龄%v \t 薪水%v \t 通过考试%v \n", name, age, sal, isPass)

}
