package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	t.Run("channelDeclarations", func(t *testing.T) {
		channel := make(chan string)
		defer close(channel)

		// value to channel
		// channel <- "Nagachi"

		// channel to value
		// data := <-channel

		go func() {
			time.Sleep(2 * time.Second)
			channel <- "Rafly Rigan Nagachi"
			fmt.Println("Selesai mengirim data")
		}()

		time.Sleep(5 * time.Second)
		data := <-channel
		fmt.Println(data)
	})

	t.Run("channelParameter", func(t *testing.T) {
		channel := make(chan string)
		defer close(channel)

		go GiveMeResponse(channel)

		data := <-channel
		fmt.Println(data)
	})

	t.Run("channelInOut", func(t *testing.T) {
		channel := make(chan string)
		defer close(channel)

		go OnlyIn(channel)
		go OnlyOut(channel)

		time.Sleep(3 * time.Second)
	})
}

func TestBuffered(t *testing.T) {
	// buffered channel
	channel := make(chan string, 1)
	defer close(channel)

	go func() {
		channel <- "Rafly"
		channel <- "Rigan"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("TestBuffered DONE")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Data ke-" + strconv.Itoa(i)
		}
		// wajib di-close untuk kondisi berhenti
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v, "diterima")
	}
	fmt.Println("TestRangeChannel DONE")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	var counter int
	for counter != 2 {
		select {
		case data := <-channel1:
			fmt.Println(data)
			counter++
		case data := <-channel2:
			fmt.Println(data)
			counter++
		default:
			fmt.Println("waiting for data...")
		}
	}
}
