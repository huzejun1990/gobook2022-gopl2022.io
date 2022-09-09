// Package memo provides a concurrency-safe memoization a function of
// 包 memo 提供了一个并发安全的 memoization 功能
// a function.  Requests for different keys proceed in parallel.
// 一个函数。 对不同密钥的请求并行进行。
// Concurrent requests for the same key block until the first completes.
// 对相同密钥块的并发请求，直到第一个完成。
// This implementation uses a Mutex.
// 此实现使用互斥体。

package memo

import "sync"

// Func is the type of the function to memoize.
// Func 是要记忆的函数的类型。
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//!+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready	// 当 res 准备好时关闭
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache	// 守卫缓存
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// 这是对该密钥的第一个请求。
		// This goroutine becomes responsible for computing
		// 这个goroutine负责计算
		// the value and broadcasting the ready condition.
		// 值和广播就绪条件。
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition // 广播就绪状态
	} else {
		// This is a repeat request for this key.
		// 这是对该键的重复请求。
		memo.mu.Unlock()

		<-e.ready // wait for ready condition	// 等待就绪状态

	}
	return e.res.value, e.res.err

}

//!-
