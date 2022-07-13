// Rev reverses a slice.

package main

import "fmt"

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) //"[5 4 3 2 1 0]"
	// !-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions
	// 将 s 向左旋转两个位置
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	// !-slice

	// Interactive test of reverse

}

// reverse reverses a slice of ints in place.
//reverse 将一片 int 反转到位。
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
