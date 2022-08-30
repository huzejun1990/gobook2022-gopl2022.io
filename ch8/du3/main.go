// The du3 command computes the disk usage of the files in a directory.
// du3 命令计算目录中文件的磁盘使用率。
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

//!+
func main() {
	// ...determine roots...	//确定根

	//!-
	flag.Parse()

	// Determine the initial directories.
	// 确定初始目录。
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	// 并行遍历文件树的每个根。
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
	//!-

	// Print the results periodically.
	// 定期打印结果
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed //关闭文件大小
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes) // final totals //最终条数
}

//!-
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nbytes, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// walkDir 递归遍历以 dir 为根的文件树
// and sends the size of each found file on fileSizes.
// 并在 fileSizes 上发送每个找到的文件的大小。
//!+walkDir
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
// sema 是一个计数信号量，用于限制 dirent 中的并发性。
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
//dirents 返回目录 dir 的条目。
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token	// 获取 token
	defer func() { <-sema }() // release token  // 重新生成token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

//!-
