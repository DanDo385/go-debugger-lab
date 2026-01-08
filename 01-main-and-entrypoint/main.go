package main

import (
	"fmt"
	"os"
)

// Global variable initialized before main()
var globalCounter int

// 🔍 SET BREAKPOINT HERE — init() runs BEFORE main()
func init() {
	globalCounter = 100
	fmt.Println("init() called, globalCounter =", globalCounter)
}

// 🔍 SET BREAKPOINT HERE — Execution enters main() after init()
func main() {
	fmt.Println("main() started")

	// 👀 WATCH globalCounter — it's already been set by init()
	fmt.Println("globalCounter in main() =", globalCounter)

	// Command-line arguments are available via os.Args
	// 🔍 SET BREAKPOINT HERE — Inspect os.Args in the Variables panel
	fmt.Println("Program name:", os.Args[0])
	if len(os.Args) > 1 {
		fmt.Println("Arguments:", os.Args[1:])
	} else {
		fmt.Println("No arguments provided")
	}

	// Environment variables
	// 👀 WATCH THIS — Look at the Variables panel for env
	env := os.Getenv("USER")
	if env != "" {
		fmt.Println("Running as user:", env)
	}

	// 🔍 SET BREAKPOINT HERE — Right before exit
	fmt.Println("main() finished")
}
