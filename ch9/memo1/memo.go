// Package memo provides a concurrency-unsafe
// 包备忘录提供了一个不安全的并发
// memoization of a function of type Func.
// 一个 Func 类型的函数的记忆。
package memo

// A Memo caches the results of calling a Func.
// 备忘录缓存调用 Func 的结果。
type Memo struct {
	f     Func
	cache map[string]result
}

// Func is the type of the function to memoize.
// Func 是要记忆的函数的类型。
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// NOTE: not concurrency-safe!
// 注意: 不是并发安全的
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

//!-
