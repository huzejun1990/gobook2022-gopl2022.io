// Countdown implements the countdown for a rocket launch.
// Countdown 实现火箭发射的倒计时。
package main

import (
	"fmt"
	"os"
	"time"
)

// NOTE: the ticker goroutine never terminates if the launch is aborted.
// 注意：如果启动被中止，ticker goroutine 永远不会终止。
// This is a "goroutine leak".
// 这是一个“goroutine 泄漏”。

//!+
func main() {
	// ...create abort channel...
	// ...创建中止通道

	//!-

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte // 读取单个字节
		abort <- struct{}{}
	}()

	//!+
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown < 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
