// basename removes directory components and a .suffix.
// basename 删除目录组件和 .suffix。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

func basename(s string) string {
	// Discard last '/' and everything before.
	// 丢弃最后一个'/'和之前的所有内容。
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	//  Preserve everything before last '.'.
	// 保留最后一个'.'之前的所有内容。
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
