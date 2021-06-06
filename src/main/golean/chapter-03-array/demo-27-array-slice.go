package main

import (
	"fmt"
)

//切片的使用
//1) 切片的英文是 slice
//2) 切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制。 3) 切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度 len(slice)都一样。 4) 切片的长度是可以变化的，因此切片是一个可以动态变化数组。
//5) 切片定义的基本语法:
//var 切片名 []类型
//var 切片名 []type = make([]type, len, [cap])
func main() {

	//1.直接引用数组，这个数组是实现存在的，程序员是可见的
	var array = [...]int{1, 2, 3, 4, 5}
	slice := array[1:3]
	fmt.Printf("chapter-03-array=%v,type=%T,地址=%p,容量=%v,元素个数=%v \n slice=%v,type=%T,地址=%p,容量=%v,元素个数=%v \n",
		array, array, &array, cap(array), len(array),
		slice, slice, &slice, cap(slice), len(slice))

	//2.通过make函数创建，make也会创建一个数组，是由切片底层维护，程序员不可见
	var mk = make([]int, 10, 12)
	mk[0] = 33
	mk[2] = 23

	fmt.Println("mk=", mk, "size=", len(mk), "cap=", cap(mk))

	var strArray = []string{"123", "hello", "rock"}
	fmt.Printf("arrayStr arr=%v, len=%v, cap=%v \n", strArray, len(strArray), cap(strArray))

	iteratorSlice()
}

//切片遍历
func iteratorSlice() {

	var array = [...]int{1, 2, 3, 4, 5}
	slice := array[1:3] //取值范围是[1,3) len=(end-start)
	//1.
	for i := range slice {
		fmt.Println("i=", slice[i])
	}
	//2.
	for i := 0; i < len(slice); i++ {
		fmt.Println("i=", slice[i])
	}
	for _, value := range slice {
		fmt.Println("value=", value)
	}

	slice2 := array[:]
	slice3 := slice2[:]
	//将切片追加到新的切片
	ints := append(slice3, slice2...)
	//切片
	fmt.Println("arr.len=", len(slice))
	fmt.Println("arr[2]=", slice[1])
	fmt.Printf("slice2=%v len=%v,cap=%v \n", slice2, len(slice2), cap(slice2))
	fmt.Printf("ints=%v len=%v,cap=%v \n", ints, len(ints), cap(ints))

	//切片拷贝，仅使用于切片类型
	var cp = make([]int, 10)
	copy(cp, slice3)
	fmt.Println("cp=", cp)

	str := "hello@1231231"

	strSlice := str[5:]
	//str[4] ='R'//string是不可变的，
	fmt.Printf("strSlice=%v \n", strSlice)

	//修改字符串：先将 string -> []byte / 或者 []rune -> 修改 -> 重写转成 string
	//
	arra2 := []byte(str) //不能转汉字。将string转成[]rune

	arra2[0] = 'D'
	str = string(arra2)
	fmt.Println("str2=", str)

	arra3 := []rune(str) //能转汉字。
	arra3[3] = '林'
	str = string(arra3)
	fmt.Println("str2=", str)
}
