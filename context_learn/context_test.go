package context_learn

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()
	contextB := context.WithValue(contextA, 1, "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println("====GET VALUE====")
	fmt.Println(contextA.Value("a"))
	fmt.Println(contextB.Value(1))
	fmt.Println(contextF.Value("c"))
}

func CreateCounter(ctx context.Context) chan int {
	dest := make(chan int)
	go func() {
		defer close(dest)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dest <- counter
				counter++
				time.Sleep(time.Second)
			}
		}
	}()
	return dest
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	dest := CreateCounter(ctx)
	for v := range dest {
		fmt.Println("counter:", v)
		if v == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context
	time.Sleep(time.Second)

	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	dest := CreateCounter(ctx)
	for v := range dest {
		fmt.Println("counter:", v)
		if v == 10 {
			break
		}
	}

	time.Sleep(time.Second)

	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	dest := CreateCounter(ctx)
	for v := range dest {
		fmt.Println("counter:", v)
		if v == 10 {
			break
		}
	}

	time.Sleep(time.Second)

	fmt.Println(runtime.NumGoroutine())
}
