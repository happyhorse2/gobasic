package basicgoroutines

import (
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err //通过channel确认所有的groutine都已执行完毕
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak!   //这样，判断error的话有个确定，如果err!=nil，则会发生groutine泄漏，
			//这样剩下的worker goroutine在向这个channel中发送值时，都会永远地阻塞下去，并且永远都不会退出。
			//这种情况叫做goroutine泄露(§8.4.4)，可能会导致整个程序卡住或者跑出out of memory的错误。
		}
	}
	return nil
}

//一个简单的channel方法
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

//通过
func makeThumbnails6(filenames []string) int64 { //filenames为从channel获取，长度不确定
	//sync.WaitGroup，这个计数器需要在goroutine启动时加1，退出时减1，
	//多个goroutine操作时，线程安全，goroutine推出时，会一直等待
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for _, f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
