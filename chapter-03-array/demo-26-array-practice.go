package main

import (
	"fmt"
	"math/rand"
	"time"
)

//数组的使用

func main() {
	printA2C()
	fmt.Println("----------")
	maxValueIndex()
	fmt.Println("----------")
	avgSum()
	fmt.Println("----------")
	radomNumberAndReverse()
}

//1.创建一个 byte 类型的 26 个元素的数组，分别 放置'A'-'Z‘。使用 for 循环访问所有元素并打印 出来。提示:字符数据运算 'A'+1 -> 'B'

func printA2C() {
	var myChar [26]byte
	for i := 0; i < 26; i++ {
		myChar[i] = 'A' + byte(i)
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c", myChar[i])
	}
}

//请求出一个数组的最大值，并得到对应的下标
func maxValueIndex() {
	var arr = [...]int{100, 2, 4, -11, 24, 299, 555, 33}
	maxValue := 0
	maxValueIndex := arr[0]
	for i := 1; i < len(arr); i++ {
		if maxValue < arr[i] {
			maxValue = arr[i]
			maxValueIndex = i
		}
	}
	fmt.Printf("arra=%v, maxValue=%v,maxValueIndex=%v", arr, maxValue, maxValueIndex)
}

//请求出一个数组的和和平均值。for-range
func avgSum() {
	var arr = [...]int{100, 2, 4, -11, 24, 299, 555, 33}
	sum := 0
	for _, value := range arr {
		sum += value
	}
	fmt.Printf("sum=%v,avg=%v", sum, float64(sum)/float64(len(arr)))
}

//要求:随机生成五个数，并将其反转打印 , 复杂应用
func radomNumberAndReverse() {
	var arr [6]int
	rand.Seed(time.Now().UnixNano())
	var len = len(arr)
	for i := 0; i < len; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("交换前=", arr)

	tmp := 0
	for i := 0; i < len/2; i++ {
		tmp = arr[len-1-i]
		arr[len-1-i] = arr[i]
		arr[i] = tmp
	}
	fmt.Println("交换后=", arr)

}
