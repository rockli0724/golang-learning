package util

import "fmt"

var va = test

func test() {
	fmt.Println("uril.global function")
}

func Operation(n1 float64, n2 float64, operation byte) float64 {
	var rs float64
	switch operation {
	case '+':
		rs = n1 + n2
	case '-':
		rs = n1 - n2
	case '*':
		rs = n1 * n2
	case '/':
		rs = n1 / n2
	default:
		fmt.Println("operation error ...")
	}
	return rs
}

func init() {
	fmt.Println("util.go init()")

}

//求和、差
func SumAndSub(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}
