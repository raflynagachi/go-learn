package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup, num int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("number: ", num)
	time.Sleep(time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go RunAsync(&group, i)
	}
	group.Wait()
	fmt.Println("Complete")
}

func TestOnce(t *testing.T) {
	group := sync.WaitGroup{}
	once := sync.Once{}
	var counter, counter2 int

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(func() {
				counter++
			})
			counter2++
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println(counter, counter2)
}

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any { return false },
	}
	group := sync.WaitGroup{}
	// mutex := sync.Mutex{}

	pool.Put("Rafly")
	pool.Put(2112)
	pool.Put("Nagachi")

	for i := 0; i < 3; i++ {
		group.Add(1)
		go func() {
			// mutex.Lock()
			defer group.Done()

			data := pool.Get()
			fmt.Println("data:", data)
			pool.Put(data)
			// mutex.Unlock()
		}()
	}

	group.Wait()
	fmt.Println("Complete")
}

func AddMap(maph *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	maph.Store(value, value)
}

// sync.Map
func TestMap(t *testing.T) {
	group := sync.WaitGroup{}
	maph := sync.Map{}

	for i := 0; i < 100; i++ {
		// group add taroh di luar go function
		group.Add(1)
		go AddMap(&maph, i, &group)
	}

	group.Wait()
	maph.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

// sync.Cond
var group = sync.WaitGroup{}
var mutex = sync.Mutex{}
var cond = sync.NewCond(&mutex)

func WaitCondition(value int) {
	defer func() {
		cond.L.Unlock()
		group.Done()
	}()

	cond.L.Lock()

	// function waiting for Signal() or Broadcast()
	cond.Wait()
	fmt.Println("DONE")
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}
	go func() {
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second)
			cond.Signal()
		}
	}()
	go func() {
		time.Sleep(2 * time.Second)
		cond.Broadcast()
	}()
	group.Wait()
}
