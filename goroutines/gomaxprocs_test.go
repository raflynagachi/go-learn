package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	totalThread := runtime.GOMAXPROCS(-1)
	totalGo := runtime.NumGoroutine()

	fmt.Println("CPU: ", totalCpu)
	fmt.Println("Thead: ", totalThread)
	fmt.Println("Goroutine: ", totalGo)

	group.Wait()
}
