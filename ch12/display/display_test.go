package display

import (
	"io"
	"net"
	"os"
	"reflect"
	"sync"
	"testing"

	"gopl2022.io/ch7/eval"
)

// NOTE: we can't use !+..!- comments to excerpt these tests
// 注意：我们不能使用 !+..!- 注释来摘录这些测试
// into the book because it defeats the Example mechanism,
// 因为它破坏了示例机制，所以进入本书，
// which requires the // Output comment to be at the end
// 这要求 // 输出注释在末尾
// of the function.
//函数的。

func Example_expr() {
	e, _ := eval.Parse("sqrt(A / pi)")
	Display("e", e)
}

func Example_slice() {
	Display("slice", []*int{new(int), nil})
}

func Example_nilInterface() {
	var w io.Writer
	Display("w", w)
}

func Example_ptrToInterface() {
	var w io.Writer
	Display("&w", &w)
}

func Example_struct() {
	Display("x", struct {
		x interface{}
	}{3})
}

func Example_interface() {
	var i interface{} = 3
	Display("i", i)
}

func Example_ptrToInterface2() {
	var i interface{} = 3
	Display("&i", &i)
}

func Example_array() {
	Display("x", [1]interface{}{3})
}

func Example_movie() {
	//!+movie
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	//!-movie
	//!-stringelove
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	//!-strangelove
	Display("strangelove", strangelove)

}

// This test ensures that the program terminates without crashing.
// 这个测试确保程序在没有崩溃的情况下终止。
func Test(t *testing.T) {
	// Some other values (YMMV)	// 一些其他值 (YMMV)
	Display("os.Stderr", os.Stderr)

	var w io.Writer = os.Stderr
	Display("&w", &w)

	var locker sync.Locker = new(sync.Mutex)
	Display("(&locker)", &locker)

	Display("locker", locker)

	locker = nil
	Display("(&locker)", &locker)

	ips, _ := net.LookupHost("golang.org")
	Display("ips", ips)

	Display("rV", reflect.ValueOf(os.Stderr))

	type P *P
	var p P
	p = &p
	if false {
		Display("p", p)
	}

	// a map that contains itself
	// 包含自身的地图
	type M map[string]M
	m := make(M)
	m[""] = m
	if false {
		Display("m", m)
	}

	// a slice that contains itself
	// 一个包含自身的切片
	type S []S
	s := make(S, 1)
	s[0] = s
	if false {
		Display("s", s)
	}

	// a linked list that eats its own tail
	// 一个吃掉自己尾巴的链表
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	if false {
		Display("c", c)
	}
}
