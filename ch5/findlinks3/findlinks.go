// Findlinks3 从命令行中的 URL 开始爬网。
package main

import (
	"fmt"
	"gopl2022.io/ch5/links"
	"log"
	"os"
)

// widthFirst 为工作列表中的每个项目调用 f。
// f 返回的任何项目都将添加到工作列表中。
// 每个项目最多调用一次 f。
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

func main() {
	// 以广度优先抓取网络，
	// 从命令行参数开始。
	breadthFirst(crawl, os.Args[1:])
}
