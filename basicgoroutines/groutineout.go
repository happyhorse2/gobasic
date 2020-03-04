package basicgoroutines

import (
	"fmt"
	"os"
	"time"
)

//var done = make(chan struct{})

// Cancelled :  推出
func Cancelled() bool {
	result, ok := <-done
	fmt.Println(result, ok) //关闭的channel读出的是默认的零值和false
	return true
}

// TestchannelClose :  推出
func TestchannelClose() {
	go func() {
		os.Stdin.Read(make([]byte, 1)) //从标准输入读取一个字符，执行goroutine退出
		close(done)
	}()
	for {
		if Cancelled() {
			time.Sleep(time.Duration(2) * time.Second)
		} else {
			time.Sleep(time.Duration(2) * time.Second)
		}
	}
}
