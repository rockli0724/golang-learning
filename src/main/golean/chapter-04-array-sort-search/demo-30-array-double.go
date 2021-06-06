package main

import "fmt"

func main() {

	var arr [3][2]int //[3]--二维数组共有多少元素（行数），[2]-一维数组中的元素个数（每一行中元素个数）
	arr[0][0] = 0
	arr[0][1] = 1
	arr[1][0] = 2
	arr[1][1] = 3
	arr[2][0] = 4
	arr[2][1] = 5

	fmt.Println("arr=", arr)

	var arr2 [3][2]int = [3][2]int{{1, 2}, {4, 5}, {7, 8}}
	var arr3 = [3][2]int{{1, 2}, {4, 5}, {7, 8}}
	var arr4 = [...][2]int{{1, 2}, {4, 5}, {7, 8}}
	var arr5 = [3][2]int{{1, 2}, {4, 5}, {7, 8}}
	arr6 := [3][2]int{{1, 2}, {4, 5}, {7, 8}}
	fmt.Println("arr2=", arr2)
	fmt.Println("arr3=", arr3)
	fmt.Println("arr4=", arr4)
	fmt.Println("arr5=", arr5)
	fmt.Println("arr6=", arr6)

	//遍历二维数组
	for i, v := range arr4 {
		for i2, i3 := range v {
			fmt.Printf("arr[%v][%v]=%v \t ", i, i2, i3)
		}
		fmt.Println()
	}
}
