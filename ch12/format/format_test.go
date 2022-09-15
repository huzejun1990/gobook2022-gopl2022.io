package format_test

import (
	"fmt"
	"gopl2022.io/ch12/format"
	"testing"
	"time"
)

func Test(t *testing.T) {
	// The pointer values are just examples, and may vary from run to run.
	// 指针值只是示例，可能因运行而异。
	//!+time
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x))                  // 1
	fmt.Println(format.Any(d))                  // 1
	fmt.Println(format.Any([]int64{x}))         // []int64 0x c00009e120
	fmt.Println(format.Any([]time.Duration{d})) // []time.Duration 0x c00009e128
}
