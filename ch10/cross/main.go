// The cross command prints the values of GOOS and GOARCH for this target.
// cross 命令为这个目标打印 GOOS 和 GOARCH 的值。
package main

import (
	"fmt"
	"runtime"
)

//!+
func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

//!-
