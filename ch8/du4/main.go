// The du4 command computes the disk usage of the files in a directory.
// du4 命令计算目录中文件的磁盘使用情况。

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// The du4 variant includes cancellation:
// du4 变体包括取消：
// it terminates quickly when the user hits return.
// 当用户点击返回时它会快速终止。

//!+1
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//!-1

func main() {
	// Determine the initial directories.
	// 确定初始目录
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+2
	// Cancel traversal when input is detected.
	// 检测到输入时取消遍历
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte	// 读取单个字节
		close(done)
	}()
	//!-2

	// Traverse each root of the file tree in parallel.
	// 并行遍历文件的每个根
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically
	// 定期打印结果
	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	//!+3
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			// 排空文件大小以允许现有的 goroutine 完成。
			for range fileSizes {
				// Do nothing.
			}
			return
		case size, ok := <-fileSizes:
			// ...
			//!-3
			if !ok {
				break loop //fileSizes was closed // 文件大小已经关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals //最终总数
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// walkDir 递归遍历以 dir 为根的文件树
// and sends the size of each found file on fileSizes.
// 并在 fileSizes 上发送每个找到的文件的大小
//!+4
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		// ...
		//!-4
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
		//!+4
	}
}

//!-4
//concurrency-limiting counting semaphore
// 并发限制计数信号量
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
// dirents 返回目录的 dir 条目
//!+5
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token 	//获取token
	case <-done:
		return nil // cancelled	//取消
	}
	defer func() { <-sema }() // release token //重新生成token

	// ... read directory ...	//读取目录
	//!-5

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries //读取所有条目
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
		// 不返回 Readdir 可能返回的部分结果
	}
	return entries
}
