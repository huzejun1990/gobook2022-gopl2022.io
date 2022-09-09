// Package bank provides a concurrency-safe single-account bank.
// 包银行提供并发安全的单账户银行。
package bank

import "sync"

//!+

var (
	mu      sync.Mutex //guards balance	// 守卫平衡
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
