// Package memo provides a concurrency-safe memoization a function of
// 包 memo 提供了一个并发安全的 memoization 功能
// type Func.  Concurrent requests are serialized by a Mutex.
// 类型函数。 并发请求由 Mutex 序列化。

package memo

import "sync"

// Func is the type of the function to memoize.
// Func 是要记忆的函数的类型。
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{}
}

//!+

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

// Get is concurrency-safe.
// Get 是并发安全的。
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
