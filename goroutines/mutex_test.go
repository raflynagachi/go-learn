package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var x int
	t.Run("TestWithoutMutex", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					x = x + 1
				}
			}()
		}
	})
	time.Sleep(2 * time.Second)
	fmt.Println("Counter:", x)

	x = 0
	t.Run("TestWithMutex", func(t *testing.T) {
		var mutex sync.Mutex
		for i := 0; i < 100; i++ {
			go func() {
				for j := 0; j < 100; j++ {
					mutex.Lock()
					x = x + 1
					mutex.Unlock()
				}
			}()
		}
	})
	time.Sleep(2 * time.Second)
	fmt.Println("Counter:", x)
}

func TestMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			account.AddBalance(1)
			fmt.Println(account.GetBalance())
		}
	}
	time.Sleep(4 * time.Second)
	fmt.Println("Balance: ", account.GetBalance())
}

// Simulasi deadlock
type Account struct {
	sync.Mutex
	name    string
	balance int
}

func (account *Account) Lock() {
	account.Mutex.Lock()
}

func (account *Account) Unlock() {
	account.Mutex.Unlock()
}

func (account *Account) Change(amount int) {
	account.balance += amount
}

func Transfer(user1, user2 *Account, amount int) {
	user1.Lock()
	fmt.Println("user1 locked")
	user1.Change(-amount)

	time.Sleep(time.Second)

	user2.Lock()
	fmt.Println("user1 locked")
	user2.Change(amount)

	time.Sleep(time.Second)
	user1.Unlock()
	user2.Unlock()
	fmt.Println("user1 & user2 unlocked")
}

func TestDeadlock(t *testing.T) {
	user1 := Account{
		name:    "Nagachi",
		balance: 100,
	}
	user2 := Account{
		name:    "Rafly",
		balance: 100,
	}
	go Transfer(&user1, &user2, 20)
	go Transfer(&user2, &user1, 15)

	time.Sleep(3 * time.Second)

	fmt.Println(user1)
	fmt.Println(user2)
}

// Atomic
func TestAtomic(t *testing.T) {
	var counter int64
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&counter, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println(counter)
}
