// Package bank provides a concurrency-safe bank with one account.
// bank包 银行提供一个具有一个帐户的并发安全银行。
package bank

var deposits = make(chan int) // send amount to deposit // 把钱发到帐号上
var balances = make(chan int) // receive balance	// 收到余额

func Deposit(amount int) {
	deposits <- amount
}
func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balance is  confined to teller goroutine // balance 仅限于柜员 goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine //启动监视器 goroutine
}

//!-
