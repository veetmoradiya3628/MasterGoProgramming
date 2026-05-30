package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex // mutex to protect balance
}

func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock() // lock the mutex before modifying the balance
	b.balance += amount
	b.mutex.Unlock() // unlock the mutex after modifying the balance
	fmt.Println("Balance after deposit : ", b.balance)
}

func (b *BankAccount) Withdraw(amount int) {
	if amount > b.balance {
		fmt.Println("Insufficient funds")
		return
	}
	b.mutex.Lock() // lock the mutex before modifying the balance
	b.balance -= amount
	b.mutex.Unlock() // unlock the mutex after modifying the balance
	fmt.Println("Balance after withdraw : ", b.balance)
}

func (b *BankAccount) Balance() int {
	b.mutex.Lock() // lock the mutex before reading the balance
	balance := b.balance
	b.mutex.Unlock() // unlock the mutex after reading the balance
	return balance
}

func main() {
	counter := 0 // critical section
	var wg sync.WaitGroup
	var mutex sync.Mutex // mutex to protect the counter

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock() // lock the mutex before modifying the counter
			counter++
			mutex.Unlock() // unlock the mutex after modifying the counter
		}()
	}

	wg.Wait()
	fmt.Println(counter)

	// BankAccount example
	var bankWg sync.WaitGroup
	var bankAccount = &BankAccount{balance: 100}
	for i := 0; i < 10; i++ {
		bankWg.Add(1)
		go func() {
			defer bankWg.Done()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			bankAccount.Deposit(50)
		}()

		bankWg.Add(1)
		go func() {
			defer bankWg.Done()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			bankAccount.Withdraw(30)
		}()
	}
	bankWg.Wait()
	fmt.Println("Final Balance: ", bankAccount.Balance())
}
