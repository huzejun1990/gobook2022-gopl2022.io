package memo

import (
	"testing"

	"gopl2022.io/ch9/memo3"
	"gopl2022.io/ch9/memotest"
)

var httpGetBody = memotest.HttpGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
