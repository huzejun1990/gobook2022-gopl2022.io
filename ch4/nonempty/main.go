// Nonempty is an example of an in-place slice algorithm.
// 非空是就地切片算法的一个例子。
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// nonempty 返回一个只包含非空字符串的切片。
// The underlying array is modified during the call.
// 底层数组在调用过程中被修改。
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original // 原始的零长度的切片
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}

	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // ["one" "three"]
	fmt.Printf("%q\n", data)           //["one" "three" "three"]
	fmt.Println("------------------")
	//fmt.Printf("%q\n", nonempty2(data))	// ["one" "three"]
	//fmt.Printf("%q\n", data)	// ["one" "three" "three"]

}
