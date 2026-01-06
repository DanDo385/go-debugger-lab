# Module 02: Variables and Scope

## What You'll Learn
Same name does not mean same variable. You'll observe variable shadowing, closure capture, and how scope creates separate memory locations.

## What to Observe
- Shadowed variables have **different memory addresses**
- Loop variables are **reused** across iterations
- Closures capture **variables, not values**

## Debugging Steps

### Step 1: Variable Shadowing
Set breakpoints at:
1. **Line 7** — Outer `x := 10`
2. **Line 13** — Inner `x := 20`
3. **Line 21** — Back to outer scope

Start debugging and step through:
- At line 7, look at `x` in the **Variables** panel
- Step to line 13 and **expand the Variables panel**
  - Notice there are now TWO `x` variables visible
  - One from the outer scope, one from the inner scope
- 👀 **Copy the addresses** — are they the same or different?

Press `F10` to step out of the inner block.
- The inner `x` disappears
- The outer `x` is still `10`

### Step 2: Loop Variable Address
Set a breakpoint at **line 27** (inside the loop).

Continue (`F5`) to the loop.
- Each time you hit the breakpoint, check the **address of `i`**
- Press `F5` multiple times to iterate through the loop
- ⚠️ **Notice: the address of `i` stays the SAME** — it's the same variable being reused

### Step 3: Closure Capture Bug
Set breakpoints at:
1. **Line 33** — Inside the closure-building loop
2. **Line 43** — Before calling the closures

Continue to line 33.
- Look at the closure being created
- What value of `i` does it see?

Step through the loop and watch `i` change to `0, 1, 2`.

Now continue to line 43 and step through the closure calls.
- 👀 **Each closure prints `3`**, not `0, 1, 2`
- Why? They all captured the **same variable `i`**, which ended at `3` after the loop

### Step 4: Correct Closure Capture
Set a breakpoint at **line 51** (inside the corrected loop).

Watch how `i := i` creates a **new variable** in each iteration.
- The inner `i` shadows the loop `i`
- Each closure captures a **different variable**

Continue to the closure calls.
- 👀 Now they print `0, 1, 2` as expected

## Questions to Answer

1. **Do shadowed variables share memory?**
   - Compare the addresses of outer `x` and inner `x`
   - Are they the same variable or different?

2. **Why does the loop variable keep the same address?**
   - Is a new `i` created each iteration, or is the same `i` reused?

3. **What do closures capture: values or variables?**
   - If they captured values, the bug wouldn't happen
   - What actually gets captured?

4. **How does `i := i` fix the closure bug?**
   - What does the shadowing create?
   - How many distinct variables exist in the corrected loop?

## Key Takeaway
**Same name ≠ same variable.** Variables are defined by their scope and memory location. Shadowing creates new variables. Closures capture variables (references), not values.
