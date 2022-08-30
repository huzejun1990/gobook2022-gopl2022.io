// Netcat is a simple read/write client for TCP servers.
// Netcat 是一个简单的 TCP 服务器读/写客户端。
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCoyp(os.Stdout, conn)
	mustCoyp(conn, os.Stdin)
}

//!-

func mustCoyp(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
