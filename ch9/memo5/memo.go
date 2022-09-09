// Package memo provides a concurrency-safe non-blocking memoization
// 包 memo 提供了一个并发安全的非阻塞 memoization
// of a function.  Requests for different keys proceed in parallel.
// 一个函数。 对不同密钥的请求并行进行。
// Concurrent requests for the same key block until the first completes.
// 对相同密钥块的并发请求，直到第一个完成。
// This implementation uses a monitor goroutine.
// 这个实现使用了一个监视器 goroutine。

package memo

//!+Func

// Func is the type of the function to memoize.
// Func 是要记忆的函数的类型。
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
// 结果是调用 Func 的结果。
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready // 当 res 准备好时关闭
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
// 请求是请求将 Func 应用于 key 的消息。
type request struct {
	key      string
	response chan<- result // the client wants a single result // 客户想要一个结果
}

type Memo struct {
	requests chan request
}

// New returns a memoization of f.  Clients must subsequently call Close.
// New 返回 f 的记忆。 客户随后必须调用 Close。
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

/*func (memo *Memo) Close() {
	close(memo.requests)
}*/

func (memo *Memo) Close() {
	close(memo.requests)
}

//!-get

//!+monitor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key. // 这是对该密钥的第一个请求。
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}

}

func (e *entry) call(f Func, key string) {
	// Evaluate the function. // 评估函数。
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition. //  广播就绪条件。
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	// 等待就绪状态。
	<-e.ready
	// Send the result to the client.
	// 将结果发送给客户端。
	response <- e.res
}

//!-monitor
