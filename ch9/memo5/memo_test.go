package memo_test

import (
	"testing"

	"gopl2022.io/ch9/memo5"
	"gopl2022.io/ch9/memotest"
)

var httpGetBody = memotest.HttpGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}

/*
//!+output
$ go test -v gopl2022.io/ch9/memo1
=== RUN   Test
2022/09/08 01:24:07 Get "https://golang.org": dial tcp 142.251.43.17:443: connectex: A connection attempt failed because
 the connected party did not properly respond after a period of time, or established connection failed because connected
 host has failed to respond.
https://godoc.org,936.5692ms, 17461 bytes
2022/09/08 01:24:29 Get "https://play.golang.org": dial tcp 108.177.125.141:443: connectex: A connection attempt failed
because the connected party did not properly respond after a period of time, or established connection failed because co
nnected host has failed to respond.
http://gopl.io,1.442829s, 4154 bytes
2022/09/08 01:24:30 Get "https://golang.org": dial tcp 142.251.43.17:443: connectex: A connection attempt failed because
 the connected party did not properly respond after a period of time, or established connection failed because connected
 host has failed to respond.
https://godoc.org,0s, 17461 bytes
2022/09/08 01:24:30 Get "https://play.golang.org": dial tcp 108.177.125.141:443: connectex: A connection attempt failed
because the connected party did not properly respond after a period of time, or established connection failed because co
nnected host has failed to respond.

nnected host has failed to respond.
http://gopl.io,0s, 4154 bytes
--- PASS: Test (44.59s)
=== RUN   TestConcurrent
https://godoc.org,458.2549ms, 17461 bytes
https://godoc.org,471.732ms, 17461 bytes
http://gopl.io,485.1701ms, 4154 bytes
http://gopl.io,675.6458ms, 4154 bytes
2022/09/08 01:24:51 Get "https://golang.org": dial tcp 142.251.43.17:443: connectex: A connection attempt failed because
 the connected party did not properly respond after a period of time, or established connection failed because connected
 host has failed to respond.
2022/09/08 01:24:51 Get "https://play.golang.org": dial tcp 142.250.206.209:443: connectex: A connection attempt failed
because the connected party did not properly respond after a period of time, or established connection failed because co
nnected host has failed to respond.
2022/09/08 01:24:51 Get "https://play.golang.org": dial tcp 142.250.206.209:443: connectex: A connection attempt failed
because the connected party did not properly respond after a period of time, or established connection failed because co
nnected host has failed to respond.
2022/09/08 01:24:51 Get "https://golang.org": dial tcp 142.251.43.17:443: connectex: A connection attempt failed because
 the connected party did not properly respond after a period of time, or established connection failed because connected
 host has failed to respond.
--- PASS: TestConcurrent (21.06s)
PASS
ok      gopl2022.io/ch9/memo1   66.448s

//!-output
*/
