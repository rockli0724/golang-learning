package main

import (
	"errors"
	"fmt"
)

//错误处理机制
/*1.在默认情况下，当发生错误后(panic) ,程序就会退出(崩溃.)
2.如果我们希望:当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。还可 以在捕获到错误后，给管理员一个提示(邮件,短信。。。)
3.这里引出我们要将的错误处理机制

1) Go 语言追求简洁优雅，所以，Go 语言不支持传统的 try...catch...finally 这种处理。
2) Go中引入的处理方式为:defer,panic,recover
3) 这几个异常的使用场景可以这么简单描述:Go 中可以抛出一个 panic 的异常，然后在 defer 中通过 recover 捕获这个异常，然后正常处理
*/
func main() {
	//testMath()
	customError()
	fmt.Println("mian()...")
}

func testMath() {

	//使用defer+recover 捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
			fmt.Println("处理错误+发送邮件信息。。。")
		}
	}()
	num1 := 20
	num2 := 0
	result := num1 / num2
	fmt.Println("result=", result)
}

//自定义错误
func customError() {
	err := readConf("configs.ini")
	if err != nil {
		//读取文件发生错误，则输出错误信息并终止程序
		panic(err)
	}
	fmt.Println("customError...")
}

func readConf(name string) error {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}
