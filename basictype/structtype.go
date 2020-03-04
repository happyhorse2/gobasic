package basictype

import (
	"fmt"
)

// Point : 点结构
type Point struct {
	X, Y int
}

// Printstructtype :打印指针类型
func Printstructtype() {
	var point Point //默认初始化
	fmt.Print(point.X)
}

func init() {
	fmt.Println("init in Printstructtype.go ")
}
