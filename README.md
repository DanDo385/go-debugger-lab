# Go Debugger Lab

**A hands-on laboratory for learning Delve and mastering Go runtime observation.**

This is not a tutorial article. This is a debugging laboratory where you build intuition by watching execution happen.

## Who This Is For

You already understand:
- Functions, structs, pointers
- Goroutines and channels
- How to write tests

You do NOT yet:
- "See" execution flow through stack frames
- Know where values live in memory
- Understand what goroutines are actually doing
- Trust the debugger when it shows surprising behavior

After this lab, you will.

---

## What is Delve?

Delve (`dlv`) is Go's debugger. It lets you:
- Pause execution at any line
- Inspect variables, memory addresses, goroutines
- Step through code line by line
- Watch how values change over time

**What Delve is NOT:**
- A magic bug finder
- A substitute for understanding your code
- Always truthful (optimizations lie to debuggers)

Delve shows you *what actually happened at runtime*. Your job is to observe and understand.

---

## Core Debugging Concepts

### Step Over vs Step Into vs Step Out

| Command | What it does | When to use |
|---------|-------------|-------------|
| **Step Over** (`F10`) | Execute the current line, don't go into functions | You trust the function, just want to see the result |
| **Step Into** (`F11`) | Enter the function being called | You want to see what happens inside |
| **Step Out** (`Shift+F11`) | Finish current function and return to caller | You've seen enough, want to go back up |
| **Continue** (`F5`) | Run until next breakpoint or program end | You want to skip to the interesting part |

### The Call Stack

The call stack shows the chain of function calls that led to the current line.

```
main.main()          ← You are here
  ↑
processRequest()     ← Called by main
  ↑
validateInput()      ← Called by processRequest
```

Each level is a **stack frame** containing:
- Function parameters
- Local variables
- Return address (where to go back)

**Click on any frame** to see its variables.

### Stack vs Heap (Practical View)

**Stack:**
- Fast, automatic cleanup
- Limited size
- Local variables that don't escape

**Heap:**
- Slower, garbage collected
- Can grow as needed
- Values that outlive their function

**How to tell in the debugger:**
```go
x := 42          // Probably stack
y := &x          // x might escape to heap now
```

Watch the memory addresses. If a value's address is reused after a function returns, it was on the stack. If it persists, it escaped to the heap.

### Why Goroutines Feel Weird

When you step through code with goroutines:
- You're stepping *one goroutine at a time*
- Other goroutines keep running (or are paused, depending on the debugger)
- Channels may block unexpectedly
- Race conditions disappear (debugger changes timing)

**The debugger is not the real world.** Use it to understand structure, not exact concurrency behavior.

### Why Optimized Builds Lie

Compiler optimizations include:
- **Inlining**: Function calls disappear
- **Dead code elimination**: Unused variables vanish
- **Register allocation**: Variables have no memory address

**Always debug with optimizations OFF:**
```bash
go build -gcflags=all=-N -l
```

Our VS Code config does this automatically.

When you debug an optimized binary:
- Variables show `<optimized out>`
- Breakpoints may not trigger
- Step over might skip multiple lines

This is why production debugging is hard.

---

## Installation

### Install Delve
```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Verify:
```bash
dlv version
```

### Install VS Code Go Extension
1. Open VS Code
2. Install the official **Go** extension (by Google)
3. Open this repository
4. VS Code will prompt to install Go tools → accept

---

## How to Use This Lab

### 1. Start with Module 01
Each module is self-contained. Go in order:
```
01-main-and-entrypoint/
02-variables-and-scope/
...
```

### 2. Read the Module README
Each module has a `README.md` explaining:
- What you'll observe
- Where to set breakpoints
- What questions to answer

### 3. Set Breakpoints
Look for markers in the code:
```go
// 🔍 SET BREAKPOINT HERE
```

Click in the left margin to set a breakpoint (red dot appears).

### 4. Start Debugging
**Option A:** Use the pre-configured launch configs
- Open "Run and Debug" panel (`Ctrl+Shift+D` / `Cmd+Shift+D`)
- Select the appropriate config (e.g., "Debug Module 01")
- Press `F5`

**Option B:** Debug current folder
- Open any `.go` file in a module
- Press `F5`
- Select "Debug Current Folder"

### 5. Observe, Don't Rush
At each breakpoint:
- Look at the **Variables** panel
- Check the **Call Stack**
- Watch values change as you step

**Answer the module's questions before moving on.**

### 6. Reset Your Mental State
Between modules, close the debug session and take a breath. Each module builds a specific mental model.

---

## Debugging Thinking, Not Just Bugs

Debuggers are for:
- **Understanding**: "How does this code actually work?"
- **Validation**: "Does this value match my assumptions?"
- **Discovery**: "Why does this behave differently than I expected?"

Debuggers are NOT for:
- Randomly stepping until something looks wrong
- Avoiding reading the code
- Replacing tests

**The goal is not to fix bugs. The goal is to see the runtime.**

---

## Success Criteria

After completing this lab, you should be able to:

✅ **Answer "where is this value stored?"** without guessing
✅ **Predict goroutine interleaving** by reading code
✅ **Recognize heap escapes** by inspecting addresses
✅ **Debug "impossible" bugs** by observing actual execution
✅ **Know when to trust the debugger** vs when it lies

If you can do these, you've developed debugging intuition.

---

## Module Overview

| Module | Focus | Key Lesson |
|--------|-------|------------|
| [01-main-and-entrypoint](01-main-and-entrypoint/) | Program startup | Execution begins before `main()` |
| [02-variables-and-scope](02-variables-and-scope/) | Variable shadowing | Same name ≠ same variable |
| [03-functions-and-call-stack](03-functions-and-call-stack/) | Stack frames | Every call creates a new frame |
| [04-pointers-and-memory](04-pointers-and-memory/) | Addresses and aliasing | Watch addresses, not just values |
| [05-slices-maps-and-aliasing](05-slices-maps-and-aliasing/) | Shared backing arrays | Mutation at a distance |
| [06-structs-and-methods](06-structs-and-methods/) | Receivers | Value vs pointer receivers matter |
| [07-interfaces-and-dynamic-dispatch](07-interfaces-and-dynamic-dispatch/) | Runtime types | Interfaces hold (type, value) pairs |
| [08-errors-and-defer](08-errors-and-defer/) | Defer execution | Deferred functions run AFTER return |
| [09-goroutines-basics](09-goroutines-basics/) | Concurrent execution | Why stepping feels broken |
| [10-channels-and-blocking](10-channels-and-blocking/) | Channel mechanics | Visualizing blocked goroutines |
| [11-data-races-and-sync](11-data-races-and-sync/) | Race conditions | Debugger changes behavior (Heisenbug) |
| [12-compiler-optimizations](12-compiler-optimizations/) | Optimization effects | Why variables "disappear" |
| [13-debugging-tests](13-debugging-tests/) | Test debugging | Debugging failing assertions |

---

## Tips

**Setting Conditional Breakpoints:**
- Right-click on a breakpoint → "Edit Breakpoint"
- Add a condition: `i == 5` or `user.Name == "Alice"`

**Watching Expressions:**
- Right-click a variable → "Add to Watch"
- Or manually add expressions in the Watch panel

**Seeing All Goroutines:**
- In the Call Stack panel, expand "Goroutines"
- Click on any goroutine to see its stack

**Keyboard Shortcuts:**
- `F5` — Continue
- `F10` — Step Over
- `F11` — Step Into
- `Shift+F11` — Step Out
- `Ctrl+Shift+F5` — Restart
- `Shift+F5` — Stop

---

## Philosophy

A debugger is not a crutch. It's a microscope.

You wouldn't study cells without a microscope. Don't study runtime behavior without a debugger.

**This lab teaches you to see what Go is actually doing.**

Now go to `01-main-and-entrypoint/` and start observing.
