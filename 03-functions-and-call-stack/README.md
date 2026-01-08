# Module 03: Functions and Call Stack

## What You'll Learn
Every function call creates a new stack frame. You'll observe the call stack growing and shrinking, see how return values flow back up, and understand stack frames as containers for local variables.

## What to Observe
- Each function call adds a **stack frame**
- Stack frames contain parameters and local variables
- The call stack shows the **chain of calls** that led to the current point
- Recursive calls create **multiple frames of the same function**
- Clicking on different stack frames shows different variable scopes

## Debugging Steps

### Step 1: Nested Function Calls
Set breakpoints at:
1. **Line 52** — `topFunction(5)` call in `main()`
2. **Line 22** — Inside `topFunction`
3. **Line 13** — Inside `middleFunction`
4. **Line 6** — Inside `deepFunction`

Start debugging and press `F5` to reach line 52.

Press `F11` (Step Into) to enter `topFunction`.
- Look at the **Call Stack** panel
- You should see: `main()` → `topFunction()`

Press `F11` again at line 23 to enter `middleFunction`.
- Call Stack now shows: `main()` → `topFunction()` → `middleFunction()`
- 👀 Notice each frame has its own variables

Press `F11` at line 15 to enter `deepFunction`.
- Call Stack: `main()` → `topFunction()` → `middleFunction()` → `deepFunction()`
- 👀 **Click on each stack frame** to see the different `value` and `result` variables

Press `F10` (Step Over) repeatedly to watch the return values flow back up:
- `deepFunction` returns
- `middleFunction` receives the value
- `topFunction` receives the value
- Finally, `main` receives the final result

### Step 2: Recursive Stack Growth
Set a breakpoint at **line 30** (inside `factorial`).

Continue (`F5`) to the `factorial(5)` call at line 58.

Press `F11` to step into `factorial`.
- Look at the Call Stack
- Press `F5` to continue to the next recursive call
- Repeat several times

Watch the call stack grow:
```
main()
  factorial(5)
    factorial(4)
      factorial(3)
        factorial(2)
          factorial(1)  ← Base case
```

When the base case is reached (line 35), the stack starts **unwinding**.
- Press `F10` to step through the returns
- 👀 Watch each frame disappear as functions return

### Step 3: Stack Frames and Variable Isolation
Set breakpoints at:
1. **Line 66** — Inside `demonstrateStackFrames`
2. **Line 77** — Inside `helperWithSameName`

Continue to line 66.
- `x` is `100` in this frame

Press `F5` to jump to `helperWithSameName`.
- `x` is `200` here — a completely different variable
- 👀 **Click on the `demonstrateStackFrames` frame in the Call Stack panel**
  - Notice `x` is still `100` in that frame

Press `Shift+F11` (Step Out) to return to `demonstrateStackFrames`.
- `x` is still `100` — the other function's `x` didn't affect it

## Questions to Answer

1. **What is a stack frame?**
   - What does it contain?
   - When is it created and destroyed?

2. **How do return values travel back up the stack?**
   - Watch `deepFunction` return `30`
   - Where does that value go?
   - How does it become `35` in `topFunction`?

3. **Why doesn't recursion run forever?**
   - What stops the stack from growing indefinitely?
   - What would happen if you called `factorial(100000)`?

4. **Can two functions have variables with the same name without conflict?**
   - How does the debugger distinguish between them?
   - What happens to `x` in `helperWithSameName` after the function returns?

## Key Takeaway
**Every function call creates a new stack frame** containing parameters and local variables. The call stack is a history of how you got to the current line. Stack frames are created on call and destroyed on return.
