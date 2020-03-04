package main

import (
	"fmt"
	basicfunc "gobasic/basicfunc"
	basicgoroutines "gobasic/basicgoroutines"
	basictype "gobasic/basictype"
)

func main() {
	//fmt.Println("hello world")
	//大写才能引用
	//basictype.PrintIntType()
	//指针类型
	// i := 5
	// basictype.Pointer(&i)
	// fmt.Println(i)

	// a := 6
	// basictype.Copyint(a)
	// fmt.Println(a)

	// basictype.TestNew()
	// basictype.Outputvarinitorder()

	// basictype.Printfloat()

	// basictype.Stringchange()

	// basictype.Stringtointerger()

	// basictype.Printshuzu()

	//basictype.Slicelencaptwo()

	basictype.Slicelencap()
	//basictype.Slicelencapnil()

	//basictype.Mapnil()

	//basictype.Printstructtype()

	//testfunc()
	//testbasictype()
	//测试协程
	// go testgoroutine()
	// time.Sleep(time.Duration(20) * time.Second)
	// fmt.Print("end")

	//testgoroutineclock()
	//go basicgoroutines.Closeosstandin()
	//basicgoroutines.Copyinout()
	//testnetchannel()
	//testchannelchuanlian()

	//testnoname()

	//testcrawl()

	//testioselect()
	//testDirents()
	//testcanshuchannel()
	//testdirectionchannel()
	//testfilesize()

	//testgroutineout()

}

func testfunc() {
	basicfunc.Httpfunerror()
}

func testbasictype() {
	basictype.Testpointerandinstanceversiontwo()
}

func testgoroutine() {
	basicgoroutines.Twogoroutines()
}

func testgoroutineclock() {
	//go basicgoroutines.Clockserver()
	//time.Sleep(time.Duration(1) * time.Second)

	//go basicgoroutines.Clockclinentone()
	basicgoroutines.Clockclinentonev2()
}

func testnetchannel() {
	basicgoroutines.Netchatv3()
}

func testchannelchuanlian() {
	basicgoroutines.Channelchuanlian()
}
func testnoname() {
	basictype.Nonamefunctiontwo()
}

func testcrawl() {
	url := "https://github.com/csunny/argo"
	basicgoroutines.Crawl(url)
}

func testioselect() {
	//测试基于select的io多路复用
	//basicgoroutines.Timerecord()
	basicgoroutines.Timerecordv2()
}

func testDirents() {
	//并发遍历目录
	dir := "/home/mayansong/java/Projectone/testgroup"
	filesize := make(chan int64, 200)
	fmt.Println(&filesize, "filesizeinitial")
	//result := basicgoroutines.Dirents(dir) //遍历目录路径
	result := basicgoroutines.WalkDir(dir, filesize) //递归遍历目录，返回大小
	fmt.Println(&result, "result")
	// fmt.Println(result)
	// for _, file := range result {
	// 	fmt.Println(file.Name())
	// }
	for {
		if len(result) > 0 {
			dirsize := <-result
			fmt.Println(dirsize)
		} else {
			fmt.Println("break")
			break
		}
	}
}

func testcanshuchannel() {
	input := make(chan int, 10)
	fmt.Println(&input)
	basictype.Canshuchannel(input)
}

func testcanshustruct() {
	//person := basictype.Persion{}
}

func testdirectionchannel() {
	basicgoroutines.Directionchannel()
}
func testfilesize() {
	//basicgoroutines.ComputeFilesize()

	//可显示进度
	basicgoroutines.WithrecordDiversedir()
	//basicgoroutines.Readkongchannel()
}

func testgroutineout() {
	basicgoroutines.TestchannelClose()
	//fmt.Println(result)
}

func testshican() {
	basictype.Testchangepersonage()
}
