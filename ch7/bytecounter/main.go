// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
// Bytecounter 演示了计算字节数的 io.Writer 的实现。
package main

import "fmt"

//!-ByteCounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter	//将 int 转换为 ByteCounter
	return len(p), nil
}

//!+ByteCounter

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5"
	c = 0          //	reset the counter // 重置计数器
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len ("hello,Dolly")
}
