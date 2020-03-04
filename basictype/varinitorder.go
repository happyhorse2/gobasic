package basictype

import (
	"fmt"
)

//不同的变量初始化顺序不同，局部变量初始化顺序，从上到下，从左到右，全局变量，根据依赖进行初始化
var a int = b + 1
var b int = 1

// Outputvarinitorder 输出变量初始化顺序
func Outputvarinitorder() {
	var (
		d int = 1
		c int = d + 1
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
