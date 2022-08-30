// Crawl2 crawls web links starting with the command-line arguments.
// Crawl2 从命令行参数开始抓取网页链接。
// This version uses a buffered channel as a counting semaphore
// 此版本使用缓冲通道作为计数信号量
// to limit the number of concurrent calls to links.Extract.
// 限制对 links.Extract 的并发调用数。
package main

import (
	"fmt"
	"gopl2022.io/ch5/links"
	"log"
	"os"
)

//!+sema
// tokens is a counting semaphore used to
// token 是一个计数信号量，用于
// enforce a limit of 20 concurrent requests.
// 强制限制 20 个并发请求。
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} //acquire a token	//获得token
	list, err := links.Extract(url)
	<-tokens // release the token	// 生成 token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist // 待发送到工作列表的数量

	// Start with the command-line arguments.
	// 从命令行参数开始。
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	// 同时抓取网页。
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

//!-
