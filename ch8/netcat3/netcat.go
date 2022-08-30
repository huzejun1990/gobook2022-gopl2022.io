// Netcat is a simple read/write client for TCP servers.
// Netcat 是一个简单的 TCP 服务器读/写客户端。
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOET : ignoring errors //排除错误
		log.Println("done")
		done <- struct{}{} // signal the main goroutine // 向主 goroutine 发出信号
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done //	wait for background goroutine to finish //等待后台 goroutine 完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
