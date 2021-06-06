package main

import "fmt"

//数组的使用
//1. 数组的地址可以通过数组名来获取 &intArr
//2. 数组的第一个元素的地址，就是数组的首地址
//3. 数组的各个元素的地址间隔是依据数组的类型决定，比如 int64 -> 8  int32->4...
//4. go的数组属于值类型，在默认情况下是值传递，因此会值拷贝。数组之间数据不会相互影响
//5. 如果想在其他函数中修改原来数组的值，可以使用引用传递（指针方式）
//6. 长度是数组类型的一部分，在传递函数参数时 需要考虑数组的长度
func main() {
	//定长数组
	var arr = [6]float64{1, 2, 3, 4, 5, 6}
	arr2 := [6]float64{1, 2, 3, 4, 5, 6}
	fmt.Println("arry1=", arr, "array2=", arr2)

	//不定长数组
	var arr_dynamic = [...]float64{1, 2, 3, 4, 5, 6}
	arr_dynamic2 := [...]float64{1, 2, 3, 4, 5, 6}
	fmt.Println("arr_dynamic=", arr_dynamic, "arr_dynamic[3]=", arr_dynamic[5], "arr_dynamic2=", arr_dynamic2)

	//数组遍历
	for index, value := range arr_dynamic {
		fmt.Println("index=", index, "value=", value)
	}
	fmt.Println("-----------------------okokok-----------")
	for _, value := range arr_dynamic {
		fmt.Println("value=", value)
	}

	//
	var arr_empty = [4]int{1, 2, 3, 4}
	fmt.Println("arr_empty=", arr_empty)
	fmt.Printf("arr_empty地址=%p ,arr[0]=%p, arr[1]=%p, arr[2]=%p, arr[3]=%p \n",
		&arr_empty, &arr_empty[0], &arr_empty[1], &arr_empty[2], &arr_empty[3])

	updateArray(arr_empty)
	fmt.Println("arr_empty-update----------")
	fmt.Println("arr_empty-update=", arr_empty)
	fmt.Printf("arr_empty-update地址=%p ,arr[0]=%p, arr[1]=%p, arr[2]=%p, arr[3]=%p",
		&arr_empty, &arr_empty[0], &arr_empty[1], &arr_empty[2], &arr_empty[3])

	updateNewArray(&arr_empty)
	fmt.Println("arr_empty-new-update----------")
	fmt.Println("arr_empty-new-update=", arr_empty)
	fmt.Printf("arr_empty-new-update地址=%p ,arr[0]=%p, arr[1]=%p, arr[2]=%p, arr[3]=%p",
		&arr_empty, &arr_empty[0], &arr_empty[1], &arr_empty[2], &arr_empty[3])

}

//数组修改，需要声明数组长度（否则编译错误，会被认为是切片）.对原数组数据不产生影响
func updateArray(arr [4]int) {
	arr[0] = 999
}

//数组修改，修改内存数据
func updateNewArray(arr *[4]int) {
	(*arr)[0] = 999
}
