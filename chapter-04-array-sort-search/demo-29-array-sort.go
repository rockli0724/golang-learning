package main

import "fmt"

func main() {
	var arr = [5]int{2, 4, 8, 5, 1}
	DoubleSort12(&arr)
}

//冒泡排序(从小到大)
func DoubleSort1(arr *[5]int) {
	fmt.Println("排序前=", *arr)

	tmp := 0 //临时变量
	//for i := 0; i < len(*arr)-1; i++ {
	for j := 0; j < len(*arr)-1; j++ {
		//如果前一个数>后面一个数
		if (*arr)[j] > (*arr)[j+1] {
			//数据交换
			tmp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = tmp
		}
	}
	fmt.Println("第1次排序=", *arr)

	for j := 0; j < len(*arr)-2; j++ {
		//如果前一个数>后面一个数
		if (*arr)[j] > (*arr)[j+1] {
			//数据交换
			tmp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = tmp
		}
	}
	fmt.Println("第2次排序=", *arr)

	for j := 0; j < len(*arr)-3; j++ {
		//如果前一个数>后面一个数
		if (*arr)[j] > (*arr)[j+1] {
			//数据交换
			tmp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = tmp
		}
	}
	fmt.Println("第3次排序=", *arr)

	for j := 0; j < len(*arr)-4; j++ {
		//如果前一个数>后面一个数
		if (*arr)[j] > (*arr)[j+1] {
			//数据交换
			tmp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = tmp
		}
	}
	fmt.Println("第4次排序=", *arr)
	//}

}

//冒泡排序
func DoubleSort12(arr *[5]int) {
	fmt.Println("排序前=", *arr)

	tmp := 0 //临时变量
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			//如果前一个数>后面一个数
			if (*arr)[j] > (*arr)[j+1] {
				//数据交换
				tmp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
		fmt.Printf("第%v次排序，结果=%v \n", i, *arr)
	}

}
