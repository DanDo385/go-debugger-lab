# Module 11: Data Races and Sync

## What You'll Learn
The debugger changes race conditions. You'll observe intentional races, see how the debugger masks bugs (Heisenbug), and learn when NOT to trust the debugger.

## What to Observe
- **Data races** cause unpredictable behavior
- The **debugger slows execution**, hiding races
- **Mutexes** protect shared data
- The **race detector** (`-race`) is more reliable than the debugger for concurrency bugs

## Debugging Steps

### Step 1: Intentional Data Race
Set a breakpoint at **line 12** (inside `racyCounter`).

**IMPORTANT:** First run **without the debugger:**
```bash
cd 11-data-races-and-sync
go run main.go
```

Note the result. Run it several times — the count varies!

Now run with the **race detector:**
```bash
go run -race main.go
```

👀 **The race detector reports the race.**

Now debug it:
- Set breakpoints and step through
- 👀 **The race might disappear** in the debugger
- Why? The debugger slows execution, changing timing

### Step 2: Fixed with Mutex
Set a breakpoint at **line 44** (inside the mutex-protected loop).

Step through:
- Watch `mu.Lock()` and `mu.Unlock()`
- Only one goroutine can hold the lock at a time
- 👀 **Counter reaches 10000 reliably**

### Step 3: Heisenbug (Race Disappears)
Set breakpoints at:
1. **Line 85** — Inside `heisenbug`
2. **Line 90** — Writer goroutine
3. **Line 96** — Reader goroutine

Step through in the debugger:
- The race might not happen
- Reader might always see `value = 42`

Now run **without the debugger**:
```bash
go run -race main.go
```

👀 **The race detector may report the race**, even though the debugger didn't show it.

**This is a Heisenbug**: a bug that disappears when you observe it.

### Step 4: WaitGroup
Set breakpoints at:
1. **Line 124** — `wg.Add(1)`
2. **Line 128** — `defer wg.Done()`
3. **Line 135** — `wg.Wait()`

Step through:
- `wg.Add(1)` increments the wait group counter
- `wg.Done()` decrements it
- `wg.Wait()` blocks until counter reaches zero

## Questions to Answer

1. **Why does `racyCounter` produce different results each time?**
   - What happens when two goroutines read and increment simultaneously?
   - Can increments be "lost"?

2. **Why does the race disappear in the debugger?**
   - How does the debugger affect timing?
   - When should you NOT trust the debugger?

3. **How does a mutex prevent the race?**
   - What does `Lock()` guarantee?
   - Can two goroutines hold the lock simultaneously?

4. **When should you use the race detector instead of the debugger?**
   - For logical bugs or concurrency bugs?
   - Does the race detector change timing?

5. **Try removing the mutex from `mutexCounter`:**
   - Run with `-race` — what does it report?
   - Run in debugger — does it always fail?

## Key Takeaway
**The debugger changes race conditions.** Races depend on timing, and the debugger slows execution. For concurrency bugs, trust the **race detector** (`go run -race`), not the debugger. Use the debugger to understand structure, not to verify race-free execution.
