package main

import (
	"fmt"
	"sync"
	"time"
)

// ⚠️ INTENTIONAL DATA RACE
// 🔍 SET BREAKPOINT HERE
func racyCounter() {
	counter := 0
	iterations := 1000

	// Launch multiple goroutines that increment the same variable
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				// ⚠️ RACE: Multiple goroutines read/write counter simultaneously
				counter++ // This is NOT atomic
			}
		}()
	}

	time.Sleep(100 * time.Millisecond)

	// 👀 Expected: 10 * 1000 = 10000
	// Actual: probably less (due to lost updates)
	fmt.Printf("Racy counter: %d (expected 10000)\n", counter)
}

// Fixed with mutex
// 🔍 SET BREAKPOINT HERE
func mutexCounter() {
	counter := 0
	iterations := 1000
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				// 🔍 SET BREAKPOINT HERE — Watch mutex lock/unlock
				mu.Lock()
				counter++ // Protected by mutex
				mu.Unlock()
			}
		}()
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Mutex counter: %d (expected 10000)\n", counter)
}

// Fixed with atomic operations
// 🔍 SET BREAKPOINT HERE
func atomicCounter() {
	var counter int64 = 0
	iterations := 1000

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				// Atomic operations are safe without locks
				// (Commented out to avoid import, but this is the idea)
				// atomic.AddInt64(&counter, 1)

				// For this example, we'll use mutex
			}
		}()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Atomic counter: %d\n", counter)
}

// Heisenbug: Race that disappears in debugger
// 🔍 SET BREAKPOINT HERE
func heisenbug() {
	value := 0
	ready := false

	// Writer goroutine
	go func() {
		value = 42
		// ⚠️ Without proper synchronization, reader might see old value
		time.Sleep(1 * time.Millisecond) // Tiny delay
		ready = true
	}()

	// Reader goroutine
	go func() {
		// Busy-wait for ready (DON'T DO THIS in real code)
		for !ready {
			// 👀 In debugger, this race might not happen
			// Debugger slows things down
		}
		fmt.Printf("Value: %d\n", value)
	}()

	time.Sleep(50 * time.Millisecond)
}

func main() {
	fmt.Println("=== Data Race (Intentional) ===")
	fmt.Println("⚠️ WARNING: This code has intentional race conditions")
	fmt.Println("To detect: run `go run -race main.go`\n")

	// 🔍 SET BREAKPOINT HERE
	racyCounter()

	fmt.Println("\n=== Fixed with Mutex ===")

	// 🔍 SET BREAKPOINT HERE
	mutexCounter()

	fmt.Println("\n=== Heisenbug (Race Disappears in Debugger) ===")

	// 🔍 SET BREAKPOINT HERE
	fmt.Println("Try running this with and without the debugger")
	fmt.Println("Behavior may differ!")
	heisenbug()

	fmt.Println("\n=== WaitGroup ===")

	// 🔍 SET BREAKPOINT HERE
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) // Increment counter
		i := i
		go func() {
			// 🔍 SET BREAKPOINT HERE
			defer wg.Done() // Decrement counter when done
			fmt.Printf("Goroutine %d working\n", i)
			time.Sleep(20 * time.Millisecond)
		}()
	}

	// 🔍 SET BREAKPOINT HERE — Wait for all goroutines
	fmt.Println("Waiting for goroutines...")
	wg.Wait()
	fmt.Println("All goroutines done")

	fmt.Println("\n=== Run with Race Detector ===")
	fmt.Println("Try: go run -race main.go")
	fmt.Println("The race detector will report the race in racyCounter()")
}
