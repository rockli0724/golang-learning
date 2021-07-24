package main

import (
	"fmt"
)

// 多重for循环

func main() {

	//多重循环
	//1.统计 3 个班成绩情况，每个班有 5 名同学，求出各个班的平均分和所有班级的平均分[学生的成 绩从键盘输入
	var classNum int = 2
	var stuNum int = 5
	var totalSum float64 = 0.0

	for j := 1; j <= classNum; j++ {
		sum := 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("清输入第%d个班,第%d个学生的成绩", j, i)
			fmt.Scanln(&score)
			//累计总分
			sum += score
		}
		fmt.Printf("第%d个班的平均成绩是%v \n", j, sum/float64(stuNum))
		totalSum += sum
	}
	fmt.Printf("各个班的总成绩是%v,平均每个班的成绩是%v", totalSum, totalSum/float64(classNum*stuNum))

}
