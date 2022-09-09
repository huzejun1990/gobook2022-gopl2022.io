// Package memo provides a concurrency-safe memoization a function of
// 包 memo 提供了一个并发安全的 memoization 功能
// type Func.  Requests for different keys run concurrently.
// 类型函数。 对不同密钥的请求同时运行。
// Concurrent requests for the same key result in duplicate work.
// 对同一个键的并发请求会导致重复工作。
package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

//!+

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the two critical sections, several goroutines
		// 在两个临界区之间，有几个 goroutine
		// may race to compute f(key) and update the map.
		// 可能会竞相计算 f(key) 并更新地图。
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}

//!-
