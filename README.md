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

## Complete Breakpoint Reference

This section lists all breakpoints across all modules. Use this as a quick reference when setting up your debugging session.

### Module 01: Main and Entrypoint
**File:** `01-main-and-entrypoint/main.go`

| Line | Description |
|------|-------------|
| 12 | Inside `init()` — observe initialization order |
| 18 | Start of `main()` — entry point |
| 25 | Inspect `os.Args` — command-line arguments |
| 40 | Right before `main()` exits |

### Module 02: Variables and Scope
**File:** `02-variables-and-scope/main.go`

| Line | Description |
|------|-------------|
| 7 | Outer `x := 10` — initial variable |
| 13 | Inner `x := 20` — shadowed variable |
| 21 | Back to outer scope — verify outer `x` unchanged |
| 26 | Loop variable — observe address reuse |
| 33 | Closure capture bug — before loop |
| 42 | Before calling closures — see capture bug |
| 48 | Correct closure capture — fixed version |
| 57 | Before calling fixed closures |

### Module 03: Functions and Call Stack
**File:** `03-functions-and-call-stack/main.go`

| Line | Description |
|------|-------------|
| 6 | `deepFunction` — third level of call stack |
| 13 | `middleFunction` — second level of call stack |
| 21 | `topFunction` — first level of call stack |
| 30 | `factorial` — watch recursive stack growth |
| 51 | Call to `topFunction(5)` in `main()` |
| 57 | Call to `factorial(5)` — recursive example |
| 66 | `demonstrateStackFrames` — stack frame isolation |
| 73 | After helper call — verify variable isolation |
| 78 | `helperWithSameName` — different stack frame |

### Module 04: Pointers and Memory
**File:** `04-pointers-and-memory/main.go`

| Line | Description |
|------|-------------|
| 7 | `modifyValue` — pass by value function |
| 15 | `modifyPointer` — pass by pointer function |
| 24 | `createPointer` — heap escape example |
| 35 | `createValue` — stack allocation example |
| 45 | Before calling `modifyValue` — observe copy |
| 51 | After `modifyValue` — verify original unchanged |
| 56 | Before calling `modifyPointer` — observe mutation |
| 62 | After `modifyPointer` — verify original changed |
| 67 | Before `createPointer` — heap escape |
| 72 | Before `createValue` — stack allocation |
| 78 | Pointer aliasing — multiple pointers to same address |
| 87 | After modifying via pointer — verify aliasing |

### Module 05: Slices, Maps, and Aliasing
**File:** `05-slices-maps-and-aliasing/main.go`

| Line | Description |
|------|-------------|
| 8 | `modifySlice` — slice mutation function |
| 19 | `appendToSlice` — append behavior function |
| 28 | `modifyMap` — map mutation function |
| 38 | After creating `original` slice |
| 43 | After creating `aliased` slice — shared backing array |
| 49 | After modifying `aliased[0]` — observe mutation |
| 59 | Before calling `modifySlice` |
| 64 | After `modifySlice` returns — verify mutation |
| 70 | After creating `small` slice |
| 74 | Before calling `appendToSlice` (not reassigning) |
| 82 | Before reassigning result of `appendToSlice` |
| 88 | Before calling `modifyMap` |
| 94 | After `modifyMap` returns — verify mutation |
| 99 | After creating `src`, `alias`, and `cpy` |
| 109 | After modifying `src[0]` — compare alias vs copy |

### Module 06: Structs and Methods
**File:** `06-structs-and-methods/main.go`

| Line | Description |
|------|-------------|
| 12 | `IncrementValue` — value receiver method |
| 20 | `IncrementPointer` — pointer receiver method |
| 28 | `IncrementAndReturn` — return modified struct |
| 43 | Before calling `IncrementValue` |
| 47 | Step into `IncrementValue` — see copy |
| 50 | After `IncrementValue` — verify no mutation |
| 55 | Before calling `IncrementPointer` |
| 59 | Step into `IncrementPointer` — see pointer |
| 62 | After `IncrementPointer` — verify mutation |
| 67 | Before calling `IncrementAndReturn` |
| 73 | After `IncrementAndReturn` — verify return value |
| 78 | Before calling `Reset` (automatic address-taking) |
| 83 | Step into `Reset` — see pointer receiver |
| 90 | After struct copying — verify separate instances |
| 97 | After modifying copied struct — verify isolation |

### Module 07: Interfaces and Dynamic Dispatch
**File:** `07-interfaces-and-dynamic-dispatch/main.go`

| Line | Description |
|------|-------------|
| 34 | `makeItSpeak` — interface function (dynamic dispatch) |
| 45 | Uninitialized interface variable |
| 51 | After assigning `Dog` to interface |
| 60 | Step into `makeItSpeak` — observe dynamic dispatch |
| 63 | After calling with `Cat` — different runtime type |
| 69 | Interface holding pointer vs value |
| 79 | Nil interface vs interface holding nil |
| 93 | Before type assertion |
| 107 | Type switch — observe runtime type determination |
| 113 | `describeType` — type switch function |

### Module 08: Errors and Defer
**File:** `08-errors-and-defer/main.go`

| Line | Description |
|------|-------------|
| 10 | `deferredCleanup` — basic defer function |
| 14 | Defer registration — defer is registered, not executed |
| 21 | Before return — defers haven't run yet |
| 26 | `namedReturn` — named return with defer |
| 28 | Inside deferred function — modify named return |
| 35 | Before return — defer will modify return value |
| 40 | `processWithError` — error handling with defer |
| 42 | Deferred error handler — inspect/modify error |
| 59 | `panicAndRecover` — panic recovery |
| 61 | Recovery handler — catch panic |
| 70 | Before panic — observe panic behavior |
| 78 | `deferInLoop` — defer in loop (buggy) |
| 84 | After loop — all defers scheduled |
| 88 | `deferInLoopFixed` — fixed version |
| 99 | Step into `deferredCleanup` |
| 105 | Before calling `namedReturn` |
| 111 | Step into `processWithError` |
| 117 | Step into `panicAndRecover` |
| 123 | Before calling `deferInLoop` |
| 129 | Before calling `deferInLoopFixed` |

### Module 09: Goroutines Basics
**File:** `09-goroutines-basics/main.go`

| Line | Description |
|------|-------------|
| 10 | `worker` — simple goroutine function |
| 21 | `increment` — goroutine with shared variable |
| 31 | Main goroutine started |
| 35 | After launching goroutines — observe Goroutines panel |
| 44 | After goroutines launched — check panel |
| 54 | Before launching increment goroutines |
| 63 | After increment goroutines — observe race condition |
| 71 | Closure capture bug — goroutine version |
| 81 | Fixed closure capture — goroutine version |
| 93 | Before anonymous goroutine |
| 96 | Inside anonymous goroutine |
| 103 | Main waiting — observe blocking |

### Module 10: Channels and Blocking
**File:** `10-channels-and-blocking/main.go`

| Line | Description |
|------|-------------|
| 10 | `sender` — channel send function |
| 13 | Channel send — will block if unbuffered |
| 19 | `receiver` — channel receive function |
| 22 | Channel receive — will block until value arrives |
| 30 | Before creating unbuffered channel |
| 34 | Before launching receiver |
| 42 | Before sending — will unblock receiver |
| 50 | Before creating buffered channel |
| 54 | First send to buffered channel |
| 64 | First receive from buffered channel |
| 75 | Deadlock example (commented out) |
| 82 | Before creating channels for select |
| 98 | Select statement — waits for first available |
| 110 | Before creating closable channel |
| 118 | Before closing channel |
| 126 | Receiving from closed, empty channel |

### Module 11: Data Races and Sync
**File:** `11-data-races-and-sync/main.go`

| Line | Description |
|------|-------------|
| 11 | `racyCounter` — intentional data race |
| 34 | `mutexCounter` — fixed with mutex |
| 43 | Mutex lock/unlock — observe synchronization |
| 57 | `atomicCounter` — atomic operations |
| 79 | `heisenbug` — race that disappears in debugger |
| 110 | Before calling `racyCounter` |
| 115 | Before calling `mutexCounter` |
| 120 | Before calling `heisenbug` |
| 127 | Before WaitGroup example |
| 134 | Inside goroutine — defer WaitGroup.Done |
| 141 | WaitGroup.Wait — wait for all goroutines |

### Module 12: Compiler Optimizations
**File:** `12-compiler-optimizations/main.go`

| Line | Description |
|------|-------------|
| 7 | `add` — may be inlined (breakpoint may not trigger) |
| 13 | `calculate` — intermediate variables may be optimized out |
| 27 | `deadCode` — dead code elimination |
| 44 | `optimizedLoop` — loop optimization |
| 66 | Before calling `add` |
| 73 | Before calling `calculate` |
| 81 | Before calling `deadCode` |
| 84 | Before calling `optimizedLoop` |

### Module 13: Debugging Tests
**File:** `13-debugging-tests/calculator_test.go`

| Line | Description |
|------|-------------|
| 6 | Inside test cases — `TestAdd` |
| 25 | Inspect result before assertion |
| 36 | Normal division test |
| 49 | Division by zero test |
| 60 | Division with remainder test (will fail) |
| 90 | Step into `FindMax` |
| 93 | Before assertion — inspect result |
| 103 | Benchmark loop — will hit `b.N` times |
| 112 | Test helper function — observe stack trace |
| 121 | Before calling helper — verify helper behavior |

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
