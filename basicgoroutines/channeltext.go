package basicgoroutines

import (
	"fmt"
	"time"
)

// Channelchuanlian : 串联channel
func Channelchuanlian() {
	naturals := make(chan int) //定义自然数channel
	squares := make(chan int)  //求平方channel

	// Counter
	go func() {
		for x := 0; ; x++ {
			time.Sleep(time.Duration(5) * time.Second)
			naturals <- x
			fmt.Println("counter", x)
		}
	}()

	// Squarer
	go func() {
		for {
			time.Sleep(time.Duration(5) * time.Second)
			x := <-naturals
			squares <- x * x
			fmt.Println("squarer", x)
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println("printer", <-squares)
	}
}

//带缓存的channel，节藕了生产者和消费者，但是消费次数，和生产次数及缓存的大小，
//不一致，会导致死锁的发生，缓存的大小也会影响效率。

//Directionchannel 单方向的channel
func Directionchannel() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)          //导致隐式转换
	go squarer(squares, naturals) //导致隐式转换
	printer(squares)
}
func counter(out chan<- int) {
	fmt.Println("producer")
	for x := 0; x < 100; x++ {
		time.Sleep(time.Duration(1) * time.Second)
		out <- x
	}
	fmt.Println("producer finish")
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	fmt.Println("middle")
	for v := range in { //range可以读取
		time.Sleep(time.Duration(5) * time.Second)
		out <- v * v
	}
	fmt.Println("middle finish")
	close(out)
}

func printer(in <-chan int) {
	fmt.Println("end")
	for v := range in {
		fmt.Println(v)
	}
	fmt.Println("end finish")
}

//Readkongchannel 读取关闭的channel
func Readkongchannel() {
	intchan := make(chan int)
	close(intchan)
	readchan, ok := <-intchan
	if !ok {
		fmt.Println(readchan, ok)
	}
}
