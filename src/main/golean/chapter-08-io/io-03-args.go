package main

import (
	"flag"
	"fmt"
)

//func main() {
//
//	fmt.Println("命令行的参数有", len(os.Args))
//	for i, arg := range os.Args {
//		fmt.Printf("args[%v]=%v \n", i, arg)
//	}
//}

func main() {

	//定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "host", "localhsot", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	//转换
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
