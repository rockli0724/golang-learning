package main

import "fmt"

//递归函数
func main() {
	//recursion(4)
	println(funcFbnc(2))
	println(funcFbnc(3))
	println(funcFbnc(4))
	println(funcFbnc(5))
	println("--------")
	println(practice(1))
	println(practice(2))
	println(practice(3))
	println(practice(4))
	println(practice(5))
	println("-----chapter-02-func param---")

	s := getSum
	fmt.Printf("s的数据类型%T, sum的类型是%T \n", s, getSum)
	i := s(1, 2) //等价于sum(i int,i int)函数
	fmt.Println("i=", i)

	i2 := myFunc(getSum, 20, 10)
	fmt.Println("i2=", i2)

	type myInt int
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1)
	fmt.Println("num1=", num1, "num2=", num2)
	fmt.Println("-------swap------")

	var a int = 10
	var b int = 200
	swapValue(&a, &b)
	fmt.Println("a=", a, "b=", b)

}

func recursion(n int) {
	if n > 2 {
		n--
		recursion(n)
	} else {
		fmt.Println("n=", n)
	}
}

//斐波那契数列
func funcFbnc(n int) int {
	if n == 1 || n == 2 {
		return n
	} else {
		return funcFbnc(n-1) + funcFbnc(n-2)
	}
}

//已知 f(1)=3; f(n) = 2*f(n-1)+1;
func practice(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*practice(n-1) + 1
	}
}

//函数当入参
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFunc(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//多个入参[可变参数，需要放在形参后面]
func sums(num int, args ...int) int {
	sum := num
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func swapValue(n1 *int, n2 *int) {
	tem := *n1
	*n1 = *n2
	*n2 = tem
}
