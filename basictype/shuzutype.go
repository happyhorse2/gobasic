package basictype

import (
	"fmt"
)

// Printshuzu : go语言数组
func Printshuzu() {
	var a [3]int
	for i, v := range a {
		fmt.Println(i, v)
	}
}
