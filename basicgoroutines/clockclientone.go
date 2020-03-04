package basicgoroutines

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Clockclinentone : 连接客户端
func Clockclinentone() {
	fmt.Println("client start")
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
	fmt.Println("client end")
}

// Clockclinentonev2 : 连接客户端
func Clockclinentonev2() {
	fmt.Println("client start")
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
	fmt.Println("client end")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
	fmt.Println("mustCopy")
}

// Netchatv3 : 通过管道
func Netchatv3() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		fmt.Println("before channel")
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	fmt.Println("middle channel")
	conn.Close()
	<-done // wait for background goroutine to finish，无缓存channel可以
	fmt.Println("after channel")

}
