package basicgoroutines

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Closeosstandin : 关闭标准输入
func Closeosstandin() {
	time.Sleep(time.Duration(20) * time.Second)
	os.Stdin.Close()
}

// Copyinout : 复制标准输入到标准输出
func Copyinout() {
	if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
		fmt.Println("end a")
		log.Fatal(err)
	}
	fmt.Println("mustCopy")
}
