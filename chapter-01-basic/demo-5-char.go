package main

import "fmt"

//字符类型的使用
//golang 没有专门的字符类型，由单个字节（byte）连接组成的
//单纯的字母一般用byte(0-255)来表示 ASCII码表字符
//英文字母-1个字节，汉字-3个字节
//1）字符型存储到计算机，需要将字符对应的码值（整数）找出来
// 存储: 字符-->对应码值-->二进制-->存储
// 读取: 二进制-->码值-->字符-->读取
// 读取: 二进制-->码值-->字符-->读取
//字符和码值的对应关系是通过字符编码表决定的（规定）
//Go语言的编码都统一成了UTF-8。非常的方便，很统一，解决了乱码的困扰
func main() {

	var c1 byte = 'a' // a==>97
	var c2 byte = '0' //数字0 ==>48
	//当我们直接输出byte值，就是输出了对应的ASCII码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2) //输出的是ASCII码

	//如果需要输出对应字符，用格式化输出即可
	fmt.Printf("c1=%c c2=%c \n", c1, c2)

	//
	//var c3 byte = '北'//overflow溢出
	var c3 int = '北' //21271
	fmt.Printf("c3=%c c3对应码值=%d \n", c3, c3)

	var c4 int = 21271 //ASCII 对应字符
	fmt.Printf("c4=%c \n", c4)

	//字符类型可以进行运算，按照码值运算
	var n1 = 10 + 'a' // ==> 10 +97 = 107
	fmt.Println("n1=", n1)
}
