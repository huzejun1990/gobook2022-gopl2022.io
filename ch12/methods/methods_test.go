package methods_test

import (
	"gopl2022.io/ch12/methods"
	"strings"
	"time"
)

func ExamplePrintDuration() {
	methods.Print(time.Hour)
}

func ExamplePrintReplacer() {
	methods.Print(new(strings.Replacer))
}
