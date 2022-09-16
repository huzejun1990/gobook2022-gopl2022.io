// Package methods provides a function to print the methods of any value.
// 封装方法提供了打印任何值的方法的功能。
package methods

import (
	"fmt"
	"reflect"
	"strings"
)

//!+print
// Print prints the method set of the value x.
// Print 打印值 x 的方法集。
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}

//!-print
