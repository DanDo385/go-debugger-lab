package main

import "fmt"

// 🔍 SET BREAKPOINT HERE — Third level of the call stack
func deepFunction(value int) int {
	result := value * 2
	fmt.Printf("deepFunction: value=%d, result=%d\n", value, result)
	return result // 👀 Watch the return value in the debugger
}

// 🔍 SET BREAKPOINT HERE — Second level of the call stack
func middleFunction(value int) int {
	fmt.Println("middleFunction: calling deepFunction")
	result := deepFunction(value + 10) // 🔍 Step Into (F11) here to see deepFunction
	fmt.Printf("middleFunction: received %d from deepFunction\n", result)
	return result + 5
}

// 🔍 SET BREAKPOINT HERE — First level of the call stack
func topFunction(value int) int {
	fmt.Println("topFunction: calling middleFunction")
	result := middleFunction(value) // 🔍 Step Into (F11) here to descend the stack
	fmt.Printf("topFunction: received %d from middleFunction\n", result)
	return result
}

// Recursive function to demonstrate stack growth
// 🔍 SET BREAKPOINT HERE — Watch the call stack grow
func factorial(n int) int {
	fmt.Printf("factorial(%d) called\n", n)

	if n <= 1 {
		// 👀 BASE CASE — Check the call stack depth here
		fmt.Println("Base case reached, returning 1")
		return 1
	}

	// 🔍 Step Into (F11) to see the stack grow
	result := n * factorial(n-1)

	// 👀 Watch the stack unwind as each recursive call returns
	fmt.Printf("factorial(%d) returning %d\n", n, result)
	return result
}

func main() {
	fmt.Println("=== Nested Function Calls ===")

	// 🔍 SET BREAKPOINT HERE
	result := topFunction(5)
	fmt.Printf("Final result: %d\n\n", result)

	fmt.Println("=== Recursive Function ===")

	// 🔍 SET BREAKPOINT HERE — Then step into factorial
	factorial(5)

	fmt.Println("\n=== Stack Frames Example ===")
	demonstrateStackFrames()
}

// Demonstrates that each function call has its own stack frame
func demonstrateStackFrames() {
	// 🔍 SET BREAKPOINT HERE
	x := 100
	fmt.Println("demonstrateStackFrames: x =", x)

	// 👀 Watch how 'x' in this frame is different from 'x' in the helper
	helperWithSameName()

	// 🔍 SET BREAKPOINT HERE — x is still 100
	fmt.Println("demonstrateStackFrames: x is still", x)
}

func helperWithSameName() {
	// 🔍 SET BREAKPOINT HERE
	x := 200 // Same name, different stack frame
	fmt.Println("helperWithSameName: x =", x)

	// 👀 Look at the Call Stack panel
	// Click on different frames to see different values of 'x'
}
