package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestGoroutines(t *testing.T) {
	t.Run("OneGoroutines", func(t *testing.T) {
		fmt.Println("Goroutines")
		go Hello("Nagachi")
		fmt.Println("Testing done")
		time.Sleep(time.Second)
	})
	t.Run("ManyGoroutines", func(t *testing.T) {
		for i := 0; i < 200; i++ {
			go Hello(strconv.FormatInt(int64(i), 10))
		}
		time.Sleep(5 * time.Second)
	})
}
