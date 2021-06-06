package main

import (
	"fmt"
	"learn-go/src/main/golean/chapter-02-func/pkg/util"
)

//运算函数计算
func main() {

	result := util.Operation(2.3, 12.4, '/')
	fmt.Printf("result=%v \n", result)

	sum, sub := util.SumAndSub(3, 5)
	fmt.Println("getSum=", sum, "sub=", sub)
}
