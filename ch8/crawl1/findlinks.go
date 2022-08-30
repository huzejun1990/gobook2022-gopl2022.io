// Also, it never terminates because the worklist is never closed.
// 此外，它永远不会终止，因为工作列表永远不会关闭。
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

func main() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	// 从命令行参数开始。
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	// 同时抓取网页。
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
