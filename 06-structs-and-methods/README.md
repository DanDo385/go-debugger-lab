# Module 06: Structs and Methods

## What You'll Learn
Value receivers get a copy. Pointer receivers get the address. You'll observe how method receivers determine whether mutations affect the original struct.

## What to Observe
- **Value receiver** methods receive a **copy** of the struct
- **Pointer receiver** methods receive the **address** of the struct
- Struct assignment **copies** the entire struct
- Go automatically takes the address when calling pointer receiver methods on values

## Debugging Steps

### Step 1: Value Receiver (No Mutation)
Set breakpoints at:
1. **Line 44** — Before calling `IncrementValue`
2. **Line 12** — Inside `IncrementValue` (value receiver)
3. **Line 49** — After the method returns

Start debugging.

At line 44:
- Note `c1.value = 10`
- Note the **address** of `c1` (e.g., `0xc000014080`)

Press `F11` to step into `IncrementValue`.
- Look at the receiver `c` in the Variables panel
- 👀 **Compare the address of `c` with the address of `c1`**
- They're DIFFERENT — this is a copy

Press `F10` to execute `c.value++`.
- `c.value` becomes `11`
- Click on the `main` stack frame
- 👀 **`c1.value` is still `10`** — the method modified a copy

### Step 2: Pointer Receiver (Mutation)
Set breakpoints at:
1. **Line 57** — Before calling `IncrementPointer`
2. **Line 20** — Inside `IncrementPointer` (pointer receiver)
3. **Line 62** — After the method returns

Continue to line 57.
- `c2.value = 10`
- Note the address of `c2`

Press `F11` to step into `IncrementPointer`.
- The receiver `c` is a **pointer** (`*Counter`)
- 👀 **Look at the value of `c`** — it's the address of `c2`
- Expand `c` to see `c.value`

Press `F10` to execute `c.value++`.
- Click on the `main` stack frame
- 👀 **`c2.value` is now `11`** — the method modified the original

### Step 3: Returning Modified Structs
Set breakpoints at:
1. **Line 68** — Before calling `IncrementAndReturn`
2. **Line 27** — Inside `IncrementAndReturn`

Continue to line 68 and step into the method.

- The receiver is a copy (value receiver)
- Press `F10` to execute `c.value++`
- The copy is incremented
- Press `F10` to return

Back in `main`:
- 👀 The return value is **assigned** to `c3`
- Now `c3.value` is `11`

### Step 4: Automatic Address-Taking
Set a breakpoint at **line 79** (`c4.Reset()`).

Continue to line 79.
- `Reset` has a **pointer receiver**
- But we're calling it on a **value** (`c4`)

Press `F11` to step into `Reset`.
- 👀 The receiver `c` is a **pointer** to `c4`
- Go automatically converted `c4.Reset()` to `(&c4).Reset()`

### Step 5: Struct Copying
Set breakpoints at:
1. **Line 86** — After `copied := original`
2. **Line 93** — After modifying `copied.value`

Continue to line 86.
- Expand both `original` and `copied` in the Variables panel
- 👀 **Compare their addresses** — they're different
- Struct assignment creates a copy

Press `F5` to reach line 93.
- `copied.value` is `999`
- 👀 **`original.value` is still `50`** — they're separate

## Questions to Answer

1. **When should you use a pointer receiver?**
   - If you need to modify the struct?
   - If the struct is large (to avoid copying)?
   - If the struct should not be copied (e.g., contains a mutex)?

2. **What happens if you mix value and pointer receivers?**
   - Can the same type have both?
   - How does Go decide which to call?

3. **Does struct assignment create a copy or an alias?**
   - Compare with slices and maps
   - What about nested structs with pointers?

4. **Why does `c4.Reset()` work even though Reset has a pointer receiver?**
   - What does Go do automatically?
   - Can you call a value receiver method on a pointer?

## Key Takeaway
**Value receivers get a copy. Pointer receivers get the address.** If a method needs to mutate the struct, use a pointer receiver. Go automatically takes the address when needed, but you still need to understand what's happening.
