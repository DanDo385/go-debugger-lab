# Module 04: Pointers and Memory

## What You'll Learn
Pointers are addresses. You'll observe pass-by-value vs pass-by-pointer, see how memory addresses reveal aliasing, and watch values escape to the heap.

## What to Observe
- Pass-by-value creates a **copy** with a different address
- Pass-by-pointer allows **mutation** of the original
- Local variables can **escape to the heap** if their address is returned
- Multiple pointers to the same address create **aliasing**

## Debugging Steps

### Step 1: Pass by Value
Set breakpoints at:
1. **Line 45** — Before calling `modifyValue`
2. **Line 7** — Inside `modifyValue`
3. **Line 51** — After `modifyValue` returns

Start debugging.

At line 45:
- Note the **address** of `original` (e.g., `0xc000014080`)
- Note the **value**: `100`

Press `F11` to step into `modifyValue`.
- Look at the address of `x` inside the function
- 👀 **Compare the addresses** — they're DIFFERENT
- `x` is a copy of `original`

Press `F10` to execute `x = 999`.
- `x` changes to `999`
- But `original` (in the `main` stack frame) is still `100`

Press `Shift+F11` to step out.
- Back in `main`, `original` is still `100`
- The function modified a copy, not the original

### Step 2: Pass by Pointer
Set breakpoints at:
1. **Line 56** — Before calling `modifyPointer`
2. **Line 15** — Inside `modifyPointer`
3. **Line 61** — After `modifyPointer` returns

Continue to line 56.
- Note the address of `original` (e.g., `0xc000014088`)

Press `F11` to step into `modifyPointer`.
- The parameter `x` is a **pointer** (address) to `original`
- 👀 **Check the value of `x`** — it's the address of `original`
- Expand `x` in the Variables panel to see `*x` (the value it points to)

Press `F10` to execute `*x = 999`.
- Now `*x` is `999`
- Click on the `main` stack frame
- 👀 **Notice `original` is also `999`** — same memory location

### Step 3: Heap Escape
Set breakpoints at:
1. **Line 67** — Before calling `createPointer`
2. **Line 25** — Inside `createPointer`

Continue to line 67 and step into `createPointer`.

At line 25:
- `local` is `42`
- Note its address (e.g., `0xc000014090`)

The function returns `&local` — a pointer to a local variable.
- Normally, local variables disappear when the function returns
- But we're returning the address, so `local` **escapes to the heap**

Press `Shift+F11` to return to `main`.
- 👀 **The pointer is still valid**
- The value `42` still exists, even though `createPointer()` finished
- This is heap allocation

### Step 4: Pointer Aliasing
Set breakpoints at:
1. **Line 77** — After creating `p1` and `p2`
2. **Line 84** — After modifying `*p1`

Continue to line 77.
- Expand `p1` and `p2` in the Variables panel
- 👀 **Both point to the same address** (the address of `x`)

Press `F5` to continue to line 84.
- `*p1` was set to `100`
- Check `x`, `*p1`, and `*p2`
- 👀 **All three show `100`** — they're all the same memory location

## Questions to Answer

1. **How can you tell if a function will modify the original variable?**
   - Check the function signature
   - Does it take a value or a pointer?

2. **What does "escape to the heap" mean?**
   - Why doesn't `local` disappear after `createPointer` returns?
   - How does the compiler decide: stack or heap?

3. **Can you have multiple pointers to the same memory?**
   - What happens when you modify through one pointer?
   - Do the other pointers see the change?

4. **How do addresses reveal copying?**
   - If two variables have different addresses, are they the same memory?
   - Can two variables with the same address have different values?

## Key Takeaway
**Watch addresses, not just values.** Pointers are addresses. Pass-by-value copies memory. Pass-by-pointer shares memory. Variables escape to the heap when their addresses outlive their scope.
