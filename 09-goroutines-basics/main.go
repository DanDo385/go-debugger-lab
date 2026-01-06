package main

import (
	"fmt"
	"time"
)

// Simple goroutine function
// 🔍 SET BREAKPOINT HERE
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// 👀 Watch this in the Goroutines panel
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Worker %d finished\n", id)
}

// Goroutine with shared variable access
// 🔍 SET BREAKPOINT HERE
func increment(counter *int, id int) {
	fmt.Printf("Goroutine %d: reading counter = %d\n", id, *counter)
	*counter++
	fmt.Printf("Goroutine %d: incremented counter to %d\n", id, *counter)
}

func main() {
	fmt.Println("=== Starting Goroutines ===")

	// 🔍 SET BREAKPOINT HERE
	fmt.Println("Main goroutine started")

	// Launch goroutines
	// 🔍 SET BREAKPOINT HERE — After launching goroutines
	go worker(1)
	go worker(2)
	go worker(3)

	// 👀 Open the Goroutines panel (in Call Stack area)
	// You'll see: main goroutine + 3 worker goroutines

	// ⚠️ If you Step Over (F10), you jump between goroutines unpredictably
	// 🔍 SET BREAKPOINT HERE
	fmt.Println("Main: goroutines launched")

	// Wait for goroutines to finish
	time.Sleep(200 * time.Millisecond)

	fmt.Println("Main: goroutines should be done\n")

	fmt.Println("=== Goroutine Interleaving ===")

	// 🔍 SET BREAKPOINT HERE
	counter := 0

	// Launch multiple goroutines that access the same variable
	// ⚠️ This has a race condition (we'll explore this in module 11)
	go increment(&counter, 1)
	go increment(&counter, 2)
	go increment(&counter, 3)

	// 🔍 SET BREAKPOINT HERE
	time.Sleep(100 * time.Millisecond)

	// 👀 What's the final value of counter?
	fmt.Printf("Final counter: %d\n\n", counter)

	fmt.Println("=== Goroutine with Closure ===")

	// 🔍 SET BREAKPOINT HERE
	for i := 0; i < 3; i++ {
		// ⚠️ Closure capture bug (same as module 02)
		go func() {
			fmt.Println("Goroutine (buggy) i =", i)
		}()
	}

	time.Sleep(50 * time.Millisecond)

	// 🔍 SET BREAKPOINT HERE
	for i := 0; i < 3; i++ {
		i := i // Capture value
		go func() {
			fmt.Println("Goroutine (fixed) i =", i)
		}()
	}

	time.Sleep(50 * time.Millisecond)

	fmt.Println("\n=== Anonymous Goroutine ===")

	// 🔍 SET BREAKPOINT HERE
	done := make(chan bool)

	go func() {
		// 🔍 SET BREAKPOINT HERE — Inside anonymous goroutine
		fmt.Println("Anonymous goroutine running")
		time.Sleep(50 * time.Millisecond)
		done <- true
	}()

	// 🔍 SET BREAKPOINT HERE — Main waiting
	fmt.Println("Main: waiting for goroutine")
	<-done
	fmt.Println("Main: goroutine finished")
}
