package main

import "fmt"

//结构体方法
//test方法和Person类型绑定
//test 方法只能通过 Person 类型的变量来调用，而不能直接调用，也不能使用其它类型变量来调用
//方法的调用和传参机制原理:(重要!)
//方法的调用和传参机制和函数基本一样，不一样的地方是方法调用时，会将调用方法的变量，当做实参也传递给方法。
//1) 结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式
//2) 如程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理
func main() {
	var ps Persons
	ps.Name = "tom"
	fmt.Println("person=", ps)
	ps.test()
	ps.caculater()
	ps.caculater2(4)
	ps.speak()
	ps.speakPoint()
	sum := ps.getSum(1, 4)
	fmt.Println("getSum()", sum)
	fmt.Println("person=", ps)

	var it integer = 20
	it.print()
	it.change()
	fmt.Println("i2=", it)

	dog := Dog{
		Name: "Dogs",
		Age:  12,
	}
	fmt.Println(dog.say())

}

type Persons struct {
	Name string `json:"name,omitempty"`
}

func (p Persons) test() {
	fmt.Println("func.37.test()")
}

func (p Persons) speak() {
	fmt.Println("speak()方法")
}
func (p *Persons) speakPoint() {
	p.Name = "speakPoint"
	fmt.Println("speakPoint()方法")
}

//计算 1+2+3+....+1000
func (p Persons) caculater() {
	result := 0
	for i := 0; i <= 1000; i++ {
		result += i
	}
	fmt.Println(p.Name, "caculater() 计算1+2+3+....+1000=", result)
}

//计算 1+2+3+....+1000
func (p Persons) caculater2(n int) {
	result := 0
	for i := 0; i <= n; i++ {
		result += i
	}
	fmt.Println(p.Name, "caculater() 计算1+2+3+....+n=", result)
}

//有返回值的结构体方法
func (p Persons) getSum(n1, n2 int) int {
	return n1 + n2
}

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}
func (i *integer) change() {
	*i = *i + 1
}

type Dog struct {
	Name   string
	Age    int
	Weight float64
}

func (dog Dog) say() string {
	sprintf := fmt.Sprintf("Name=[%v], Age=[%v], Weight=[%v]", dog.Name, dog.Age, dog.Weight)
	return sprintf
}

func (dog Dog) String() string {
	sprintf := fmt.Sprintf("Name=[%v], Age=[%v], Weight=[%v]", dog.Name, dog.Age, dog.Weight)
	return sprintf
}
