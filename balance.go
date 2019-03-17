// +build ignore

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	wg.Add(4)
}

func (acc *bankAccount) doWithdraw(amount int) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
	runtime.Gosched()
	acc.Withdraw(amount)
}

func (acc *bankAccount) String() string {
	if acc.Balance < 0 {
		return fmt.Sprintf("-£%d", -acc.Balance)
	}
	return fmt.Sprintf("£%d", acc.Balance)
}

func (acc *bankAccount) PrintBalance() {
	wg.Wait()
	fmt.Println()
	fmt.Println("Balance:", acc.String())
}

type bankAccount struct{ Balance int }

func (acc *bankAccount) Withdraw(amount int) {
	if bal := acc.Balance; bal >= amount { // if has enough money
		fmt.Printf("✓ Balance = %d\tWithdrawing %d\n", bal, amount)
		acc.Balance -= amount
		return
	}
	fmt.Printf("✗ Balance = %d\tCan't withdraw %d\n", acc.Balance, amount)
}

func main() {
	acc := &bankAccount{Balance: 10}
	go acc.doWithdraw(11) // impossible
	go acc.doWithdraw(5)
	go acc.doWithdraw(4)
	go acc.doWithdraw(7)

	// .. wait for goroutines to complete
	acc.PrintBalance() // Print final balance
}

// end OMIT
