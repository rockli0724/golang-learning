package main

import "fmt"

//字符串类型的使用
// 一串固定长度的字符连接起来的字符序列，由单个字节连接起来的，
//使用细节
//1.统一使用utf-8编码，避免乱码问题
//2.字符串一但赋值，字符串就不能修改了，在golang中string是不可变的
//3.字符创的两种表现形式
//	1.双引号，会识别转义字符
//	2.反引号，以字符创的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
//4.字符串拼接方式注意点

func main() {
	var str string = "abc\n * abd"
	fmt.Println("str=", str)

	//字符串一但赋值，字符串就不能修改了
	//var str2 = "hello"
	//str2[0] = 'a'//不能修改str2的内容

	//2.反引号，以字符创的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代
	str3 := "abc\ndef"
	fmt.Println(str3)

	//输出源代码 反引号 ` `
	str4 := `
package main

import "fmt"

//字符串类型的使用
// 一串固定长度的字符连接起来的字符序列，由单个字节连接起来的，
//使用细节
//1.统一使用utf-8编码，避免乱码问题
//2.字符串一但赋值，字符串就不能修改了，在golang中string是不可变的
//3.字符创的两种表现形式
//	1.双引号，会识别转义字符
//	2.反引号，以字符创的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
chapter-02-func main() {
	var str string = "abc\n * abd"
	fmt.Println("str=", str)

	//字符串一但赋值，字符串就不能修改了
	//var str2 = "hello"
	//str2[0] = 'a'//不能修改str2的内容

	//2.反引号，以字符创的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代
	str3 := "abc\ndef"
	fmt.Println(str3)

	//输出源代码 反引号
}
`
	fmt.Println("str4=", str4)

	//字符串拼接方式,多行时，需要将+放入行尾
	var str5 = "hello" + "world" +
		" hahhaha"
	fmt.Println(str5)
}
