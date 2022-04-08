package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(3 * time.Second) // cara 1: time with channel
	channel := time.After(3 * time.Second)  // cara 2: only channel return
	fmt.Println(time.Now())

	timex := <-timer.C
	fmt.Println(timex)
	timex = <-channel
	fmt.Println(timex)

	// cara 3: AfterFunc
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(3*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	group.Wait()
	fmt.Println(time.Now())
}

func TestTicker(t *testing.T) {
	// ticker := time.NewTicker(2 * time.Second)
	tickerChannel := time.Tick(2 * time.Second)

	// for tick := range ticker.C {
	// 	fmt.Println(tick)
	// 	ticker.Stop()
	// }

	for tick := range tickerChannel {
		fmt.Println(tick)
	}
}
