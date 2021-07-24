package main

import "fmt"

type Box struct {
	length float64
	width  float64
	hight  float64
}

func (b *Box) getVum() float64 {
	return b.length * b.width * b.hight
}

type Student struct {
	name   string  `json:"name,omitempty"`
	gender string  `json:"gender,omitempty"`
	age    int     `json:"age,omitempty"`
	id     int     `json:"id,omitempty"`
	score  float64 `json:"score,omitempty"`
}

func (student Student) say() string {
	sprintf := fmt.Sprintf("student 的信息 name=[%v] gender=[%v], age=[%v] id=[%v] score=[%v]", student.name, student.gender, student.age, student.id, student.score)
	return sprintf
}

type Calcuator struct {
	Num1 float64
	Num2 float64
	Name string
}

func (c *Calcuator) getSum() float64 {
	return c.Num1 + c.Num2
}

func (c *Calcuator) getSumb() float64 {
	return c.Num1 - c.Num2
}

func (c *Calcuator) getResult(operation byte) float64 {
	result := 0.0
	switch operation {
	case '+':
		result = c.Num1 + c.Num2
	case '-':
		result = c.Num1 - c.Num2
	case '*':
		result = c.Num1 * c.Num2
	case '/':
		result = c.Num1 / c.Num2
	default:
		fmt.Println("输入有误")
		result = 0.0
	}
	return result
}

func (c Calcuator) chengfa(num int) {
	fmt.Println("------99乘法表----")
	for i := 1; i <= num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v x %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}

//函数

func test01(c Calcuator) {
	fmt.Println(c.Name)
}
func test02(c *Calcuator) {
	fmt.Println(c.Name)
}

//方法
func (c Calcuator) test03() {
	fmt.Println("--->test03()-before=", c.Name)
	c.Name = "test03-jack"
	fmt.Println("test03()=", c.Name)
}
func (c *Calcuator) test04() {
	fmt.Println("--->test04-before=", c.Name)
	c.Name = "test04-toming"
	fmt.Println("test04=", c.Name)
}

//对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
//总结:
//1) 不管调用形式如何，真正决定是值拷贝还是地址拷贝，看这个方法是和哪个类型绑定.
//2) 如果是和值类型，比如 (p Person) , 则是值拷贝， 如果和指针类型，比如是 (p *Person) 则是地址拷贝
func main() {
	var ca Calcuator
	ca.chengfa(5)

	calcuator := Calcuator{
		0, 0, "tom",
	}
	test01(calcuator)
	fmt.Println("main.test01 after=", calcuator.Name)
	test02(&calcuator)
	fmt.Println("main.test02 after=", calcuator.Name)

	calcuator.test03()
	fmt.Println("calcuator.test03 after=", calcuator.Name)
	calcuator.test04()
	fmt.Println("calcuator.test04 after=", calcuator.Name)

	//创建结构体1，可指定字段的值，这样不依赖字段顺序
	student := Student{
		name:   "lilin",
		gender: "男",
		age:    0,
		id:     10,
		score:  8.9,
	}
	fmt.Println("studentInfo=", student.say())

	box := Box{
		length: 10.2,
		width:  10.2,
		hight:  3.0,
	}
	fmt.Println(box.getVum())

	//创建结构体2，返回结构体的指针类型
	c := &Calcuator{
		Num1: 222,
		Num2: 0,
		Name: "heihhe",
	}
	fmt.Println("test8=", *c)

}
