// Chat is a server that lets clients chat with each other.
// Chat 是一个服务器，让客户端可以互相聊天。
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages // 所有传入的客户端消息
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients // 所有连接的客户端
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// 向所有人广播传入的消息
			// clients' outgoing message channels.
			// 客户端的传出消息通道。
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!-handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages // 传出客户端消息
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()
	//注意：忽略来自 input.Err() 的潜在错误

	leaving <- ch
	messages <- who + "has left "
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors // 忽略网络错误
	}
}

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
