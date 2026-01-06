package main

import (
	"errors"
	"fmt"
)

// Function with defer
// 🔍 SET BREAKPOINT HERE
func deferredCleanup() {
	fmt.Println("Function started")

	// 🔍 SET BREAKPOINT HERE — Defer is REGISTERED, not executed
	defer fmt.Println("Defer 1: This runs last")
	defer fmt.Println("Defer 2: This runs second")
	defer fmt.Println("Defer 3: This runs first")

	// 👀 Deferred functions run in LIFO order (stack)

	fmt.Println("Function about to return")
	// 🔍 SET BREAKPOINT HERE — Defers haven't run yet
}

// Named return with defer
// 🔍 SET BREAKPOINT HERE
func namedReturn() (result string) {
	// 🔍 SET BREAKPOINT HERE
	defer func() {
		// 👀 This can modify the named return value
		result = result + " (modified by defer)"
	}()

	result = "original"
	// 🔍 SET BREAKPOINT HERE — Before return
	return result // Defer runs AFTER this assignment but BEFORE the actual return
}

// Error handling with defer
// 🔍 SET BREAKPOINT HERE
func processWithError(shouldFail bool) (err error) {
	// 🔍 SET BREAKPOINT HERE
	defer func() {
		if err != nil {
			// 👀 Defer can inspect and modify the error
			fmt.Println("Defer saw error:", err)
			err = fmt.Errorf("wrapped: %w", err)
		}
	}()

	if shouldFail {
		return errors.New("something failed")
	}

	return nil
}

// Panic and recover
// 🔍 SET BREAKPOINT HERE
func panicAndRecover() {
	// 🔍 SET BREAKPOINT HERE
	defer func() {
		if r := recover(); r != nil {
			// 👀 Recover catches the panic
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to panic")
	// 🔍 SET BREAKPOINT HERE
	panic("something went wrong!")

	// ⚠️ This line never executes
	fmt.Println("This never prints")
}

// Defer with loop variable capture bug
// 🔍 SET BREAKPOINT HERE
func deferInLoop() {
	for i := 0; i < 3; i++ {
		// ⚠️ All defers will see i = 3 (loop variable capture)
		defer fmt.Println("Deferred i (buggy):", i)
	}
	// 🔍 SET BREAKPOINT HERE — All defers are scheduled but not run
}

// Correct version
// 🔍 SET BREAKPOINT HERE
func deferInLoopFixed() {
	for i := 0; i < 3; i++ {
		i := i // Shadow to capture current value
		defer fmt.Println("Deferred i (fixed):", i)
	}
}

func main() {
	fmt.Println("=== Basic Defer ===")

	// 🔍 SET BREAKPOINT HERE — Step into deferredCleanup
	deferredCleanup()
	fmt.Println("Back in main\n")

	fmt.Println("=== Named Return with Defer ===")

	// 🔍 SET BREAKPOINT HERE
	result := namedReturn()
	fmt.Printf("Result: %s\n\n", result)

	fmt.Println("=== Error Handling with Defer ===")

	// 🔍 SET BREAKPOINT HERE — Step into processWithError
	err := processWithError(true)
	fmt.Printf("Error: %v\n\n", err)

	fmt.Println("=== Panic and Recover ===")

	// 🔍 SET BREAKPOINT HERE — Step into panicAndRecover
	panicAndRecover()
	fmt.Println("Survived the panic!\n")

	fmt.Println("=== Defer in Loop (Buggy) ===")

	// 🔍 SET BREAKPOINT HERE
	deferInLoop()
	fmt.Println()

	fmt.Println("=== Defer in Loop (Fixed) ===")

	// 🔍 SET BREAKPOINT HERE
	deferInLoopFixed()
}
