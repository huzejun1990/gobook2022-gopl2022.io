package main

import (
	"bufio"
	"fmt"
	"os"
)

// once in the standard input,preceded by its count.
//一次在标准输入中，前面是它的计数。
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Error()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
