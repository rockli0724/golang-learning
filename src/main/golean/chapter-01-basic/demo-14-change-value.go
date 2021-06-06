package main

import "fmt"

// a = 10 b =20 将a b 值进行交换但不允许复中间值
func main() {
	var a int = 10
	var b int = 20

	a = a + b
	b = a - b // b = a + b - b ==> b = a
	a = a - b // a = a + b - a ==> a = b

	fmt.Printf("a=%v , b=%v \n", a, b)

	//
	c := 100
	fmt.Println("c的地址=", &c)

	var ptr *int = &c
	fmt.Println("ptr 指向的值=", *ptr)

}
