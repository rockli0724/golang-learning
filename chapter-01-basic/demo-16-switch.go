package main

import (
	"fmt"
)

// switch 语句

func main() {
	/*	var num int = 10
		switch num {
		case 10:
			fmt.Println("ok1")
			fallthrough //switch穿透
		case 20:
			fmt.Println("ok2")
		case 30:
			fmt.Println("ok3")
		default:
			fmt.Println("nothing...")
		}*/

	//type switch
	/*	var x interface{}
		var y = false
		x = y
		switch i := x.(type) {
		case nil:
			fmt.Println("x的类型：%T", i)
		case int:
			fmt.Println("x的类型：int")
		case float64:
			fmt.Println("x的类型：float64")
		case chapter-02-func(int) float64:
			fmt.Println("x的类型：chapter-02-func(int) 类型")
		case bool, string:
			fmt.Println("x的类型：bool/string")
		default:
			fmt.Println("未知类型")
		}*/

	/*	//practice
		//1.使用switch把小写类型的char类型转成大写（键盘输入）
		//只转换 a,b,c,d,e,其他的输出"other"
		var char byte
		fmt.Println("请输入一个字符")
		fmt.Scanf("%c", &char)

		switch char {
		case 'a':
			fmt.Println("A")
		case 'b':
			fmt.Println("B")
		case 'c':
			fmt.Println("C")
		case 'd':
			fmt.Println("D")
		case 'e':
			fmt.Println("E")
		default:
			fmt.Println("other")
		}
	*/
	//2.输出学生成绩 >60分 输出“合格”，否则“不合格”
	/*	var course float64
		fmt.Println("请输入成绩")
		fmt.Scanln(&course)
		switch int(course / 60) {
		case 1:
			fmt.Println("ok")
		case 0:
			fmt.Println("not ok")
		default:
			fmt.Println("error...")
		}*/

	//根据用户指定月份
	//打印该月份所属的季节。 3、4、5春季 6、7、8夏季 9、10、11秋季 12、1、2冬季
	var month byte
	fmt.Println("请输入月份")
	fmt.Scanln(&month)

	switch month {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	}
}
