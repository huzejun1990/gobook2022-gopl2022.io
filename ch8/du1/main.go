// The du1 command computes the disk usage of the files in a directory.
// du1 命令计算目录中文件的磁盘使用情况。
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//!+main
func main() {
	// Determine the initial directories.
	// 确定初始目录。
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	// 遍历文件树。
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results	// 打印结果
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nfiles)/1e9)
}

//!-main

//!+walkDir
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
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

//!-walkDir

// The du1 variant uses two goroutines and
// du1 变体使用两个 goroutine 和
// prints the total after every file is found.
// 在找到每个文件后打印总数。
