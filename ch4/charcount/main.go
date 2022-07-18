// Charcount computes counts of Unicode characters.
// Charcount 计算 Unicode 字符的计数。
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    //counts of unicode characters // unicode 字符数
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings // UTF-8 编码的长度计数
	invaliid := 0                   // count of invalid UTF-8 characters // 无效 UTF-8 字符数

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invaliid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invaliid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invaliid)
	}
}
