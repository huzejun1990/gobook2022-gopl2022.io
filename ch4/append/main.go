// Append illustrates the behavior of the built-in append function.
// Append 说明了内置 append 函数的行为

package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.  //有成长的空间。 扩展切片。
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.	// 空间不足。 分配一个新数组。
		// Grow by doubling, for amortized linear complexity.	// 通过加倍增长，用于摊销线性复杂度
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) //a built-in function; see text 	// 内置函数； 见文字
	}
	z[len(x)] = y
	return z

}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

}

/**
0 cap=1 [0]
1 cap=2 [0 1]
2 cap=4 [0 1 2]
3 cap=4 [0 1 2 3]
4 cap=8 [0 1 2 3 4]
5 cap=8 [0 1 2 3 4 5]
6 cap=8 [0 1 2 3 4 5 6]
7 cap=8 [0 1 2 3 4 5 6 7]
8 cap=16        [0 1 2 3 4 5 6 7 8]
9 cap=16        [0 1 2 3 4 5 6 7 8 9]

*/
