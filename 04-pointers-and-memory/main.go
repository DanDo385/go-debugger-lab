package main

import "fmt"

// Pass by value — the function receives a COPY
// 🔍 SET BREAKPOINT HERE
func modifyValue(x int) {
	fmt.Printf("modifyValue received x=%d at address %p\n", x, &x)
	x = 999 // This modifies the COPY, not the original
	fmt.Printf("modifyValue changed x to %d\n", x)
}

// Pass by pointer — the function receives the ADDRESS
// 🔍 SET BREAKPOINT HERE
func modifyPointer(x *int) {
	fmt.Printf("modifyPointer received pointer %p, pointing to value %d\n", x, *x)
	*x = 999 // This modifies the ORIGINAL via the pointer
	fmt.Printf("modifyPointer changed *x to %d\n", *x)
}

// Return a pointer to a local variable
// 🤔 DOES THIS ESCAPE TO THE HEAP?
// 🔍 SET BREAKPOINT HERE
func createPointer() *int {
	local := 42
	fmt.Printf("createPointer: local=%d at address %p\n", local, &local)

	// 👀 Normally, local would be on the stack and disappear after return
	// But we're returning a pointer to it, so it ESCAPES to the heap
	return &local
}

// Return a value (stays on stack)
// 🔍 SET BREAKPOINT HERE
func createValue() int {
	local := 42
	fmt.Printf("createValue: local=%d at address %p\n", local, &local)
	return local // Just the value is returned, not the address
}

func main() {
	fmt.Println("=== Pass by Value ===")

	// 🔍 SET BREAKPOINT HERE
	original := 100
	fmt.Printf("Before modifyValue: original=%d at address %p\n", original, &original)

	modifyValue(original) // 👀 Watch: original is NOT modified

	// 🔍 SET BREAKPOINT HERE
	fmt.Printf("After modifyValue: original=%d (unchanged)\n\n", original)

	fmt.Println("=== Pass by Pointer ===")

	// 🔍 SET BREAKPOINT HERE
	original = 100
	fmt.Printf("Before modifyPointer: original=%d at address %p\n", original, &original)

	modifyPointer(&original) // 👀 Watch: original IS modified

	// 🔍 SET BREAKPOINT HERE
	fmt.Printf("After modifyPointer: original=%d (changed!)\n\n", original)

	fmt.Println("=== Heap Escape ===")

	// 🔍 SET BREAKPOINT HERE
	ptr := createPointer() // 👀 This variable escaped to the heap
	fmt.Printf("main: received pointer %p, pointing to value %d\n", ptr, *ptr)
	// The value still exists even though createPointer() returned!

	// 🔍 SET BREAKPOINT HERE
	val := createValue()
	fmt.Printf("main: received value %d\n\n", val)

	fmt.Println("=== Pointer Aliasing ===")

	// 🔍 SET BREAKPOINT HERE
	x := 50
	p1 := &x
	p2 := &x // Both pointers point to the same variable

	fmt.Printf("x=%d at address %p\n", x, &x)
	fmt.Printf("p1 points to address %p, value=%d\n", p1, *p1)
	fmt.Printf("p2 points to address %p, value=%d\n", p2, *p2)

	// 🔍 SET BREAKPOINT HERE
	*p1 = 100 // Modify via p1

	// 👀 Watch: both p1 and p2 see the change, and so does x
	fmt.Printf("After *p1=100: x=%d, *p1=%d, *p2=%d\n", x, *p1, *p2)
}
