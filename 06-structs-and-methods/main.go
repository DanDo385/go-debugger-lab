package main

import "fmt"

type Counter struct {
	value int
	name  string
}

// Value receiver — receives a COPY of the struct
// 🔍 SET BREAKPOINT HERE
func (c Counter) IncrementValue() {
	fmt.Printf("IncrementValue (before): c.value=%d, address=%p\n", c.value, &c)
	c.value++ // This modifies the COPY, not the original
	fmt.Printf("IncrementValue (after): c.value=%d, address=%p\n", c.value, &c)
}

// Pointer receiver — receives a POINTER to the struct
// 🔍 SET BREAKPOINT HERE
func (c *Counter) IncrementPointer() {
	fmt.Printf("IncrementPointer (before): c.value=%d, address=%p\n", c.value, c)
	c.value++ // This modifies the ORIGINAL
	fmt.Printf("IncrementPointer (after): c.value=%d, address=%p\n", c.value, c)
}

// Value receiver that returns the modified struct
// 🔍 SET BREAKPOINT HERE
func (c Counter) IncrementAndReturn() Counter {
	c.value++
	return c // Returns a new struct with the incremented value
}

// 👀 Method with pointer receiver can be called on value
// Go automatically takes the address
func (c *Counter) Reset() {
	c.value = 0
}

func main() {
	fmt.Println("=== Value Receiver ===")

	// 🔍 SET BREAKPOINT HERE
	c1 := Counter{value: 10, name: "c1"}
	fmt.Printf("Before IncrementValue: c1.value=%d, address=%p\n", c1.value, &c1)

	// 🔍 SET BREAKPOINT HERE — Step Into (F11) to see the copy
	c1.IncrementValue()

	// 🔍 SET BREAKPOINT HERE
	fmt.Printf("After IncrementValue: c1.value=%d (unchanged)\n\n", c1.value)

	fmt.Println("=== Pointer Receiver ===")

	// 🔍 SET BREAKPOINT HERE
	c2 := Counter{value: 10, name: "c2"}
	fmt.Printf("Before IncrementPointer: c2.value=%d, address=%p\n", c2.value, &c2)

	// 🔍 SET BREAKPOINT HERE — Step Into (F11) to see the pointer
	c2.IncrementPointer()

	// 🔍 SET BREAKPOINT HERE
	fmt.Printf("After IncrementPointer: c2.value=%d (changed!)\n\n", c2.value)

	fmt.Println("=== Returning Modified Struct ===")

	// 🔍 SET BREAKPOINT HERE
	c3 := Counter{value: 10, name: "c3"}
	fmt.Printf("Before IncrementAndReturn: c3.value=%d\n", c3.value)

	c3 = c3.IncrementAndReturn() // Must reassign to see the change

	// 🔍 SET BREAKPOINT HERE
	fmt.Printf("After IncrementAndReturn: c3.value=%d\n\n", c3.value)

	fmt.Println("=== Automatic Address-Taking ===")

	// 🔍 SET BREAKPOINT HERE
	c4 := Counter{value: 100, name: "c4"}

	// Reset has a pointer receiver, but we can call it on a value
	// Go automatically converts this to (&c4).Reset()
	// 🔍 SET BREAKPOINT HERE — Step Into to see it receive a pointer
	c4.Reset()

	fmt.Printf("After Reset: c4.value=%d\n\n", c4.value)

	fmt.Println("=== Struct Copying ===")

	// 🔍 SET BREAKPOINT HERE
	original := Counter{value: 50, name: "original"}
	copied := original // This creates a COPY, not an alias

	fmt.Printf("original: value=%d, address=%p\n", original.value, &original)
	fmt.Printf("copied:   value=%d, address=%p\n", copied.value, &copied)

	// 🔍 SET BREAKPOINT HERE
	copied.value = 999

	// 👀 original is unchanged because they're separate structs
	fmt.Printf("After copied.value=999:\n")
	fmt.Printf("  original: value=%d\n", original.value)
	fmt.Printf("  copied:   value=%d\n", copied.value)
}
