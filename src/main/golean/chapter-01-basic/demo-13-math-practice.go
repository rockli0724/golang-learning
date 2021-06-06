package main

import "fmt"

func main() {
	//加入还有97天放假，问还有XX个星期XX天
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d个星期零%d天 \n", week, day)

	//定义一个变量保存化华氏温度，话是问题--》摄氏温度转换公式为
	//5/9*(华氏温度-100)，求华氏温度对应的摄氏温度
	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("%v华氏度对应摄氏度=%v \n", huashi, sheshi)

	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age <= 40 {
		fmt.Println("ok2")
	}
}
