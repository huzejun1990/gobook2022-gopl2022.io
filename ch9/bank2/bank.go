// Package bank provides a concurrency-safe bank with one account.
// 包银行提供一个具有一个帐户的并发安全银行。
package bank

//!+
var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance // 二进制信号量保护平衡
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token // 获取token
	balance = balance + amount
	<-sema // release token  //重新生成token
}

func Balance() int {
	sema <- struct{}{} // acquire token // 获取token
	b := balance
	<-sema // release token // 重新生成token
	return b
}

//!-
