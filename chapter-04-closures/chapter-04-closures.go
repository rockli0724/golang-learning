package main

import "fmt"

func main() {

	nextInt := intSeq(2)
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq(2)
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
}

func intSeq(i int) func() int {
	return func() int {
		i++
		return i
	}
}
