package main

import "fmt"

func main() {
	slice := fblc(10)
	fmt.Println("fin(10)=", slice)
}

//practice
/*说明:编写一个函数 fbn(n int) ，要求完成
1) 可以接收一个 n int
2) 能够将斐波那契的数列放到切片中 3) 提示, 斐波那契的数列形式:
arr[0] = 1; arr[1] = 1; arr[2]=2; arr[3] = 3; arr[4]=5; arr[5]=8*/
func fblc(n int) []uint64 {
	//声明一个切片，切片大小n
	fblcSlice := make([]uint64, n)
	//第一个和第二个为1
	fblcSlice[0] = 1
	fblcSlice[1] = 1
	//for 循环存在数列
	for i := 2; i < n; i++ {
		fblcSlice[i] = fblcSlice[i-1] + fblcSlice[i-2]
	}
	return fblcSlice
}
