package main

import (
	"fmt"
	"strconv"
)

//基本数据类型和string的转换
//方式一：fmt.Sprintf("%参数",表达式)
//方式二：使用strconv包的函数
func main() {

	var num1 int = 99
	var num2 float64 = 23.3453
	var b bool = true
	var myChar byte = 'h'
	var str string //空的str

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%v \n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%v \n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%v \n", str, str)

	//第二种方式，使用strconv 函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	//说明： 'f' 格式 10 表示小数位保留10位 64 表示这个小叔是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	//strconv 函数Itoa
	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n", str, str)

}
