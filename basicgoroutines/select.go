package basicgoroutines

import (
	"fmt"
	"os"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// Timerecord : 计时器开始工作
func Timerecord() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		i := <-tick
		fmt.Println(i, "time tick") //没1秒触发一次
	}
}

//Timerecordv2 : 计时器开始工作,使用newticker
func Timerecordv2() {
	fmt.Println("Commencing countdown.")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		i := <-ticker.C
		fmt.Println(i, "time tick") //没1秒触发一次
	}
	fmt.Println("launch")
}

// Timeafter : 10秒后开始，运行发射
func Timeafter() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	fmt.Println("launch")
}
