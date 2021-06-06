package main

import "fmt"

// break

func main() {

	for i := 0; i < 4; i++ {
	label1:
		for j := 0; j < 10; j++ {
			if j == 3 {
				break label1
			}
			fmt.Println("OK")
		}
		fmt.Println("------------")
	}
}
