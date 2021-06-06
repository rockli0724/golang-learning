package main

import "fmt"

// for循环

func main() {
	//1.输出10句“你好”
	for i := 0; i < 10; i++ {
		fmt.Println("你好")
	}
	fmt.Println("-------------")

	//2.
	j := 0
	for j < 10 {
		fmt.Println("hello")
		j++
	}

	fmt.Println("-------------")

	//3.
	k := 1
	for {
		if k <= 10 {
			fmt.Println("ok")
		} else {
			break //跳出循环
		}
		k++
	}

	fmt.Println("-------------")
	//4.循环遍历字符串[传统方式]
	var str string = "hello,北京!"
	str2 := []rune(str) //切片，解决循环中文问题，utf8中一个中文对应3个字节
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	fmt.Println("-------------")
	str = "abc~def北京"
	//5. for range
	for index, value := range str {
		fmt.Printf("index=%d, value=%c \n", index, value)
	}

	fmt.Println("-------------")
	//6.打印1~100中所有9的倍数个个数以及和
	var max uint64 = 100
	var count uint64 = 0
	var sum uint64 = 0
	var i uint64 = 0

	for ; i < max; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%v sum=%v \n", count, sum)

	fmt.Println("-------------")
	//输出
	var ct int = 6
	for i := 0; i <= ct; i++ {
		fmt.Printf("%v + %v = %v \n", i, ct-i, ct)
	}

}
