// Defer2 demonstrates a deferred call to runtime.Stack during a panic.
// Defer2 演示了恐慌期间对 runtime.Stack 的延迟调用。
package main

import (
	"fmt"
	"os"
	"runtime"
)

//!+
func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

//!-
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/*
//!+pringstack
$ go run defer.go
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
goroutine 1 [running]:
main.printStack()
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:20 +0x39
panic({0x189dc0, 0x22ef50})
        F:/environment/golangdownload/gopath/go1.17.11/src/runtime/panic.go:1038 +0x215
main.f(0x1b8c40)
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:26 +0x157
main.f(0x1)
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:28 +0x132
main.f(0x2)
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:28 +0x132
main.f(0x3)
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:28 +0x132
main.main()
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:15 +0x45
panic: runtime error: integer divide by zer


*/
