// The squares program demonstrates a function value with state.
// squares 程序演示了一个带状态的函数值。
package main

import "fmt"

// squares 返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) //"1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) //  "9"
	fmt.Println(f()) //	"16"
}
