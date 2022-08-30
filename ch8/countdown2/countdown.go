// Countdown implements the countdown for a rocket launch.
// Countdown 实现火箭发射的倒计时。
package main

import (
	"fmt"
	"os"
	"time"
)

//!+
func main() {
	//...create abort channel...

	//!-

	//!+abort
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte // 读取单个字节
		abort <- struct{}{}
	}()
	//!-abort

	//!+
	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
