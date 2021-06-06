package main

import "fmt"

//小数类型的使用
//浮点型的存储分为三个部分：符号位+指数位+尾数位，在存储过程中，会有精度损失
func main() {
	//float32 单精度,有精度损失的
	var price float32 = 43.22
	fmt.Println("price=", price)
	var num1 float32 = 432.3312312312311231312312423412341
	var num2 float64 = 432.3312312312311231312312423412341
	fmt.Println("num1=", num1, "num2=", num2)

	var num4 = .1
	fmt.Printf("类型是 %T \n", num4)

	//十进制的形式：如5.12 .231
	num5 := 5.12
	num6 := .123 // =>0.123
	fmt.Println("num5=", num5, "num6=", num6)

	//科学计数法
	num7 := 5.1214e2  // =>5.1214 * 10^2
	num8 := 5.1214e2  // =>5.1214 * 10^2
	num9 := 5.1214e-2 // =>5.1214 / 10^2
	fmt.Println("num7=", num7, "num8=", num8, "num9=", num9)

	//通常情况我们都推荐使用float64位【约定俗成】

}
