package main

import "fmt"

// Simple function to demonstrate inlining
// 🔍 SET BREAKPOINT HERE — May not trigger with optimizations
func add(a, b int) int {
	return a + b
}

// Function with intermediate variables
// 🔍 SET BREAKPOINT HERE
func calculate(x int) int {
	// These intermediate variables may be optimized away
	temp1 := x * 2
	temp2 := temp1 + 10
	temp3 := temp2 * 3

	// 👀 With optimizations, temp1/temp2 may show "<optimized out>"
	result := temp3 - 5

	return result
}

// Function demonstrating dead code elimination
// 🔍 SET BREAKPOINT HERE
func deadCode() {
	x := 10

	// This code is "dead" — result is never used
	y := x * 2
	z := y + 100

	// Optimizer may eliminate y and z entirely

	fmt.Println("x =", x)
	// 👀 Try to inspect y and z in the debugger
	// With optimizations, they may not exist
	_ = z // Suppress "unused" warning
}

// Loop with optimization opportunities
// 🔍 SET BREAKPOINT HERE
func optimizedLoop() {
	sum := 0
	// The optimizer may unroll or transform this loop
	for i := 0; i < 10; i++ {
		sum += i
	}

	// 👀 Try stepping through the loop
	// With optimizations, you may skip iterations
	fmt.Println("Sum:", sum)
}

func main() {
	fmt.Println("=== Compiler Optimizations ===")
	fmt.Println("Run this in two modes:")
	fmt.Println("1. Optimizations OFF (default in our VS Code config):")
	fmt.Println("   go build -gcflags=all=-N -l")
	fmt.Println("2. Optimizations ON:")
	fmt.Println("   go build")
	fmt.Println()

	// 🔍 SET BREAKPOINT HERE
	result := add(5, 10)
	fmt.Println("add(5, 10) =", result)

	// 👀 With optimizations, `add` may be inlined
	// The breakpoint inside `add` may never trigger

	// 🔍 SET BREAKPOINT HERE
	calc := calculate(7)
	fmt.Println("calculate(7) =", calc)

	// 👀 Step into calculate and watch intermediate variables
	// With optimizations: "<optimized out>"
	// Without optimizations: visible values

	// 🔍 SET BREAKPOINT HERE
	deadCode()

	// 🔍 SET BREAKPOINT HERE
	optimizedLoop()

	fmt.Println()
	fmt.Println("=== Try This Experiment ===")
	fmt.Println("1. Debug with current config (optimizations OFF)")
	fmt.Println("   - All breakpoints should work")
	fmt.Println("   - All variables should be visible")
	fmt.Println()
	fmt.Println("2. Build an optimized binary:")
	fmt.Println("   go build -o optimized")
	fmt.Println("   dlv exec ./optimized")
	fmt.Println()
	fmt.Println("   - Breakpoints may not trigger")
	fmt.Println("   - Variables show '<optimized out>'")
	fmt.Println("   - Stepping may skip lines")
	fmt.Println()
	fmt.Println("This is why production debugging is hard.")
}
