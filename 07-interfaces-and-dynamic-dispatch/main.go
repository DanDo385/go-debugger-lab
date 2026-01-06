package main

import "fmt"

// Interface definition
type Speaker interface {
	Speak() string
}

// Two types implementing Speaker
type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

// Type that does NOT implement Speaker
type Rock struct {
	weight int
}

// Function that accepts the interface
// 🔍 SET BREAKPOINT HERE — Watch dynamic dispatch
func makeItSpeak(s Speaker) {
	// 👀 Inspect 's' in the Variables panel
	// You'll see the (type, value) pair
	fmt.Printf("Type: %T, Value: %+v\n", s, s)
	fmt.Println("Says:", s.Speak())
}

func main() {
	fmt.Println("=== Interface (Type, Value) Pairs ===")

	// 🔍 SET BREAKPOINT HERE
	var s Speaker

	// 👀 s is nil (no type, no value)
	fmt.Printf("s (uninitialized): %v, %T\n", s, s)

	// 🔍 SET BREAKPOINT HERE
	dog := Dog{name: "Buddy"}
	s = dog // Assign a Dog to the interface

	// 👀 Now s holds (Dog, {name: "Buddy"})
	fmt.Printf("s (holding Dog): %v, %T\n\n", s, s)

	fmt.Println("=== Dynamic Dispatch ===")

	// 🔍 SET BREAKPOINT HERE — Step Into (F11) makeItSpeak
	makeItSpeak(dog)

	// 🔍 SET BREAKPOINT HERE
	cat := Cat{name: "Whiskers"}
	makeItSpeak(cat) // Same function, different runtime type

	fmt.Println("\n=== Interface Holding Pointer vs Value ===")

	// 🔍 SET BREAKPOINT HERE
	s = dog         // Value receiver, interface holds a copy
	dogPtr := &dog
	s = dogPtr      // Interface holds a pointer

	// 👀 Inspect the interface holding pointer vs value
	fmt.Printf("s (holding *Dog): %v, %T\n\n", s, s)

	fmt.Println("=== Nil Interface vs Interface Holding Nil ===")

	// 🔍 SET BREAKPOINT HERE
	var nilInterface Speaker   // Nil interface (type=nil, value=nil)
	var nilPointer *Dog = nil
	var nonNilInterface Speaker = nilPointer // Interface holding nil pointer (type=*Dog, value=nil)

	// 👀 Watch these in the Variables panel
	fmt.Printf("nilInterface: %v, is nil? %v\n", nilInterface, nilInterface == nil)
	fmt.Printf("nonNilInterface: %v, is nil? %v\n", nonNilInterface, nonNilInterface == nil)

	// ⚠️ nonNilInterface is NOT nil, even though it holds a nil pointer!
	// This is a common gotcha

	fmt.Println("\n=== Type Assertions ===")

	// 🔍 SET BREAKPOINT HERE
	s = Dog{name: "Max"}

	// 👀 Step through these type assertions
	// Successful assertion
	if d, ok := s.(Dog); ok {
		fmt.Printf("s is a Dog: %+v\n", d)
	}

	// Failed assertion (using safe form)
	if _, ok := s.(Cat); !ok {
		fmt.Println("s is NOT a Cat")
	}

	// 🔍 SET BREAKPOINT HERE — Type switch
	describeType(dog)
	describeType(cat)
	describeType(Rock{weight: 100})
}

// 🔍 SET BREAKPOINT HERE — Observe type switch
func describeType(i interface{}) {
	// 👀 Watch the runtime type determination
	switch v := i.(type) {
	case Dog:
		fmt.Printf("It's a dog named %s\n", v.name)
	case Cat:
		fmt.Printf("It's a cat named %s\n", v.name)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
