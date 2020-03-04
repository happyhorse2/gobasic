package basictype

import (
	"fmt"
	"strconv"
)

// SingleString : go语言float类型基础打印
func SingleString() {
	s := "hello world"
	fmt.Println(s[len(s)-1])
}

//SliceandSting 字符串与切片之间的关系和操作
func SliceandSting() {
	s := "hello world"
	fmt.Println(s[:5])
	fmt.Println(s[7:])
}

//Stringchange 查看字符串内存变化情况
func Stringchange() {
	s := "left foot"
	t := s
	fmt.Println(&s)
	fmt.Println(&t)
	s += ", right foot"
	fmt.Print(&s)
	fmt.Println(s)
	//s内存地址不变，s与t内存地址不同
	//字符串的值是不变的，但是可以重新个字符传变量复制.
	//s[0]= 'h',   报错了
}

//Stringtointerger 整形转字符串
func Stringtointerger() {
	x := 123
	y := fmt.Sprintf("%d", x) //返回字符串
	z := strconv.Itoa(x)
	k := strconv.Itoa(x)
	fmt.Println(&x, &y, &z, &k)
	fmt.Println(y, strconv.Itoa(x))
}
