package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Hello(name string) {
	fmt.Println("Hello, ", name)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Rafly Rigan Nagachi"
}

func OnlyIn(channel chan<- string) {
	fmt.Println("Data in: ")
	// data := <-channel 	// error send-only channel
	channel <- "Rafly Nagachi"
	time.Sleep(1 * time.Second)
}

func OnlyOut(channel <-chan string) {
	fmt.Print("Data out: ")
	data := <-channel
	fmt.Println(data)
	time.Sleep(1 * time.Second)
}

// ======================Mutex======================
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (ba *BankAccount) AddBalance(amount int) {
	ba.RWMutex.Lock()
	ba.Balance += amount
	ba.RWMutex.Unlock()
}

func (ba *BankAccount) GetBalance() int {
	ba.RWMutex.RLock()
	balance := ba.Balance
	ba.RWMutex.RUnlock()
	return balance
}
