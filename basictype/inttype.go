package basictype

import (
	"fmt"
)

// PrintIntType : go语言int类型基础打印
func PrintIntType() {
	var uIntType uint8 = 255
	fmt.Println(uIntType*uIntType, uIntType+1, uIntType+uIntType)

}

// Max :比较大小
func Max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//不管它们的具体大小，int、uint和uintptr是不同类型的兄弟类型。其中int和int32也是不同的类型，即使int的大小也是32bit，在需要将int当作int32类型的地方需要一个显式的类型转换操作，反之亦然。
