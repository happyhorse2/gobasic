package basictype

import (
	"fmt"
)

//Pointer 指针操作
func Pointer(p *int) int {
	*p++
	return *p
}

//Copyint 普通类型,值复制
func Copyint(p int) int {
	p++
	return p
}

//TestNew 测试new函数创建变量
func TestNew() {
	/*new函数返回初始化变量的指针*/
	p := new(int)
	fmt.Println(p)
	fmt.Println(*p)
}

// Person Person结构体
type Person struct {
	age  int
	name string
}

// Pointerandinstance 隐式解引用
func Pointerandinstance() {
	var studenta Person
	fmt.Println(studenta.age)
	var studentapoint *Person
	studentapoint = &studenta
	fmt.Println(studentapoint.age) //
}

// pointerversiontwo 隐式解引用
func (ptr *Person) pointerversiontwo() {
	ptr.age = 20
}

// instanceverssion 隐式解引用
func (ptr Person) instanceverssion() {
	ptr.age = 40
}

// pointerversionthree 隐式解引用
func (ptr *Person) pointerversionthree() {
	ptr.age = 30
}

// instanceverssiontwo 隐式解引用
func (ptr Person) instanceverssiontwo() {
	ptr.age = 50
}

// Testpointerandinstanceversiontwo 隐式解引用
func Testpointerandinstanceversiontwo() {
	var teachera Person
	teachera.pointerversiontwo()
	fmt.Println(teachera.age)   //20
	teachera.instanceverssion() //会新生成一个结构体，执行里面的程序，对原结构体无影响
	fmt.Println(teachera.age)   //20

	//Person{1, "xiaoming"}.pointerversiontwo() //接收器是指针，需要用变量调用，否则找不到指针
	(&teachera).pointerversionthree() //30
	fmt.Println(teachera.age)
	(&teachera).instanceverssiontwo() //30
	fmt.Println(teachera.age)         //会新生成一个结构体，执行里面的程序，对原结构体无影响

	// 方法接收器是对象，不是指针，运行程序对调用对象无影响
}

// Testchangepersonage 测试实参
func Testchangepersonage() {
	var a Persion
	changeperssonage(a)
	fmt.Println(a)
}

func changeperssonage(per Persion) {
	per.age = 90
}
