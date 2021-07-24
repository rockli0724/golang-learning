package main

import (
	"fmt"
	"strconv"
	"time"
)

//golang 内建函数
func main() {
	//timeFUnc()

	//printSleep()
	//unix 和unixNano
	fmt.Println("unix=", time.Now().Unix(), "unixNano=", time.Now().UnixNano())

	start := time.Now().Unix()
	//timeFunc2()
	end := time.Now().Unix()
	fmt.Println("耗时秒数=", end-start)
	intType()

}

//日期函数
func timeFUnc() {
	now := time.Now()
	fmt.Println("Year=", now.Year())
	fmt.Println("Month=", now.Month())
	fmt.Println("Day=", now.Day())
	fmt.Println("Hour=", now.Hour())
	fmt.Println("Minute=", now.Minute())
	fmt.Println("Second=", now.Second())
	format := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("format=", format)

}

//每一秒输出一个数字，直到打印到100停止
func printSleep() {
	i := 0
	for {
		i++
		fmt.Println("i=", i)
		time.Sleep(time.Second * 1)
		if i == 20 {
			break
		}
	}
}

func timeFunc2() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func intType() {
	num1 := 100
	fmt.Printf("num1的类型%T, 值=%v, 地址=%v \n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2的类型%T, 值=%v, 地址=%v, 指针值=%v", num2, num2, &num2, *num2)

}
