# Module 08: Errors and Defer

## What You'll Learn
Deferred functions run AFTER return. You'll observe defer execution order, see how defer modifies named returns, and watch panic/recover in action.

## What to Observe
- `defer` **registers** a function but doesn't execute it immediately
- Deferred functions run in **LIFO order** (last deferred, first executed)
- Deferred functions execute **after the return statement** but **before returning to the caller**
- `defer` can **modify named return values**
- `recover()` **catches panics** but only inside deferred functions

## Debugging Steps

### Step 1: Basic Defer Execution Order
Set breakpoints at:
1. **Line 14** — `defer fmt.Println("Defer 1")`
2. **Line 20** — Before function returns

Start debugging and press `F5` to reach line 14.

- Notice the defer statement is **executed** (registered), but the function inside is **not called yet**
- Press `F10` three times to register all three defers

Press `F5` to reach line 20 (before return).
- The deferred functions **haven't run yet**

Press `F10` to return.
- 👀 **Watch the console**: defers run in reverse order (3, 2, 1)
- LIFO: Last In, First Out

### Step 2: Named Return with Defer
Set breakpoints at:
1. **Line 27** — Inside the deferred function
2. **Line 34** — Before `return result`

Continue to the `namedReturn` function.

Press `F5` to reach line 34.
- `result` is `"original"`

Press `F10` to execute the return statement.
- The debugger jumps to the deferred function (line 27)
- 👀 **`result` is still `"original"` here**

Press `F10` to modify `result`.
- Now `result` is `"original (modified by defer)"`

Press `F10` to finish the defer and return.
- 👀 **The returned value is the modified version**
- Defer runs AFTER `return result` but BEFORE the value leaves the function

### Step 3: Error Wrapping with Defer
Set breakpoints at:
1. **Line 42** — Inside the deferred error handler
2. **Line 51** — `return errors.New(...)`

Continue to `processWithError(true)`.

Press `F5` to reach line 51.
- An error is returned

The debugger jumps to the deferred function (line 42).
- 👀 **`err` is not nil** — the defer sees the error
- Press `F10` to wrap the error
- The final error is wrapped

### Step 4: Panic and Recover
Set breakpoints at:
1. **Line 60** — Inside the recovery handler
2. **Line 67** — `panic("something went wrong!")`

Continue to `panicAndRecover`.

Press `F5` to reach line 67.
- Press `F10` to execute `panic`

The debugger jumps to the deferred function (line 60).
- 👀 **`recover()` returns the panic value**
- The panic is caught, execution continues

Back in `main`, the next line after `panicAndRecover()` executes.
- 👀 **The program didn't crash** — panic was recovered

### Step 5: Defer in Loop (Capture Bug)
Set a breakpoint at **line 79** (after the loop in `deferInLoop`).

Continue and step through the loop.
- Each iteration **registers** a defer
- The loop variable `i` is captured

Press `F5` to reach line 79 (after loop).
- The loop has finished, `i` is `3`

Press `F10` to return from the function.
- 👀 **All defers print `i = 3`** — they all captured the same variable

### Step 6: Defer in Loop (Fixed)
Repeat the same steps for `deferInLoopFixed`.

This time:
- `i := i` creates a **new variable** per iteration
- Each defer captures a **different** `i`
- 👀 **Defers print `2, 1, 0`** (in LIFO order)

## Questions to Answer

1. **When do deferred functions actually execute?**
   - Before or after the return statement?
   - Can they modify return values?

2. **What order do multiple defers execute?**
   - FIFO (first in, first out) or LIFO (last in, first out)?
   - Why might this order matter?

3. **Can defer modify named return values?**
   - Try changing `namedReturn` to return a non-named value
   - Does defer still affect the return?

4. **Where does recover() work?**
   - Try calling `recover()` outside a defer — what happens?
   - Can you recover from a panic in a different goroutine?

5. **Why does the loop variable capture happen?**
   - Compare with Module 02 (closure capture)
   - How is this the same bug?

## Key Takeaway
**Deferred functions run AFTER return, in LIFO order.** They can modify named return values. `recover()` only works inside `defer`. Defer captures variables, not values (same closure bug as before).
