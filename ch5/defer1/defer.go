// Defer1 demonstrates a deferred call being invoked during a panic.
// Defer1 演示了在恐慌期间调用的延迟调用。
package main

import (
	"fmt"
)

//!+
func main() {
	f(3)
}

//!+f
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

//!-f

/*
//!+stdout
$ go run defer.go
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.f(0x247ba0)
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:16 +0x157
main.f(0x1)
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:18 +0x132
main.f(0x2)
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:18 +0x132
main.f(0x3)
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:18 +0x132
main.main()
        F:/gobook2022/gopl2022.io/ch5/defer1/defer.go:11 +0x1e
exit status 2


!-stdout
*/
