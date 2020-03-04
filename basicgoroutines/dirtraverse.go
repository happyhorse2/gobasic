package basicgoroutines

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Traveersedir : 两个协程
func Traveersedir() {
	//并发遍历目录

}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range Dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func walkDirv2(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range Dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDirv2(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20) //控制并发遍历进度

// Dirents returns the entries of directory dir.
func Dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: //acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() //release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

// WalkDir returns the entries of directory dir.
func WalkDir(dir string, filesize chan int64) chan int64 {
	fmt.Println("digui")
	fmt.Println(&filesize, "filesize", len(filesize))
	for _, entry := range Dirents(dir) {
		if entry.IsDir() {
			fmt.Println("is dir")
			fmt.Println(dir, entry.Name())
			subdir := filepath.Join(dir, entry.Name())
			fmt.Println(subdir)
			WalkDir(subdir, filesize)
			fmt.Println("error  or not")
		} else {
			fmt.Println("is size")
			fmt.Println(entry.Name())
			filesize <- entry.Size()
		}
	}
	fmt.Println(&filesize, "filesize", len(filesize))
	return filesize
}

// ComputeFilesize returns the entire size of directory dir.
func ComputeFilesize() {
	flag.Parse()
	roots := flag.Args()
	fmt.Println(roots)
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			fmt.Println("di gui")
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results.
	var nfiles, nbytes int64
	fmt.Println(len(fileSizes), "zhu groutine")
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes))
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// WithrecordDiversedir 遍历目录，并通过计时显示进度
func WithrecordDiversedir() {
	flag.Parse()
	roots := flag.Args()
	fmt.Println(roots)
	if len(roots) == 0 {
		roots = []string{"."}
	}
	var n sync.WaitGroup
	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			n.Add(1)
			go walkDirv2(root, &n, fileSizes)
		}
	}()

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			//to allow existing goroutines to finish
			for range fileSizes { //fileSizes关闭时，for循环会自动结束
				//Do nothing
			}
			return
		case size, ok := <-fileSizes: //Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据；如果channel中已经没有数据的话将产生一个零值的数据。
			fmt.Println("size", size, "ok", ok)
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-ticker.C:
			printDiskUsage(nfiles, nbytes)
		}
	}
}
