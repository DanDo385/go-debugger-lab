package main

import "fmt"

// Modify a slice passed by value
// ⚠️ This modifies the BACKING ARRAY, not a copy
// 🔍 SET BREAKPOINT HERE
func modifySlice(s []int) {
	fmt.Printf("modifySlice received: %v (len=%d, cap=%d)\n", s, len(s), cap(s))
	if len(s) > 0 {
		s[0] = 999 // 👀 This WILL affect the original slice
	}
	fmt.Printf("modifySlice changed to: %v\n", s)
}

// Append to a slice
// ⚠️ Append might create a NEW backing array
// 🔍 SET BREAKPOINT HERE
func appendToSlice(s []int) []int {
	fmt.Printf("appendToSlice received: %v (len=%d, cap=%d)\n", s, len(s), cap(s))
	s = append(s, 100) // 👀 This might or might not affect the caller's slice
	fmt.Printf("appendToSlice after append: %v (len=%d, cap=%d)\n", s, len(s), cap(s))
	return s
}

// Modify a map
// 🔍 SET BREAKPOINT HERE
func modifyMap(m map[string]int) {
	fmt.Printf("modifyMap received: %v\n", m)
	m["key"] = 999 // 👀 Maps are reference types, this WILL affect the original
	fmt.Printf("modifyMap changed to: %v\n", m)
}

func main() {
	fmt.Println("=== Slice Aliasing (Shared Backing Array) ===")

	// 🔍 SET BREAKPOINT HERE
	original := []int{1, 2, 3, 4, 5}
	fmt.Printf("original: %v (len=%d, cap=%d)\n", original, len(original), cap(original))

	// Create a slice that shares the backing array
	// 🔍 SET BREAKPOINT HERE
	aliased := original[1:4] // Shares the backing array with original
	fmt.Printf("aliased: %v (len=%d, cap=%d)\n", aliased, len(aliased), cap(aliased))

	// 👀 Watch both slices in the Variables panel
	// Modify through the aliased slice
	// 🔍 SET BREAKPOINT HERE
	aliased[0] = 999

	// ⚠️ This changes original[1] because they share memory
	fmt.Printf("After aliased[0]=999:\n")
	fmt.Printf("  original: %v\n", original)
	fmt.Printf("  aliased:  %v\n\n", aliased)

	fmt.Println("=== Passing Slices to Functions ===")

	// 🔍 SET BREAKPOINT HERE
	nums := []int{10, 20, 30}
	fmt.Printf("Before modifySlice: %v\n", nums)

	modifySlice(nums) // 👀 nums[0] will change

	// 🔍 SET BREAKPOINT HERE
	fmt.Printf("After modifySlice: %v (changed!)\n\n", nums)

	fmt.Println("=== Append and Capacity ===")

	// 🔍 SET BREAKPOINT HERE
	small := []int{1, 2}
	fmt.Printf("small: %v (len=%d, cap=%d)\n", small, len(small), cap(small))

	// Append without reassigning
	// 🔍 SET BREAKPOINT HERE — Step into appendToSlice
	appendToSlice(small)

	// ⚠️ small is UNCHANGED because append created a new backing array
	fmt.Printf("small after appendToSlice (not reassigned): %v\n", small)

	// Append with reassigning
	// 🔍 SET BREAKPOINT HERE
	small = appendToSlice(small)
	fmt.Printf("small after appendToSlice (reassigned): %v\n\n", small)

	fmt.Println("=== Map Aliasing ===")

	// 🔍 SET BREAKPOINT HERE
	m := map[string]int{"key": 42}
	fmt.Printf("Before modifyMap: %v\n", m)

	modifyMap(m) // 👀 Maps are reference types, m will change

	// 🔍 SET BREAKPOINT HERE
	fmt.Printf("After modifyMap: %v (changed!)\n\n", m)

	fmt.Println("=== Copy vs Alias ===")

	// 🔍 SET BREAKPOINT HERE
	src := []int{1, 2, 3}
	alias := src              // Just an alias, shares backing array
	cpy := make([]int, len(src))
	copy(cpy, src)            // Actual copy, different backing array

	fmt.Printf("src:   %v\n", src)
	fmt.Printf("alias: %v\n", alias)
	fmt.Printf("cpy:   %v\n", cpy)

	// 🔍 SET BREAKPOINT HERE
	src[0] = 999

	// 👀 alias changes, cpy does not
	fmt.Printf("After src[0]=999:\n")
	fmt.Printf("  src:   %v\n", src)
	fmt.Printf("  alias: %v (changed!)\n", alias)
	fmt.Printf("  cpy:   %v (unchanged)\n", cpy)
}
