package main

import "fmt"

// 多重for循环
//打印金字塔

func main() {

	//打印半个金字塔
	for i := 1; i < 6; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("----------")

	//打印全部金字塔 2n-1

	/*	  *				1层 1个* 规律： 2*层数-1 空格规律： 总层数-当前层数
		 ***
		*****
	   *******
	*/
	var deep int = 4
	//层数
	for i := 1; i <= deep; i++ {
		for space := 1; space <= deep-i; space++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("----------")
	//打印空心金字塔
	/*	  *				1层 1个* 规律： 2*层数-1 空格规律： 总层数-当前层数
		 * *
		*   *
	   *******
	*/

	//层数
	for i := 1; i <= 4; i++ {
		for space := 1; space <= 4-i; space++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == 4 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println("------99乘法表----")

	var num int = 9
	for i := 1; i <= num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}
