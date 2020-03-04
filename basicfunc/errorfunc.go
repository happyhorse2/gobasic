package basicfunc

import (
	"fmt"
	"net/http"
)

//Httpfunerror 错误通常作为，函数的额外返回值返回，通常是函数的最后一个返回值
func Httpfunerror() (error, error) { //返回多个指
	url := "hello"
	resp, err := http.Get(url)
	fmt.Println(resp, err)
	fmt.Println("aaaaaaaaa")
	if err != nil {
		fmt.Println(err == nil, "bbbbbbbb")
		return nil, err
	}
	fmt.Println(err == nil, "ccccccccccccc")
	return nil, err
}

//文件读取错误EOF
