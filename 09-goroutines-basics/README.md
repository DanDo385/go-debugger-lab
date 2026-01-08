# Module 09: Goroutines Basics

## What You'll Learn
Stepping through concurrent code feels broken. You'll observe goroutines running concurrently, see the Goroutines panel, and understand why the debugger "jumps around."

## What to Observe
- Goroutines appear in the **Goroutines panel**
- Stepping can **switch between goroutines** unexpectedly
- Goroutines **run concurrently** (not just in the debugger)
- Closure capture bugs are **worse with goroutines** (they run later)

## Debugging Steps

### Step 1: Launching Goroutines
Set breakpoints at:
1. **Line 33** — Before launching goroutines
2. **Line 10** — Inside `worker` function
3. **Line 42** — After launching goroutines

Start debugging.

At line 33:
- Open the **Call Stack** panel
- Expand **Goroutines** section
- 👀 You should see `1 Goroutine(1) main.main`

Press `F10` three times to launch the goroutines.
- Check the Goroutines panel again
- 👀 **You should now see 4 goroutines:**
  - Main goroutine
  - Worker 1
  - Worker 2
  - Worker 3

Click on different goroutines in the panel.
- Each goroutine has its own **call stack**
- Each has its own **local variables**

### Step 2: Why Stepping Feels Broken
Set a breakpoint at **line 10** (inside `worker`).

Press `F5` (Continue).
- The debugger stops at `worker` — but which goroutine?
- 👀 **Check the Goroutines panel** — the active goroutine is highlighted

Press `F10` (Step Over).
- ⚠️ **The debugger might jump to a different goroutine**
- This is not a bug — all goroutines are running
- The debugger shows whichever one hits a breakpoint next

Press `F10` several times.
- You'll jump between `worker(1)`, `worker(2)`, `worker(3)` unpredictably
- 👀 **This is concurrency** — you can't predict the order

### Step 3: Goroutine Interleaving
Set breakpoints at:
1. **Line 52** — Before launching increment goroutines
2. **Line 21** — Inside `increment` function

Continue to line 52.
- `counter` is `0`

Press `F5` to hit the first goroutine.
- Check which goroutine ID you're in
- Step through and watch `counter` change

Press `F5` repeatedly to see different goroutines.
- Each goroutine reads and modifies `counter`
- 👀 **The order is unpredictable**

What's the final value of `counter`?
- You might expect `3`, but it could be different due to race conditions
- (We'll explore this in Module 11)

### Step 4: Closure Capture with Goroutines
Set a breakpoint at **line 71** (buggy closure).

Continue and step through the loop.
- Each `go func()` is launched
- But they **don't run immediately**

Press `F5` to let them run.
- 👀 **Most (or all) print `i = 3`**
- Why? The loop finished before the goroutines ran
- They all captured the same `i` variable, which ended at `3`

Compare with the fixed version (line 78):
- `i := i` creates a **new variable** per iteration
- Each goroutine captures a **different** `i`

### Step 5: Anonymous Goroutines
Set breakpoints at:
1. **Line 92** — Inside the anonymous goroutine
2. **Line 99** — Main waiting

Continue to line 99.
- Main goroutine is blocked on `<-done`

Check the Goroutines panel:
- 👀 **Two goroutines:**
  - Main (blocked)
  - Anonymous goroutine (running)

Press `F5` to jump to the anonymous goroutine.
- Step through and send to `done` channel
- Main goroutine unblocks and continues

## Questions to Answer

1. **Why does stepping "jump around"?**
   - Are you stepping through one goroutine or all of them?
   - How can you follow a single goroutine?

2. **How many goroutines are running?**
   - Check the Goroutines panel
   - What's the goroutine ID for `main`? (Usually 1)

3. **Why do all goroutines print `i = 3` in the buggy version?**
   - When do the goroutines actually run?
   - What value does `i` have by then?

4. **Can you set a breakpoint that only triggers for one specific goroutine?**
   - Try using conditional breakpoints with goroutine IDs

## Key Takeaway
**You're debugging one goroutine at a time, but all of them are running.** Stepping can jump between goroutines unpredictably. Use the Goroutines panel to see what's running. The debugger changes timing, so concurrency bugs may disappear.
