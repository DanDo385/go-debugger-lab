package main

import "fmt"

func main() {
	// 🔍 SET BREAKPOINT HERE
	x := 10
	fmt.Println("Outer x:", x) // 👀 x = 10

	// Variable shadowing — same name, different variable
	{
		// 🔍 SET BREAKPOINT HERE
		x := 20 // ⚠️ This is a DIFFERENT variable
		fmt.Println("Inner x:", x) // 👀 x = 20

		// 👀 WATCH THE ADDRESS of x here vs outer x
		fmt.Printf("Inner x address: %p\n", &x)
	}

	// 🔍 SET BREAKPOINT HERE — Back to outer scope
	fmt.Println("Outer x again:", x) // 👀 x is still 10
	fmt.Printf("Outer x address: %p\n", &x)

	// Loop variable shadowing
	// 🔍 SET BREAKPOINT HERE
	for i := 0; i < 3; i++ {
		// Each iteration uses the SAME i variable
		fmt.Printf("Loop i=%d, address=%p\n", i, &i)
	}

	// Closure capture gotcha
	// 🔍 SET BREAKPOINT HERE
	var funcs []func()
	for i := 0; i < 3; i++ {
		// ⚠️ All closures capture the SAME i variable
		funcs = append(funcs, func() {
			fmt.Println("Captured i:", i) // 👀 What value will this print?
		})
	}

	// 🔍 SET BREAKPOINT HERE — Before calling closures
	fmt.Println("\nCalling closures:")
	for _, f := range funcs {
		f() // ⚠️ What will each function print?
	}

	// Correct closure capture
	// 🔍 SET BREAKPOINT HERE
	var correctFuncs []func()
	for i := 0; i < 3; i++ {
		i := i // 👀 Shadow the loop variable to create a new variable per iteration
		correctFuncs = append(correctFuncs, func() {
			fmt.Println("Correctly captured i:", i)
		})
	}

	// 🔍 SET BREAKPOINT HERE
	fmt.Println("\nCalling correct closures:")
	for _, f := range correctFuncs {
		f()
	}
}
