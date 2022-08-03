// Fetch saves the contents of a URL into a local file.
// Fetch 将 URL 的内容保存到本地文件中。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

//!+
// Fetch downloads the URL and returns the // Fetch 下载 URL 并返回
// name and length of the local file. // 本地文件的名称和长度。
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	//关闭文件，但更喜欢复制中的错误（如果有）。
	if closeErr := f.Close(); err != nil {
		err = closeErr
	}
	return local, n, err
}

//!-

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}
