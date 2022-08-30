// Clock1 is a TCP server that periodically writes the time.
// Clock1 是定期写入时间的 TCP 服务器。
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //e.g., connection aborted	// 连接中止
			continue
		}
		handleConn(conn) // handle one connection at a time	// 一次处理一个连接
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected	//客户端断开连接
		}
		time.Sleep(1 * time.Second)
	}
}
