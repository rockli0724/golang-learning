package main

import "fmt"

//结构体默认值拷贝
func main() {

	//声明方式1 直接声明
	var person0 Person
	person0.Name = "liln3"
	fmt.Println("person0=", person0)

	//声明方式2
	person1 := Person{"lilin", 29}
	person1.Age = 44
	fmt.Println("person1=", person1)

	//声明方式3
	var person2 *Person = new(Person)
	//go 的设计者，为了程序员使用方便，底层会将person2.Name="lilin" 进行处理
	// 会再person2 上取值运算(*person2).Name="lilin"
	(*person2).Name = "lilin2"
	person2.Age = 33
	fmt.Println("person2=", *person2)

	//声明方式4 {}
	//因为person4是一个指针，因此标准的访问字段的方式
	//(*person4).Name = "scott"
	//go 的设计者，为了程序员使用方便，底层会将person2.Name="lilin" 进行处理
	// 会再person2 上取值运算(*person2).Name="lilin"
	var person4 *Person = &Person{"mary", 22}

	////标准写法
	//(*person4).Age = 12
	////也可以这样
	//person4.Name = "person4"
	fmt.Println("person2=", *person4)

	//练习值拷贝

	var p1 Person
	p1.Name = "小明"
	p1.Age = 22
	var p2 *Person = &p1

	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)

	p2.Name = "tom"

	fmt.Println("p1.name=", p1.Name, "p2.name=", p2.Name)

	var a A
	var b B
	a = A(b) //结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段(名字、个数和类 型)

	fmt.Println(a, b)

}

type Person struct {
	Name string
	Age  int
}

type A struct {
	Num int
}
type B struct {
	Num int
}
