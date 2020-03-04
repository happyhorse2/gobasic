package basicgoroutines

import (
	"fmt"
	"time"
)

// 当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用go语句来创建

// Twogoroutines : 两个协程
func Twogoroutines() {
	go spinner(10 * time.Millisecond)
	const n = 2
	fibN := fib(n)
	fmt.Printf("Fibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	fmt.Printf("2222222222")
	// for {
	// 	for _, r := range `-\|/` {
	// 		fmt.Printf("%c", r)
	// 		time.Sleep(delay)
	// 	}
	// }
	time.Sleep(time.Duration(60) * time.Second)
	fmt.Printf("3333333333") // 主goroutine(运行main函数的),运行结束后，所有子goroutine（在运行中）也结束，
	//（非master）父goroutine中开启的子goroutine,结束不先后影响
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
