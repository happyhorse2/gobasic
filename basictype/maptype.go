package basictype

import "fmt"

// Mapnil : 映射的初始化，空指针
func Mapnil() {
	var ages map[string]int         //默认初始化ages为nil，没有开辟空间
	fmt.Print(ages == nil)          //true
	agesmap := make(map[string]int) //空map， 开辟了空间
	fmt.Print(agesmap == nil)       //false
}
