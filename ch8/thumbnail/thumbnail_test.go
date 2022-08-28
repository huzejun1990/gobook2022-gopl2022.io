// This file is just a place to put example code from the book.
// It does not actually run any code in gopl.io/ch8/thumbnail.

package thumbnail_test

import (
	"log"
	"os"
	"sync"

	"gopl2022.io/ch8/thumbnail"
)

//!+1
// makeThumbnails makes thumbnails of the specified files.
// makeThumbnails 制作指定文件的缩略图。
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

//!-1

//!+2
// NOTE: incorrect! (不确定)
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // NOTE: ignoring errors // 忽略错误
	}
}

//!-2

//!+3
// makeThumbnails3 makes thumbnails of the specified files in parallel.
// makeThumbnails3 并行生成指定文件的缩略图。
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE:ignoring errors // 排除错误
			ch <- struct{}{}
		}(f)
	}
	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

//!-3

//!-4
// makeThumbnails4 makes thumbnails for the specified files in parallel.
// makeThumbnails4 并行生成指定文件的缩略图。
// It returns an error if any step failed.
// 如果任何步骤失败，则返回错误。
func makeThumbnails64(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak! // goroutine 泄漏
		}
	}

	return nil
}

//!-4

//!-5
// makeThumbnails5 makes thumbnails for the specified files in parallel.
// makeThumbnails5 并行生成指定文件的缩略图。
// It returns the generated file names in an arbitrary order,
// 它以任意顺序返回生成的文件名，
// or an error if any step failed.
// 如果任何步骤失败，则返回错误。
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
	}
	return thumbfiles, nil

}

//!-5

//!+6
// makeThumbnails6 makes thumbnails for each file received from the channel.
// makeThumbnails6 为从频道接收到的每个文件制作缩略图。
// It returns the number of bytes occupied by the files it creates.
// 它返回它创建的文件占用的字节数。
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines	// 工作 goroutine 的数量
	for f := range filenames {
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

//!-6
