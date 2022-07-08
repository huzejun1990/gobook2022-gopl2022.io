// Printints demonstrates the use of bytes.Buffer to format a string.
// Printints 演示了使用 bytes.Buffer 来格式化字符串。
package main

import (
	"bytes"
	"fmt"
)

//Printints 演示了使用 bytes.Buffer 来格式化字符串。intsToString 类似于 fmt.Sprint(values) 但添加了逗号。
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1,2,3]"
}
