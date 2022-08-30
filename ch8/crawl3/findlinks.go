// Crawl3 从命令行参数开始抓取网页链接。
//
// 此版本使用有界并行。
// 为简单起见，它没有解决终止问题。
package main

import (
	"fmt"
	"gopl2022.io/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	worklist := make(chan []string)  //URL 列表，可能有重复
	unseenLinks := make(chan string) //去重的 URL

	//将命令行参数添加到工作列表。
	go func() { worklist <- os.Args[1:] }()

	// 创建 20 个爬虫 goroutine 来获取每个看不见的链接。
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// 主 goroutine 去重工作列表项
	// 并将看不见的发送给爬虫。
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}

//!-
