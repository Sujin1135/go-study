package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup
	count := 10

	account := &Account{10}
	wg.Add(count)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func DepositAndWithdraw(acc *Account) {
	fmt.Printf("account balance is %d\n", acc.Balance)
	if acc.Balance < 0 || acc.Balance > 1010 {
		panic(fmt.Sprintf("Balance should not be negative value or bigger than 1010: %d", acc.Balance))
	}
	mutex.Lock()
	defer mutex.Unlock()

	acc.Balance += 1000
	time.Sleep(time.Millisecond)
	acc.Balance -= 1000
}
