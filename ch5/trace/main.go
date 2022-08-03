// The trace program uses defer to add entry/exit diagnostics to a function.
// 跟踪程序使用 defer 向函数添加进入/退出诊断。
package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the
	//extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow
	//operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}

/*
!+output
$ go build gopl2022.io/ch5/trace
$ ./trace
2022/08/04 02:37:50 enter bigSlowOperation
2022/08/04 02:38:00 exit bigSlowOperation (10.0915963s)
!-output
*/
