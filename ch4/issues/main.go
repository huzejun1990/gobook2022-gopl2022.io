// Issues prints a table of GitHub issues matching the search terms.
// 问题打印与搜索词匹配的 GitHub 问题表。
package main

import (
	"fmt"
	"gopl2022.io/ch4/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
