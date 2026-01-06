package main

import (
	"fmt"
	"time"
)

// Send data on a channel
// 🔍 SET BREAKPOINT HERE
func sender(ch chan int, value int) {
	fmt.Printf("Sender: about to send %d\n", value)
	// 🔍 SET BREAKPOINT HERE — Will block if channel is unbuffered
	ch <- value
	fmt.Printf("Sender: sent %d\n", value)
}

// Receive data from a channel
// 🔍 SET BREAKPOINT HERE
func receiver(ch chan int, id int) {
	fmt.Printf("Receiver %d: waiting for value\n", id)
	// 🔍 SET BREAKPOINT HERE — Will block until value arrives
	value := <-ch
	fmt.Printf("Receiver %d: received %d\n", id, value)
}

func main() {
	fmt.Println("=== Unbuffered Channel (Synchronous) ===")

	// 🔍 SET BREAKPOINT HERE
	unbuffered := make(chan int)

	// Launch receiver first
	// 🔍 SET BREAKPOINT HERE
	go receiver(unbuffered, 1)

	// Give receiver time to start
	time.Sleep(10 * time.Millisecond)

	// 👀 Check Goroutines panel — receiver is blocked on channel read

	// 🔍 SET BREAKPOINT HERE — Send will unblock receiver
	unbuffered <- 42

	time.Sleep(10 * time.Millisecond)
	fmt.Println()

	fmt.Println("=== Buffered Channel (Asynchronous) ===")

	// 🔍 SET BREAKPOINT HERE
	buffered := make(chan int, 2) // Buffer size = 2

	// Send without receiver (won't block while buffer has space)
	// 🔍 SET BREAKPOINT HERE
	buffered <- 1
	fmt.Println("Sent 1 (buffer: 1/2)")

	buffered <- 2
	fmt.Println("Sent 2 (buffer: 2/2)")

	// 👀 Next send would block (buffer full)

	// Receive values
	// 🔍 SET BREAKPOINT HERE
	v1 := <-buffered
	fmt.Printf("Received %d (buffer: 1/2)\n", v1)

	v2 := <-buffered
	fmt.Printf("Received %d (buffer: 0/2)\n\n", v2)

	fmt.Println("=== Deadlock Detection ===")

	// Uncomment to see deadlock:
	// deadlockChan := make(chan int)
	// 🔍 SET BREAKPOINT HERE
	// <-deadlockChan // ⚠️ This will deadlock (no sender)

	fmt.Println("(Deadlock example commented out)\n")

	fmt.Println("=== Select Statement ===")

	// 🔍 SET BREAKPOINT HERE
	ch1 := make(chan int)
	ch2 := make(chan string)

	// Send to ch1 after a delay
	go func() {
		time.Sleep(30 * time.Millisecond)
		ch1 <- 100
	}()

	// Send to ch2 after a longer delay
	go func() {
		time.Sleep(60 * time.Millisecond)
		ch2 <- "hello"
	}()

	// 🔍 SET BREAKPOINT HERE — Select waits for first available channel
	select {
	case v := <-ch1:
		fmt.Printf("Received from ch1: %d\n", v)
	case v := <-ch2:
		fmt.Printf("Received from ch2: %s\n", v)
	}

	fmt.Println()

	fmt.Println("=== Closing Channels ===")

	// 🔍 SET BREAKPOINT HERE
	closable := make(chan int, 3)

	// Send some values
	closable <- 1
	closable <- 2
	closable <- 3

	// 🔍 SET BREAKPOINT HERE — Close the channel
	close(closable)

	// 👀 Can still receive from closed channel until empty
	for v := range closable {
		fmt.Printf("Received: %d\n", v)
	}

	// 🔍 SET BREAKPOINT HERE — Receiving from closed, empty channel returns zero value
	v, ok := <-closable
	fmt.Printf("After close: v=%d, ok=%v\n", v, ok)
}
