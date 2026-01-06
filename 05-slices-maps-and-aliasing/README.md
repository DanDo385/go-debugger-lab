# Module 05: Slices, Maps, and Aliasing

## What You'll Learn
Slices and maps share memory in surprising ways. You'll observe shared backing arrays, see when append creates new memory, and understand mutation at a distance.

## What to Observe
- Slices created from slicing **share the backing array**
- Modifying through one slice **affects aliased slices**
- `append` may or may not create a **new backing array** (depends on capacity)
- Maps are **reference types** (always shared)
- `copy()` creates a **true copy** (separate backing array)

## Debugging Steps

### Step 1: Slice Aliasing
Set breakpoints at:
1. **Line 39** — After creating `original`
2. **Line 43** — After creating `aliased`
3. **Line 48** — After modifying `aliased[0]`

Start debugging.

At line 39:
- Expand `original` in the Variables panel
- Note: `len=5, cap=5`

Press `F5` to reach line 43.
- `aliased` is `original[1:4]`
- Expand `aliased`: `len=3, cap=4`
- 👀 **Both slices share the same backing array**

Press `F5` to reach line 48.
- You've just set `aliased[0] = 999`
- 👀 **Look at `original`** — `original[1]` is now `999`
- Why? Because `aliased[0]` and `original[1]` point to the same memory

### Step 2: Passing Slices to Functions
Set breakpoints at:
1. **Line 59** — Before calling `modifySlice`
2. **Line 8** — Inside `modifySlice`
3. **Line 63** — After `modifySlice` returns

Continue to line 59.
- `nums` is `[10, 20, 30]`

Press `F11` to step into `modifySlice`.
- The slice `s` is passed **by value** (the slice header is copied)
- But the backing array is **shared**

Press `F10` to execute `s[0] = 999`.
- Click on the `main` stack frame
- 👀 **Notice `nums[0]` is also `999`** — shared backing array

### Step 3: Append and Capacity
Set breakpoints at:
1. **Line 69** — After creating `small`
2. **Line 73** — Before calling `appendToSlice` (not reassigning)
3. **Line 20** — Inside `appendToSlice`

Continue to line 69.
- `small` is `[1, 2]`
- Expand it: `len=2, cap=2` (full capacity)

Press `F11` at line 73 to step into `appendToSlice`.
- At line 21, `append` is called
- Because capacity is full, `append` creates a **new backing array**
- Press `F10` to execute the append
- 👀 **Notice `cap` increased** (e.g., from `2` to `4`)

Press `Shift+F11` to return to `main`.
- 👀 **`small` is still `[1, 2]`**
- Why? The function modified its local slice header, not `main`'s
- The return value was ignored

Press `F5` to reach the reassignment line (line 79).
- This time, we capture the return value
- Now `small` is updated

### Step 4: Map Mutation
Set breakpoints at:
1. **Line 84** — Before calling `modifyMap`
2. **Line 28** — Inside `modifyMap`
3. **Line 89** — After `modifyMap` returns

Continue to line 84.
- `m` is `{"key": 42}`

Press `F11` to step into `modifyMap`.
- Maps are **reference types** (like pointers)
- Modifying `m` inside the function affects the original

Press `F10` to execute `m["key"] = 999`.
- Click on the `main` stack frame
- 👀 **`m["key"]` is now `999`** — maps share memory

### Step 5: Copy vs Alias
Set breakpoints at:
1. **Line 94** — After creating `src`, `alias`, and `cpy`
2. **Line 104** — After modifying `src[0]`

Continue to line 94.
- `alias` and `src` point to the **same backing array**
- `cpy` is a **true copy** with a different backing array

Press `F5` to reach line 104.
- `src[0]` was changed to `999`
- 👀 **Check all three slices:**
  - `src`: `[999, 2, 3]`
  - `alias`: `[999, 2, 3]` (shared)
  - `cpy`: `[1, 2, 3]` (separate copy)

## Questions to Answer

1. **When do slices share memory?**
   - Does slicing create a copy or an alias?
   - How can you tell if two slices share a backing array?

2. **Why doesn't `append` always modify the original slice?**
   - What happens when capacity is exceeded?
   - Should you always reassign the result of `append`?

3. **Are maps passed by value or by reference?**
   - Does passing a map to a function copy it?
   - Can a function modify the caller's map?

4. **How do you create a true copy of a slice?**
   - Is `alias := src` enough?
   - What does `copy()` do differently?

## Key Takeaway
**Slices are not arrays.** They're headers pointing to a backing array. Multiple slices can share the same backing array, causing mutation at a distance. Maps are always shared. Use `copy()` to create independent slices.
