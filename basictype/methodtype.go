package basictype

import (
	"fmt"
	"os"
)

var rmdirs []func()

// Persion : 人
type Persion struct {
	age  int
	name string
}

func nonamefunction() {
	dirs := []string{"one", "two", "three", "four"}
	for i := 0; i < len(dirs); i++ {
		os.MkdirAll(dirs[i], 0755) // OK
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dirs[i]) // NOTE: incorrect!    //用于回调这里就有问题了，i的地址不变，最后指向4
		})
	}
}

// Nonamefunctiontwo : go语言匿名函数
func Nonamefunctiontwo() {
	tempDirs := []string{"one", "two", "three", "four"}
	for _, dir := range tempDirs { //循环过程中，dir地址不变
		d := dir //重点关注，赋值完成了新的地址的创建，所有for循环中的代码应该使用d
		fmt.Printf("d=%v, *d=%p\n", d, &d)
		fmt.Printf("dir=%v, *dir=%p\n", dir, &dir) //这里没问题
		rmdirs = append(rmdirs, func() {
			fmt.Printf("dir=%v, *dir=%p\n", dir, &dir) //回调这里就会有问题
		})
	}

	fmt.Println("---")

	for _, f := range rmdirs {
		f()
	}
}

// Canshuchannel : channel作为实参
func Canshuchannel(input chan int) chan int {
	fmt.Println(&input)
	input <- 1
	Canshuchanneltwo(input)
	fmt.Println(len(input), "input length")
	fmt.Println(&input, "end")
	return input
}

// Canshuchanneltwo : channel作为实参
func Canshuchanneltwo(input chan int) chan int {
	fmt.Println(&input)
	input <- 2
	return input
}

//Canshustruct  结构体作为实参
func Canshustruct(person Persion) {
	fmt.Println(&person, "Canshustruct start")
	person.age = 10
	addName(person)
	fmt.Println(&person, "Canshustruct end")
}
func addName(person Persion) {
	fmt.Println(&person, "addName")
	person.name = "xiaoming"
}
