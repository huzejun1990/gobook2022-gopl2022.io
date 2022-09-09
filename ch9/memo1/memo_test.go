package memo_test

import (
	"testing"

	"gopl2022.io/ch9/memo1"
	"gopl2022.io/ch9/memotest"
)

var httpGetBody = memotest.HttpGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
// 注意：不是并发安全的！ 测试失败。
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

/*
//!+output

$ go test -run=TestConcurrent -race -v gopl2022.io/ch9/memo1
=== RUN   TestConcurrent
https://godoc.org,1.0143542s, 17461 bytes
https://godoc.org,1.0608759s, 17461 bytes
http://gopl.io,1.5393446s, 4154 bytes
http://gopl.io,1.5508439s, 4154 bytes
2022/09/09 18:48:52 Get "https://play.golang.org": dial tcp 142.251.43.17:443: connectex: A connection attempt failed be
cause the connected party did not properly respond after a period of time, or established connection failed because conn
ected host has failed to respond.
2022/09/09 18:48:52 Get "https://play.golang.org": dial tcp 142.251.43.17:443: connectex: A connection attempt failed be
cause the connected party did not properly respond after a period of time, or established connection failed because conn
ected host has failed to respond.
2022/09/09 18:48:52 Get "https://golang.org": dial tcp 142.251.43.17:443: connectex: A connection attempt failed because
 the connected party did not properly respond after a period of time, or established connection failed because connected
 host has failed to respond.
2022/09/09 18:48:52 Get "https://golang.org": dial tcp 142.251.43.17:443: connectex: A connection attempt failed because
 the connected party did not properly respond after a period of time, or established connection failed because connected
 host has failed to respond.
--- PASS: TestConcurrent (21.19s)
PASS
ok      gopl2022.io/ch9/memo1   23.128s

//!-output


*/
