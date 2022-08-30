// The du2 command computes the disk usage of the files in a directory.
// du2 命令计算目录中文件的磁盘使用情况。
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// The du2 variant uses select and a time.Ticker
// du2 变体使用 select 和 time.Ticker
// to print the totals periodically if -v is set.
// 如果设置了 -v，则定期打印总数。

//!+
var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// ...start background goroutine...	//后台启动 goroutine

	//!-
	// Determine the initial directories.	// 确定初始目录。
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.	// 遍历文件树。
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	//!+
	// Print the results periodically.	// 定期打印结果。
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed	// 文件大小已关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals	// 最终总数

}

//!-
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// walkDir 递归遍历以 dir 为根的文件树
// and sends the size of each found file on fileSizes.
// 并在 fileSizes 上发送每个找到的文件的大小。
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
// dirents 返回目录 dir 的条目。
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries

}
